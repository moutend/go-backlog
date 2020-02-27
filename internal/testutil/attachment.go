package testutil

import (
	"bytes"
	"io/ioutil"
	"net/http"
)

func postSpaceAttachment() *http.Response {
	// http.MethodPost+" "+"/api/v2/space/attachment"
	return &http.Response{
		StatusCode: http.StatusOK,
		Body: ioutil.NopCloser(bytes.NewBufferString(`{
  "id": 1,
  "name": "test.txt",
  "size": 8857
}`)),
		Header: make(http.Header),
	}
}
