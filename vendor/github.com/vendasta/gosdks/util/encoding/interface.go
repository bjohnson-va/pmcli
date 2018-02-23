// Package encoding provides patterns and implementations for encoding and decoding strings in various formats.
// When you need to encode and/or decode values, give your signatures references to these interfaces
//
// type MyStruct {
//   encoder encoding.IntEncoder
//   decoder encoding.IntDecoder
// }
//
// s := &MyStruct{encoder: encoding.EncodeIntegerToBase64, decoder encoding.DecodeBase64ToInt}
package encoding

const (
	// DecodingFailure happens when the encoded input is not encoded as expected by the decoder.
	// This should be interpreted as an invalid argument-class error.
	DecodingFailure = "Encoded string was invalid, could not be decoded."
)

// IntDecoder attempts to decode an integer from a string, throwing an error if the decoding fails
type IntDecoder func(in string) (int64, error)

// IntEncoder encodes an integer as a string
type IntEncoder func(in int64) string
