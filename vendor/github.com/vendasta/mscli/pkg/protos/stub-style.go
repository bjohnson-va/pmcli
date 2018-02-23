package protos

import "fmt"

// Stub Style Enum
type StubStyle string

const (
	StubEmpty   StubStyle = "empty"
	StubHandler StubStyle = "handler"
)

func StubStyleFromString(stubStyle string) (StubStyle, error) {
	switch stubStyle {
	case "empty":
		return StubEmpty, nil
	case "handler":
		return StubHandler, nil
	default:
		return StubEmpty, fmt.Errorf("could not detect stub style, must be one of ['empty', 'handler']")
	}
}
