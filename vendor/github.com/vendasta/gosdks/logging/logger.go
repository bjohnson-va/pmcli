package logging

import (
	"golang.org/x/net/context"
)

// Logger provides the methods for logging.
type Logger interface {
	request(ctx context.Context, r *requestData)
	Debugf(ctx context.Context, f string, a ...interface{})
	Infof(ctx context.Context, f string, a ...interface{})
	Warningf(ctx context.Context, f string, a ...interface{})
	Errorf(ctx context.Context, f string, a ...interface{})
	Criticalf(ctx context.Context, f string, a ...interface{})
	Alertf(ctx context.Context, f string, a ...interface{})
	Emergencyf(ctx context.Context, f string, a ...interface{})
	Tag(ctx context.Context, key, value string) Logger
	RequestID() string
}

var loggerInstance Logger

// GetLogger returns the current Logger instance.
func GetLogger() Logger {
	if loggerInstance == nil {
		loggerInstance = &stderrLogger{config: &config{}}
	}
	return loggerInstance
}

func logRequest(ctx context.Context, r *requestData) {
	GetLogger().request(ctx, r)
}

// Debugf emits a debug log
func Debugf(ctx context.Context, f string, a ...interface{}) {
	GetLogger().Debugf(ctx, f, a...)
}

// Infof emits a info log
func Infof(ctx context.Context, f string, a ...interface{}) {
	GetLogger().Infof(ctx, f, a...)
}

// Warningf emits a warning log
func Warningf(ctx context.Context, f string, a ...interface{}) {
	GetLogger().Warningf(ctx, f, a...)
}

// Errorf emits an error log
func Errorf(ctx context.Context, f string, a ...interface{}) {
	GetLogger().Errorf(ctx, f, a...)
}

// Criticalf emits a critical log
func Criticalf(ctx context.Context, f string, a ...interface{}) {
	GetLogger().Criticalf(ctx, f, a...)
}

// Alertf emits an alert log
func Alertf(ctx context.Context, f string, a ...interface{}) {
	GetLogger().Alertf(ctx, f, a...)
}

// Emergencyf emits an emergency log
func Emergencyf(ctx context.Context, f string, a ...interface{}) {
	GetLogger().Emergencyf(ctx, f, a...)
}

// Tag sets a tag on the request
func Tag(ctx context.Context, key, value string) Logger {
	return GetLogger().Tag(ctx, key, value)
}
