package testingutils

import (
	"io"
	"net/http"
	"net/http/httptest"
)

// ResponseFromString builds an HTTP response whose body is the given string
func ResponseFromString(body string) *http.Response {
	handler := func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, body)
	}
	req := httptest.NewRequest("GET", "http://example.com/foo", nil)
	w := httptest.NewRecorder()
	handler(w, req)
	return w.Result()
}
