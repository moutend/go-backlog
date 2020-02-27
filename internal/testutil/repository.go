package testutil

import (
	"bytes"
	"io/ioutil"
	"net/http"
)

func getRepositories() *http.Response {
	return &http.Response{
		StatusCode: http.StatusOK,
		Body: ioutil.NopCloser(bytes.NewBufferString(`[
  {
    "id": 1,
    "projectId": 151,
    "name": "app",
    "description": "",
    "hookUrl": null,
    "httpUrl": "https://xx.backlogtool.com/git/BLG/app.git",
    "sshUrl": "xx@xx.git.backlogtool.com:/BLG/app.git",
    "displayOrder": 0,
    "pushedAt": null,
    "createdUser": {
      "id": 1,
      "userId": "admin",
      "name": "admin",
      "roleType": 1,
      "lang": "ja",
      "mailAddress": "eguchi@nulab.example"
    },
    "created": "2013-05-30T09:11:36Z",
    "updatedUser": {
      "id": 1,
      "userId": "admin",
      "name": "admin",
      "roleType": 1,
      "lang": "ja",
      "mailAddress": "eguchi@nulab.example"
    },
    "updated": "2013-05-30T09:11:36Z"
  }
]`)),
		Header: make(http.Header),
	}
}

func getRepository() *http.Response {
	return &http.Response{
		StatusCode: http.StatusOK,
		Body: ioutil.NopCloser(bytes.NewBufferString(`{
  "id": 1,
  "projectId": 151,
  "name": "app",
  "description": "",
  "hookUrl": null,
  "httpUrl": "https://xx.backlog.jp/git/BLG/app.git",
  "sshUrl": "xx@xx.git.backlog.jp:/BLG/app.git",
  "displayOrder": 0,
  "pushedAt": null,
  "createdUser": {
    "id": 1,
    "userId": "admin",
    "name": "admin",
    "roleType": 1,
    "lang": "ja",
    "mailAddress": "eguchi@nulab.example"
  },
  "created": "2013-05-30T09:11:36Z",
  "updatedUser": {
    "id": 1,
    "userId": "admin",
    "name": "admin",
    "roleType": 1,
    "lang": "ja",
    "mailAddress": "eguchi@nulab.example"
  },
  "updated": "2013-05-30T09:11:36Z"
}`)),
		Header: make(http.Header),
	}
}
