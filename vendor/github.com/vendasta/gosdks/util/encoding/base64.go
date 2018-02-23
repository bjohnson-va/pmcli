package encoding

import (
	"strconv"
	"encoding/base64"
	"errors"
)

// EncodeIntegerToBase64 encodes an integer to a base64 string
func EncodeIntegerToBase64(i int64) string {
	s := strconv.FormatInt(i, 10)
	return string(base64.URLEncoding.EncodeToString([]byte(s)))
}

// DecodeBase64ToInt decodes a base64 string to an integer, throwing an error if the string is not encoded as expected.
func DecodeBase64ToInt(s string) (int64, error) {
	if s == "" {
		return 0, nil
	}
	d, err := base64.URLEncoding.DecodeString(s)
	if err != nil {
		return 0, errors.New(DecodingFailure)
	}
	i, err := strconv.ParseInt(string(d), 10, 64)
	if err != nil {
		return 0, errors.New(DecodingFailure)
	}
	return i, nil
}