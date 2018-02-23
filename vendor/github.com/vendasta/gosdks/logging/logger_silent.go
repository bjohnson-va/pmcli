package logging

import "golang.org/x/net/context"

type silentLogger struct{}

func (m silentLogger) request(ctx context.Context, r *requestData)                {}
func (m silentLogger) Debugf(ctx context.Context, f string, a ...interface{})     {}
func (m silentLogger) Infof(ctx context.Context, f string, a ...interface{})      {}
func (m silentLogger) Warningf(ctx context.Context, f string, a ...interface{})   {}
func (m silentLogger) Errorf(ctx context.Context, f string, a ...interface{})     {}
func (m silentLogger) Criticalf(ctx context.Context, f string, a ...interface{})  {}
func (m silentLogger) Alertf(ctx context.Context, f string, a ...interface{})     {}
func (m silentLogger) Emergencyf(ctx context.Context, f string, a ...interface{}) {}
func (m silentLogger) Tag(ctx context.Context, key, value string) Logger {return m}
func (m silentLogger) RequestID() string {
	return ""
}

//InitSilentLogger is used for tests to clean up the output
func InitSilentLogger() {
	loggerInstance = silentLogger{}
}
