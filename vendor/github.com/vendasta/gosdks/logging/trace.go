package logging

import (
	crand "crypto/rand"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"
	"time"

	"context"

	cloudtrace "cloud.google.com/go/trace"
	api "google.golang.org/api/cloudtrace/v1"
	"google.golang.org/api/option"
	"google.golang.org/api/transport"
	"google.golang.org/grpc/metadata"
	gce_metadata "cloud.google.com/go/compute/metadata"
	"net/url"
)

const (
	httpHeader          = `x-cloud-trace-context`
	userAgent           = `gcloud-golang-trace/20160501`
	cloudPlatformScope  = `https://www.googleapis.com/auth/cloud-platform`
	spanKindClient      = `RPC_CLIENT`
	spanKindServer      = `RPC_SERVER`
	maxStackFrames      = 20
	labelHost           = `trace.cloud.google.com/http/host`
	labelMethod         = `trace.cloud.google.com/http/method`
	labelStackTrace     = `trace.cloud.google.com/stacktrace`
	labelStatusCode     = `trace.cloud.google.com/http/status_code`
	labelURL            = `trace.cloud.google.com/http/url`
	labelSamplingPolicy = `trace.cloud.google.com/sampling_policy`
	labelSamplingWeight = `trace.cloud.google.com/sampling_weight`
)

type contextKey struct{}

type stackLabelValue struct {
	Frames []stackFrame `json:"stack_frame"`
}

type stackFrame struct {
	Class    string `json:"class_name,omitempty"`
	Method   string `json:"method_name"`
	Filename string `json:"file_name"`
	Line     int64  `json:"line_number"`
}

var (
	spanIDCounter   uint64
	spanIDIncrement uint64
)

func init() {
	// Set spanIDCounter and spanIDIncrement to random values.  nextSpanID will
	// return an arithmetic progression using these values, skipping zero.  We set
	// the LSB of spanIDIncrement to 1, so that the cycle length is 2^64.
	binary.Read(crand.Reader, binary.LittleEndian, &spanIDCounter)
	binary.Read(crand.Reader, binary.LittleEndian, &spanIDIncrement)
	spanIDIncrement |= 1
}

// nextSpanID returns a new span ID.  It will never return zero.
func nextSpanID() uint64 {
	var id uint64
	for id == 0 {
		id = atomic.AddUint64(&spanIDCounter, spanIDIncrement)
	}
	return id
}

// nextTraceID returns a new trace ID.
func nextTraceID() string {
	id1 := nextSpanID()
	id2 := nextSpanID()
	return fmt.Sprintf("%016x%016x", id1, id2)
}

// Client is a client for uploading traces to the Google Stackdriver Trace server.
type Client struct {
	service   *api.Service
	projectID string
	policy    cloudtrace.SamplingPolicy
}

// NewClient creates a new Google Stackdriver Trace client.
func NewClient(ctx context.Context, projectID string, opts ...option.ClientOption) (*Client, error) {
	o := []option.ClientOption{
		option.WithScopes(cloudPlatformScope),
		option.WithUserAgent(userAgent),
	}
	o = append(o, opts...)
	hc, basePath, err := transport.NewHTTPClient(ctx, o...)
	if err != nil {
		return nil, fmt.Errorf("creating HTTP client for Google Stackdriver Trace API: %v", err)
	}
	apiService, err := api.New(hc)
	if err != nil {
		return nil, fmt.Errorf("creating Google Stackdriver Trace API client: %v", err)
	}
	if basePath != "" {
		// An option set a basepath, so override api.New's default.
		apiService.BasePath = basePath
	}
	policy, err := cloudtrace.NewLimitedSampler(0, 2)
	if err != nil {
		return nil, err
	}
	return &Client{
		service:   apiService,
		projectID: projectID,
		policy:    policy,
	}, nil
}

// SetSamplingPolicy sets the sampling policy.
func (client *Client) SetSamplingPolicy(p cloudtrace.SamplingPolicy) {
	if client != nil {
		client.policy = p
	}
}

// SpanFromContext returns a new trace span.  If the incoming HTTP request's
// headers don't specify that the request should be traced, and the sampling
// policy doesn't determine the request should be traced, the returned span
// will be nil.
// It also returns nil if the client is nil.
// When Finish is called on the returned span, the span and its descendants are
// uploaded to the Google Stackdriver Trace server.
func (client *Client) SpanFromContext(ctx context.Context, path string) *Span {
	if client == nil {
		return nil
	}
	span := traceInfoFromContext(ctx, path)
	if client.policy != nil {
		d := client.policy.Sample(cloudtrace.Parameters{HasTraceHeader: span != nil})
		if !d.Trace {
			return nil
		}
		if d.Sample {
			if span == nil {
				t := &trace{
					traceID: nextTraceID(),
					options: optionTrace,
					client:  client,
				}
				span = startNewChildWithContext(ctx, path, t, 0 /* parentSpanID */)
				span.span.Kind = spanKindServer
				span.rootSpan = true
			}
			span.SetLabel(labelSamplingPolicy, d.Policy)
			span.SetLabel(labelSamplingWeight, fmt.Sprint(d.Weight))
		}
	}

	if span == nil {
		return nil
	}
	span.trace.client = client
	return span
}

// NewContext returns a derived context containing the span.
func NewContext(ctx context.Context, s *Span) context.Context {
	if s == nil {
		return ctx
	}
	md, _ := metadata.FromOutgoingContext(ctx)  // This may not be necessary?
	ctx = metadata.NewOutgoingContext(ctx, metadata.Join(md, metadata.Pairs(httpHeader, s.spanHeader())))
	return context.WithValue(ctx, contextKey{}, s)
}

// FromContext returns the span contained in the context, or nil.
func FromContext(ctx context.Context) *Span {
	s, _ := ctx.Value(contextKey{}).(*Span)
	return s
}

func traceInfoFromContext(ctx context.Context, path string) *Span {
	// See https://cloud.google.com/trace/docs/faq for the header format.
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil
	}
	hl, ok := md["x-cloud-trace-context"]
	if !ok {
		return nil
	}
	if len(hl) != 1 {
		return nil
	}
	h := hl[0]

	// Return if the header is empty or missing, or if the header is unreasonably
	// large, to avoid making unnecessary copies of a large string.
	if h == "" || len(h) > 200 {
		return nil
	}

	// Parse the trace id field.
	slash := strings.Index(h, `/`)
	if slash == -1 {
		return nil
	}
	traceID, h := h[:slash], h[slash+1:]

	// Parse the span id field.
	semicolon := strings.Index(h, `;`)
	if semicolon == -1 {
		return nil
	}
	spanstr, h := h[:semicolon], h[semicolon+1:]
	spanID, err := strconv.ParseUint(spanstr, 10, 64)
	if err != nil {
		return nil
	}

	// Parse the options field.
	if !strings.HasPrefix(h, "o=") {
		return nil
	}
	o, err := strconv.ParseUint(h[2:], 10, 64)
	if err != nil {
		return nil
	}
	options := optionFlags(o)
	if options&optionTrace == 0 {
		return nil
	}

	t := &trace{
		traceID: traceID,
		options: options,
	}
	rootSpan := startNewChildWithContext(ctx, path, t, spanID)
	rootSpan.span.Kind = spanKindServer
	rootSpan.rootSpan = true
	return rootSpan
}

type optionFlags uint32

const (
	optionTrace optionFlags = 1 << iota
	optionStack
)

type trace struct {
	mu      sync.Mutex
	client  *Client
	traceID string
	options optionFlags
	spans   []*Span // finished spans for this trace.
}

// finish appends s to t.spans.  If s is the root span, uploads the trace to the
// server.
func (t *trace) finish(s *Span, wait bool, opts ...FinishOption) error {
	for _, o := range opts {
		o.modifySpan(s)
	}
	s.end = time.Now()
	t.mu.Lock()
	t.spans = append(t.spans, s)
	spans := t.spans
	t.mu.Unlock()
	if s.rootSpan {
		if wait {
			return t.upload(spans)
		}
		go func() {
			err := t.upload(spans)
			if err != nil {
				log.Println("error uploading trace:", err)
			}
		}()
	}
	return nil
}

func (t *trace) upload(spans []*Span) error {
	apiSpans := make([]*api.TraceSpan, len(spans))
	for i, sp := range spans {
		sp.span.StartTime = sp.start.In(time.UTC).Format(time.RFC3339Nano)
		sp.span.EndTime = sp.end.In(time.UTC).Format(time.RFC3339Nano)
		if t.options&optionStack != 0 {
			sp.setStackLabel()
		}
		sp.SetLabel(labelHost, sp.host)
		sp.SetLabel(labelURL, sp.url)
		sp.SetLabel(labelMethod, sp.method)
		if sp.statusCode != 0 {
			sp.SetLabel(labelStatusCode, strconv.Itoa(sp.statusCode))
		}
		apiSpans[i] = &sp.span
	}

	traces := &api.Traces{
		Traces: []*api.Trace{
			{
				ProjectId: t.client.projectID,
				TraceId:   t.traceID,
				Spans:     apiSpans,
			},
		},
	}
	_, err := t.client.service.Projects.PatchTraces(t.client.projectID, traces).Do()
	return err
}

func startTracing(context context.Context, requestURL *url.URL) (context.Context, *Span) {
	if gce_metadata.OnGCE() {
		span := configValue.TracingClient.SpanFromContext(context, requestURL.Path)
		return NewContext(context, span), span
	}
	return context, nil
}

// A span describes the amount of time it takes an application to complete a
// suboperation in a trace. For example, it can describe how long it takes for
// the application to perform a round-trip RPC call to another system when
// handling a request, or how long it takes to perform another task that is part
// of a larger operation.
type Span struct {
	trace      *trace
	span       api.TraceSpan
	start      time.Time
	end        time.Time
	rootSpan   bool
	stack      [maxStackFrames]uintptr
	host       string
	method     string
	url        string
	statusCode int
}

// NewChild creates a new span with the given name as a child of s.
// If s is nil, does nothing and returns nil.
func (s *Span) NewChild(name string) *Span {
	if s == nil {
		return nil
	}
	return startNewChild(name, s.trace, s.span.SpanId)
}

// NewRemoteChild creates a new span as a child of s.
// Span details are set from an outgoing *http.Request r.
// A header is set in r so that the trace context is propagated to the destination.
// If s is nil, does nothing and returns nil.
func (s *Span) NewRemoteChild(ctx context.Context, path string) (context.Context, *Span) {
	if s == nil {
		return ctx, nil
	}
	newSpan := startNewChildWithContext(ctx, path, s.trace, s.span.SpanId)
	md, _ := metadata.FromOutgoingContext(ctx)
	newMD := metadata.Join(md, metadata.New(map[string]string{httpHeader: newSpan.spanHeader()}))
	return metadata.NewOutgoingContext(ctx, newMD), newSpan
}

func startNewChildWithContext(ctx context.Context, path string, trace *trace, parentSpanID uint64) *Span {
	newSpan := startNewChild(path, trace, parentSpanID)
	newSpan.method = "POST"
	return newSpan
}

// NewHTTPRemoteChild creates a new span as a child of s.
// Span details are set from an outgoing *http.Request r.
// A header is set in r so that the trace context is propagated to the destination.
// If s is nil, does nothing and returns nil.
func (s *Span) NewHTTPRemoteChild(r *http.Request) *Span {
	if s == nil {
		return nil
	}
	newSpan := startNewChildWithRequest(r, s.trace, s.span.SpanId)
	r.Header.Set(httpHeader, newSpan.spanHeader())
	return newSpan
}

func startNewChildWithRequest(r *http.Request, trace *trace, parentSpanID uint64) *Span {
	newSpan := startNewChild(r.URL.Path, trace, parentSpanID)
	newSpan.method = r.Method
	newSpan.host = r.Host
	newSpan.url = r.URL.String()
	return newSpan
}

func startNewChild(name string, trace *trace, parentSpanID uint64) *Span {
	newSpan := &Span{
		trace: trace,
		span: api.TraceSpan{
			Kind:         spanKindClient,
			Name:         name,
			ParentSpanId: parentSpanID,
			SpanId:       nextSpanID(),
		},
		start: time.Now(),
	}
	if trace.options&optionStack != 0 {
		_ = runtime.Callers(1, newSpan.stack[:])
	}
	return newSpan
}

// TraceID returns the ID of the trace to which s belongs.
func (s *Span) TraceID() string {
	if s == nil {
		return ""
	}
	return s.trace.traceID
}

// SetLabel sets the label for the given key to the given value.
// If the value is empty, the label for that key is deleted.
// If a label is given a value automatically and by SetLabel, the
// automatically-set value is used.
// If s is nil, does nothing.
func (s *Span) SetLabel(key, value string) {
	if s == nil {
		return
	}
	if value == "" {
		if s.span.Labels != nil {
			delete(s.span.Labels, key)
		}
		return
	}
	if s.span.Labels == nil {
		s.span.Labels = make(map[string]string)
	}
	s.span.Labels[key] = value
}

// FinishOption allows a span to be modified.
type FinishOption interface {
	modifySpan(s *Span)
}

type withResponse struct {
	*http.Response
}

// WithResponse returns an option that can be passed to Finish that indicates
// that some labels for the span should be set using the given *http.Response.
func WithResponse(resp *http.Response) FinishOption {
	return withResponse{resp}
}
func (u withResponse) modifySpan(s *Span) {
	if u.Response != nil {
		s.statusCode = u.StatusCode
	}
}

// Finish declares that the span has finished.
//
// If s is nil, Finish does nothing and returns nil.
//
// If the option trace.WithResponse(resp) is passed, then some labels are set
// for s using information in the given *http.Response.  This is useful when the
// span is for an outgoing http request; s will typically have been created by
// NewRemoteChild in this case.
//
// If s is a root span (one created by SpanFromRequest) then s, and all its
// descendant spans that have finished, are uploaded to the Google Stackdriver
// Trace server asynchronously.
func (s *Span) Finish(opts ...FinishOption) {
	if s == nil {
		return
	}
	s.trace.finish(s, false, opts...)
}

// FinishWait is like Finish, but if s is a root span, it waits until uploading
// is finished, then returns an error if one occurred.
func (s *Span) FinishWait(opts ...FinishOption) error {
	if s == nil {
		return nil
	}
	return s.trace.finish(s, true, opts...)
}

func (s *Span) spanHeader() string {
	// See https://cloud.google.com/trace/docs/faq for the header format.
	return fmt.Sprintf("%s/%d;o=%d", s.trace.traceID, s.span.SpanId, s.trace.options)
}

func (s *Span) setStackLabel() {
	var stack stackLabelValue
	lastSigPanic, inTraceLibrary := false, true
	for _, pc := range s.stack {
		if pc == 0 {
			break
		}
		if !lastSigPanic {
			pc--
		}
		fn := runtime.FuncForPC(pc)
		file, line := fn.FileLine(pc)
		// Name has one of the following forms:
		// path/to/package.Foo
		// path/to/package.(Type).Foo
		// For the first form, we store the whole name in the Method field of the
		// stack frame.  For the second form, we set the Method field to "Foo" and
		// the Class field to "path/to/package.(Type)".
		name := fn.Name()
		if inTraceLibrary && !strings.HasPrefix(name, "cloud.google.com/go/trace.") {
			inTraceLibrary = false
		}
		var class string
		if i := strings.Index(name, ")."); i != -1 {
			class, name = name[:i+1], name[i+2:]
		}
		frame := stackFrame{
			Class:    class,
			Method:   name,
			Filename: file,
			Line:     int64(line),
		}
		if inTraceLibrary && len(stack.Frames) == 1 {
			stack.Frames[0] = frame
		} else {
			stack.Frames = append(stack.Frames, frame)
		}
		lastSigPanic = fn.Name() == "runtime.sigpanic"
	}
	if label, err := json.Marshal(stack); err == nil {
		s.SetLabel(labelStackTrace, string(label))
	}
}
