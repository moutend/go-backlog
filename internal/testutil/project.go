package testutil

import (
	"bytes"
	"io/ioutil"
	"net/http"
)

func getProject12345Statuses() *http.Response {
	return &http.Response{
		StatusCode: http.StatusOK,
		Body: ioutil.NopCloser(bytes.NewBufferString(`[
  {
    "id": 1,
    "projectId": 1,
    "name": "Open",
    "color": "#ed8077",
    "displayOrder": 1000
  }
]`)),
		Header: make(http.Header),
	}
}

func getProject12345() *http.Response {
	return &http.Response{
		StatusCode: http.StatusOK,
		Body: ioutil.NopCloser(bytes.NewBufferString(`{
  "id": 1,
  "projectKey": "TEST",
  "name": "test",
  "chartEnabled": false,
  "subtaskingEnabled": false,
  "projectLeaderCanEditProjectLeader": false,
  "textFormattingRule": "markdown",
  "archived": false
}`)),
		Header: make(http.Header),
	}
}

func getProjects() *http.Response {
	return &http.Response{
		StatusCode: http.StatusOK,
		Body: ioutil.NopCloser(bytes.NewBufferString(`[
  {
    "id": 1,
    "projectKey": "TEST",
    "name": "test",
    "chartEnabled": false,
    "subtaskingEnabled": false,
    "projectLeaderCanEditProjectLeader": false,
    "textFormattingRule": "markdown",
    "archived": false
  }
]`)),
		Header: make(http.Header),
	}
}
