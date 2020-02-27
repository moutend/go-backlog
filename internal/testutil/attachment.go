package testutil

import (
	"bytes"
	"io/ioutil"
	"net/http"
)

func init() {
	responseMap[http.MethodPost+" "+"/api/v2/space/attachment"] = &http.Response{
		StatusCode: http.StatusOK,
		Body: ioutil.NopCloser(bytes.NewBufferString(`{
  "id": 1,
  "name": "test.txt",
  "size": 8857
}`)),
		Header: make(http.Header),
	}
}
