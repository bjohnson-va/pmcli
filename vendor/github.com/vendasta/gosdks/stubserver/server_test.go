package stubserver

import (
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func readResponse(t *testing.T, resp *http.Response) string {
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		assert.Fail(t, "ReadAll failed")
	}
	return string(body)
}

func Test_ServerHandlesStopCallGracefullyIfServerHasNotStarted(t *testing.T) {
	s := NewServer()
	s.Stop()
}

func Test_ServerHandlesStopCallGracefullyIfServerHasAlreadyBeenStopped(t *testing.T) {
	s := NewServer()
	s.Stop()
	s.Stop()
}

func Test_EmptyStringWillBeReturnedFromURLIfServerHasNotStarted(t *testing.T) {
	s := NewServer()
	assert.Empty(t, s.URL())
}

func Test_StubServerReturnsCorrectResponsesAndStatusCodes(t *testing.T) {
	s := NewServer()
	defer s.Stop()
	s.Start()
	s.PushResponse(&Response{StatusCode: 200, Response: "Resp 1"})
	s.PushResponse(&Response{StatusCode: 401, Response: "Resp 2"})

	type out struct {
		statusCode int
		response   string
	}

	cases := []struct {
		name   string
		output *out
	}{
		{
			name: "First response should be returned",
			output: &out{
				statusCode: 200,
				response:   "Resp 1",
			},
		},
		{
			name: "Second response should be returned",
			output: &out{
				statusCode: 401,
				response:   "Resp 2",
			},
		},
		{
			name: "internal error should be returned if there are no more responses",
			output: &out{
				statusCode: 500,
				response:   "internal error\n",
			},
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			resp, err := http.Get(s.URL())
			assert.Nil(t, err)
			assert.Equal(t, c.output.statusCode, resp.StatusCode)
			body := readResponse(t, resp)
			assert.Equal(t, c.output.response, string(body))
		})
	}
}
