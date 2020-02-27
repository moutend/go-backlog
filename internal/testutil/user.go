package testutil

import (
	"bytes"
	"io/ioutil"
	"net/http"
)

func getUsers() *http.Response {
	return &http.Response{
		StatusCode: http.StatusOK,
		Body: ioutil.NopCloser(bytes.NewBufferString(`[
  {
    "id": 1,
    "userId": "admin",
    "name": "admin",
    "roleType": 1,
    "lang": "ja",
    "mailAddress": "eguchi@nulab.example"
  }
]`)),
		Header: make(http.Header),
	}
}

func getMyself() *http.Response {
	return &http.Response{
		StatusCode: http.StatusOK,
		Body: ioutil.NopCloser(bytes.NewBufferString(`{
  "id": 1,
  "userId": "admin",
  "name": "admin",
  "roleType": 1,
  "lang": "ja",
  "mailAddress": "eguchi@nulab.example"
}`)),
		Header: make(http.Header),
	}
}
