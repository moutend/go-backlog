package testutil

import (
	"bytes"
	"io/ioutil"
	"net/http"
)

func getIssuesComments() *http.Response {
	return &http.Response{
		StatusCode: http.StatusOK,
		Body: ioutil.NopCloser(bytes.NewBufferString(`[
  {
    "id": 6586,
    "content": "test",
    "changeLog": null,
    "createdUser": {
      "id": 1,
      "userId": "admin",
      "name": "admin",
      "roleType": 1,
      "lang": "ja",
      "mailAddress": "eguchi@nulab.example"
    },
    "created": "2013-08-05T06:15:06Z",
    "updated": "2013-08-05T06:15:06Z",
    "stars": [],
    "notifications": []
  }
]`)),
		Header: make(http.Header),
	}
}

func getPullRequestsComments() *http.Response {
	return &http.Response{
		StatusCode: http.StatusOK,
		Body: ioutil.NopCloser(bytes.NewBufferString(`[
  {
    "id": 35,
    "content": "from api",
    "changeLog": [
      {
        "field": "dependentIssue",
        "newValue": "GIT-3",
        "originalValue": null
      }
    ],
    "createdUser": {
      "id": 1,
      "userId": "admin",
      "name": "admin",
      "roleType": 1,
      "lang": "ja",
      "mailAddress": "eguchi@nulab.example"
    },
    "created": "2015-05-14T01:53:38Z",
    "updated": "2015-05-14T01:53:38Z",
    "stars": [],
    "notifications": []
  }
]`)),
		Header: make(http.Header),
	}
}
