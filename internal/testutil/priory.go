package testutil

import (
	"bytes"
	"io/ioutil"
	"net/http"
)

func getPriorities() *http.Response {
	return &http.Response{
		StatusCode: http.StatusOK,
		Body: ioutil.NopCloser(bytes.NewBufferString(`[
  {
    "id": 2,
    "name": "High"
  },
  {
    "id": 3,
    "name": "Normal"
  },
  {
    "id": 4,
    "name": "Low"
  }
]`)),
		Header: make(http.Header),
	}
}
