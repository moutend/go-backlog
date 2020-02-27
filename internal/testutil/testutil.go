package testutil

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"testing"
)

var (
	defaultResponse *http.Response = nil
	responseMap                    = make(map[string]*http.Response)
)

type RoundTripFunc func(req *http.Request) *http.Response

func (fn RoundTripFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	return fn(req), nil
}

func NewTestClient(fn RoundTripFunc) *http.Client {
	return &http.Client{
		Transport: RoundTripFunc(fn),
	}
}

func NewFakeClient(t *testing.T) *http.Client {
	t.Helper()

	return NewTestClient(func(req *http.Request) *http.Response {
		t.Logf("%s %s\n", req.Method, req.URL.Path)

		if responseMap[req.Method+" "+req.URL.Path] != nil {
			return responseMap[req.Method+" "+req.URL.Path]
		}

		return defaultResponse
	})
}

func init() {
	defaultResponse = &http.Response{
		StatusCode: http.StatusBadRequest,
		Body: ioutil.NopCloser(bytes.NewBufferString(`{
  "errors": [
    {
      "message": "",
      "code": 999,
      "moreInfo": ""
    }
  ]
}`)),
		Header: make(http.Header),
	}
}
