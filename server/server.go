package server

import "context"

type Details struct {
	Port           int64
	AllowedOrigin  string
	RootDir        string
	ConfigFilePath string
	RandomSeed     string
	Interactive    bool
}

type Definition interface {
	Shutdown(ctx context.Context) error
	ListenAndServe() error
}
