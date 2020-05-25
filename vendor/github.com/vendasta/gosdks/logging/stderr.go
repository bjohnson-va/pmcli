package logging

import (
	"fmt"
	"math/rand"
	"os"
	"runtime"
	"strings"

	"cloud.google.com/go/logging"
	"golang.org/x/net/context"
)

const (
	colorRed = iota + 30
	colorGreen
	colorYellow
	colorMagenta
	colorCyan
)

type color int

var (
	colors = []string{
		logging.Debug:     colorSeq(colorMagenta),
		logging.Info:      colorSeq(colorRed),
		logging.Warning:   colorSeq(colorYellow),
		logging.Error:     colorSeq(colorGreen),
		logging.Critical:  colorSeq(colorCyan),
		logging.Alert:     colorSeq(colorCyan),
		logging.Emergency: colorSeq(colorCyan),
	}
)

func colorSeq(color color) string {
	return fmt.Sprintf("\033[%dm", int(color))
}

type stderrLogger struct {
	config *config
}

func (l *stderrLogger) request(ctx context.Context, r *requestData) {
	l.Log(ctx, logging.Debug, "Served gRPC request for handler %s with code %d", r.HTTPRequest.Request.URL.Path, r.HTTPRequest.Status)
}
func (l *stderrLogger) Debugf(ctx context.Context, f string, a ...interface{}) {
	l.Log(ctx, logging.Debug, f, a...)
}

func (l *stderrLogger) Infof(ctx context.Context, f string, a ...interface{}) {
	l.Log(ctx, logging.Info, f, a...)
}

func (l *stderrLogger) Warningf(ctx context.Context, f string, a ...interface{}) {
	l.Log(ctx, logging.Warning, f, a...)
}

func (l *stderrLogger) Errorf(ctx context.Context, f string, a ...interface{}) {
	l.Log(ctx, logging.Error, f, a...)
}

func (l *stderrLogger) Criticalf(ctx context.Context, f string, a ...interface{}) {
	l.Log(ctx, logging.Critical, f, a...)
}

func (l *stderrLogger) Alertf(ctx context.Context, f string, a ...interface{}) {
	l.Log(ctx, logging.Alert, f, a...)
}

func (l *stderrLogger) Emergencyf(ctx context.Context, f string, a ...interface{}) {
	l.Log(ctx, logging.Emergency, f, a...)
}

func (m *stderrLogger) Tag(ctx context.Context, key, value string) Logger {return m}

func (l *stderrLogger) Log(ctx context.Context, level logging.Severity, f string, a ...interface{}) {
	col := colors[level]
	_, file, line, ok := runtime.Caller(3)
	if !ok {
		file = "???"
		line = 0
	} else {
		filePieces := strings.Split(file, "/")
		file = strings.Join(filePieces[len(filePieces)-2:], "/")
	}
	if !strings.HasSuffix(f, "\n") {
		f = f + "\n"
	}
	prefix := fmt.Sprintf("%s%-6s%50s:%-4d\033[0m", col, level.String(), file, line)
	fmt.Fprintf(os.Stderr, prefix+" "+f, a...)
}

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func (l *stderrLogger) RequestID() string {
	b := make([]rune, 16)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func newStdErrLogger(config *config) (*stderrLogger, error) {
	return &stderrLogger{config: config}, nil
}
