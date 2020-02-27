package testutil

import (
	"bytes"
	"io/ioutil"
	"net/http"
)

func getSpace() *http.Response {
	return &http.Response{
		StatusCode: http.StatusOK,
		Body: ioutil.NopCloser(bytes.NewBufferString(`{
  "spaceKey": "nulab",
  "name": "Nulab Inc.",
  "ownerId": 1,
  "lang": "ja",
  "timezone": "Asia/Tokyo",
  "reportSendTime": "08:00:00",
  "textFormattingRule": "markdown",
  "created": "2008-07-06T15:00:00Z",
  "updated": "2013-06-18T07:55:37Z"
}`)),
		Header: make(http.Header),
	}
}
