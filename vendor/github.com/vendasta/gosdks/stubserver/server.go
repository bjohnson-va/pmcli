package stubserver

import (
	"io"
	"net/http"
	"net/http/httptest"
)

// Server provides an interface for adding responses to be returned from the server
type Server interface {
	Start()
	Stop()
	URL() string
	PushResponse(r *Response)
}

// NewServer create a new server stub
func NewServer() Server {
	return &server{}
}

// Response represents a response which will be returned by the test server
type Response struct {
	StatusCode int
	Response   string
}

type server struct {
	server    *httptest.Server
	responses []*Response
}

func (s *server) handler(w http.ResponseWriter, r *http.Request) {
	if len(s.responses) == 0 {
		w.WriteHeader(http.StatusInternalServerError)
		io.WriteString(w, "internal error\n")
		return
	}
	w.WriteHeader(s.responses[0].StatusCode)
	io.WriteString(w, s.responses[0].Response)
	s.responses[0] = nil
	s.responses = s.responses[1:]
}

// Start the server
func (s *server) Start() {
	s.server = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		s.handler(w, r)
	}))
}

// Stop the server
func (s *server) Stop() {
	if s.server != nil {
		s.server.Close()
	}
}

// URL returns the url of the server
func (s *server) URL() string {
	if s.server != nil {
		return s.server.URL
	}
	return ""
}

// PushResponse adds a new response to the response queue
func (s *server) PushResponse(r *Response) {
	s.responses = append(s.responses, r)
}
