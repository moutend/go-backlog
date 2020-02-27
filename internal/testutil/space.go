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

func getSpaceDiskUsage() *http.Response {
	return &http.Response{
		StatusCode: http.StatusOK,
		Body: ioutil.NopCloser(bytes.NewBufferString(`{
  "capacity": 1073741824,
  "issue": 119511,
  "wiki": 48575,
  "file": 0,
  "subversion": 0,
  "git": 0,
  "gitLFS": 0,
  "details": [
    {
      "projectId": 1,
      "issue": 11931,
      "wiki": 0,
      "file": 0,
      "subversion": 0,
      "git": 0,
      "gitLFS": 0
    }
  ]
}`)),
		Header: make(http.Header),
	}
}
