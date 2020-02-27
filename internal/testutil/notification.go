package testutil

import (
	"bytes"
	"io/ioutil"
	"net/http"
)

func getNotifications() *http.Response {
	return &http.Response{
		StatusCode: http.StatusOK,
		Body: ioutil.NopCloser(bytes.NewBufferString(`[
  {
    "id": 299,
    "alreadyRead": true,
    "reason": 2,
    "resourceAlreadyRead": true,
    "project": {
      "id": 2,
      "projectKey": "TEST2",
      "name": "test2",
      "chartEnabled": true,
      "subtaskingEnabled": true,
      "projectLeaderCanEditProjectLeader": false,
      "textFormattingRule": "backlog",
      "archived": false,
      "displayOrder": 0
    },
    "issue": {
      "id": 4531,
      "projectId": 2,
      "issueKey": "TEST2-17",
      "keyId": 17,
      "issueType": {
        "id": 7,
        "projectId": 2,
        "name": "Bug",
        "color": "#990000",
        "displayOrder": 0
      },
      "summary": "aaa",
      "description": "",
      "resolutions": null,
      "priority": {
        "id": 3,
        "name": "Normal"
      },
      "status": {
        "id": 1,
        "projectId": 2,
        "name": "Open",
        "color": "#ed8077",
        "displayOrder": 1000
      },
      "assignee": {
        "id": 2,
        "name": "eguchi",
        "roleType": 2,
        "lang": null,
        "mailAddress": "eguchi@nulab.example"
      },
      "category": [],
      "versions": [],
      "milestone": [],
      "startDate": "2013-08-29T15:00:00Z",
      "dueDate": "2013-09-03T15:00:00Z",
      "estimatedHours": null,
      "actualHours": null,
      "parentIssueId": null,
      "createdUser": {
        "id": 1,
        "userId": "admin",
        "name": "admin",
        "roleType": 1,
        "lang": "ja",
        "mailAddress": "eguchi@nulab.example"
      },
      "created": "2013-04-23T07:38:59Z",
      "updatedUser": {
        "id": 1,
        "userId": "admin",
        "name": "admin",
        "roleType": 1,
        "lang": "ja",
        "mailAddress": "eguchi@nulab.example"
      },
      "updated": "2013-09-06T09:25:41Z",
      "customFields": [],
      "attachments": [],
      "sharedFiles": [],
      "stars": []
    },
    "comment": {
      "id": 7007,
      "content": "hoge",
      "changeLog": null,
      "createdUser": {
        "id": 2,
        "userId": "eguchi",
        "name": "eguchi",
        "roleType": 2,
        "lang": null,
        "mailAddress": "eguchi@nulab.example"
      },
      "created": "2013-10-31T06:58:58Z",
      "updated": "2013-10-31T06:58:58Z",
      "stars": [],
      "notifications": []
    },
    "pullRequest": null,
    "pullRequestComment": null,
    "sender": {
      "id": 2,
      "userId": "eguchi",
      "name": "eguchi",
      "roleType": 2,
      "lang": null,
      "mailAddress": "eguchi@nulab.example"
    },
    "created": "2013-10-31T06:58:59Z"
  }
]`)),
		Header: make(http.Header),
	}
}
