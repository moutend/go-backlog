package testutil

import (
	"bytes"
	"io/ioutil"
	"net/http"
)

func postIssues() *http.Response {
	return &http.Response{
		StatusCode: http.StatusOK,
		Body: ioutil.NopCloser(bytes.NewBufferString(`{
  "id": 1,
  "projectId": 1,
  "issueKey": "BLG-1",
  "keyId": 1,
  "issueType": {
    "id": 2,
    "projectId": 1,
    "name": "Task",
    "color": "#7ea800",
    "displayOrder": 0
  },
  "summary": "first issue",
  "description": "",
  "resolutions": null,
  "priority": {
    "id": 3,
    "name": "Normal"
  },
  "status": {
    "id": 1,
    "projectId": 1,
    "name": "Open",
    "color": "#ed8077",
    "displayOrder": 1000
  },
  "assignee": {
    "id": 2,
    "userId": "eguchi",
    "name": "eguchi",
    "roleType": 2,
    "lang": null,
    "mailAddress": "eguchi@nulab.example"
  },
  "category": [],
  "versions": [],
  "milestone": [
    {
      "id": 30,
      "projectId": 1,
      "name": "wait for release",
      "description": "",
      "startDate": null,
      "releaseDueDate": null,
      "archived": false,
      "displayOrder": 0
    }
  ],
  "startDate": null,
  "dueDate": null,
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
  "created": "2012-07-23T06:10:15Z",
  "updatedUser": {
    "id": 1,
    "userId": "admin",
    "name": "admin",
    "roleType": 1,
    "lang": "ja",
    "mailAddress": "eguchi@nulab.example"
  },
  "updated": "2012-07-23T06:10:15Z",
  "customFields": [],
  "attachments": [
    {
      "id": 1,
      "name": "IMGP0088.JPG",
      "size": 85079
    }
  ],
  "sharedFiles": [],
  "stars": []
}`)),
		Header: make(http.Header),
	}
}

func getIssues12345() *http.Response {
	return &http.Response{
		StatusCode: http.StatusOK,
		Body: ioutil.NopCloser(bytes.NewBufferString(`{
  "id": 1,
  "projectId": 1,
  "issueKey": "BLG-1",
  "keyId": 1,
  "issueType": {
    "id": 2,
    "projectId": 1,
    "name": "Task",
    "color": "#7ea800",
    "displayOrder": 0
  },
  "summary": "first issue",
  "description": "",
  "resolutions": null,
  "priority": {
    "id": 3,
    "name": "Normal"
  },
  "status": {
    "id": 1,
    "projectId": 1,
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
  "milestone": [
    {
      "id": 30,
      "projectId": 1,
      "name": "wait for release",
      "description": "",
      "startDate": null,
      "releaseDueDate": null,
      "archived": false,
      "displayOrder": 0
    }
  ],
  "startDate": null,
  "dueDate": null,
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
  "created": "2012-07-23T06:10:15Z",
  "updatedUser": {
    "id": 1,
    "userId": "admin",
    "name": "admin",
    "roleType": 1,
    "lang": "ja",
    "mailAddress": "eguchi@nulab.example"
  },
  "updated": "2013-02-07T08:09:49Z",
  "customFields": [],
  "attachments": [
    {
      "id": 1,
      "name": "IMGP0088.JPG",
      "size": 85079
    }
  ],
  "sharedFiles": [],
  "stars": [
    {
      "id": 10,
      "comment": null,
      "url": "https://xx.backlogtool.com/view/BLG-1",
      "title": "[BLG-1] first issue | Show issue - Backlog",
      "presenter": {
        "id": 2,
        "userId": "eguchi",
        "name": "eguchi",
        "roleType": 2,
        "lang": "ja",
        "mailAddress": "eguchi@nulab.example"
      },
      "created": "2013-07-08T10:24:28Z"
    }
  ]
}`)),
		Header: make(http.Header),
	}
}

func patchIssues12345() *http.Response {
	return &http.Response{
		StatusCode: http.StatusOK,
		Body: ioutil.NopCloser(bytes.NewBufferString(`{
  "id": 1,
  "projectId": 1,
  "issueKey": "BLG-1",
  "keyId": 1,
  "issueType": {
    "id": 2,
    "projectId": 1,
    "name": "Task",
    "color": "#7ea800",
    "displayOrder": 0
  },
  "summary": "first issue",
  "description": "",
  "resolutions": null,
  "priority": {
    "id": 3,
    "name": "Normal"
  },
  "status": {
    "id": 1,
    "projectId": 1,
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
  "milestone": [
    {
      "id": 30,
      "projectId": 1,
      "name": "wait for release",
      "description": "",
      "startDate": null,
      "releaseDueDate": null,
      "archived": false,
      "displayOrder": 0
    }
  ],
  "startDate": null,
  "dueDate": null,
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
  "created": "2012-07-23T06:10:15Z",
  "updatedUser": {
    "id": 1,
    "userId": "admin",
    "name": "admin",
    "roleType": 1,
    "lang": "ja",
    "mailAddress": "eguchi@nulab.example"
  },
  "updated": "2013-02-07T08:09:49Z",
  "customFields": [],
  "attachments": [
    {
      "id": 1,
      "name": "IMGP0088.JPG",
      "size": 85079
    }
  ],
  "sharedFiles": [],
  "stars": [
    {
      "id": 10,
      "comment": null,
      "url": "https://xx.backlogtool.com/view/BLG-1",
      "title": "[BLG-1] first issue | Show issue - Backlog",
      "presenter": {
        "id": 2,
        "userId": "eguchi",
        "name": "eguchi",
        "roleType": 2,
        "lang": "ja",
        "mailAddress": "eguchi@nulab.example"
      },
      "created": "2013-07-08T10:24:28Z"
    }
  ]
}`)),
		Header: make(http.Header),
	}
}

func deleteIssues12345() *http.Response {
	return &http.Response{
		StatusCode: http.StatusOK,
		Body: ioutil.NopCloser(bytes.NewBufferString(`{
  "id": 1,
  "projectId": 1,
  "issueKey": "BLG-1",
  "keyId": 1,
  "issueType": {
    "id": 2,
    "projectId": 1,
    "name": "Task",
    "color": "#7ea800",
    "displayOrder": 0
  },
  "summary": "first issue",
  "description": "",
  "resolutions": null,
  "priority": {
    "id": 3,
    "name": "Normal"
  },
  "status": {
    "id": 1,
    "projectId": 1,
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
  "milestone": [
    {
      "id": 30,
      "projectId": 1,
      "name": "wait for release",
      "description": "",
      "startDate": null,
      "releaseDueDate": null,
      "archived": false,
      "displayOrder": 0
    }
  ],
  "startDate": null,
  "dueDate": null,
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
  "created": "2012-07-23T06:10:15Z",
  "updatedUser": {
    "id": 1,
    "userId": "admin",
    "name": "admin",
    "roleType": 1,
    "lang": "ja",
    "mailAddress": "eguchi@nulab.example"
  },
  "updated": "2013-02-07T08:09:49Z",
  "customFields": [],
  "attachments": [
    {
      "id": 1,
      "name": "IMGP0088.JPG",
      "size": 85079
    }
  ],
  "sharedFiles": [],
  "stars": [
    {
      "id": 10,
      "comment": null,
      "url": "https://xx.backlogtool.com/view/BLG-1",
      "title": "[BLG-1] first issue | Show issue - Backlog",
      "presenter": {
        "id": 2,
        "userId": "eguchi",
        "name": "eguchi",
        "roleType": 2,
        "lang": "ja",
        "mailAddress": "eguchi@nulab.example"
      },
      "created": "2013-07-08T10:24:28Z"
    }
  ]
}`)),
		Header: make(http.Header),
	}
}

func getIssueTypes() *http.Response {
	return &http.Response{
		StatusCode: http.StatusOK,
		Body: ioutil.NopCloser(bytes.NewBufferString(`[
  {
    "id": 1,
    "projectId": 1,
    "name": "Bug",
    "color": "#990000",
    "displayOrder": 0,
    "templateSummary": "Subject",
    "templateDescription": "Description"
  }
]`)),
		Header: make(http.Header),
	}
}

func getIssuesCount() *http.Response {
	return &http.Response{
		StatusCode: http.StatusOK,
		Body: ioutil.NopCloser(bytes.NewBufferString(`{
  "count": 43
}`)),
		Header: make(http.Header),
	}
}

func getIssues() *http.Response {
	return &http.Response{
		StatusCode: http.StatusOK,
		Body: ioutil.NopCloser(bytes.NewBufferString(`[
  {
    "id": 1,
    "projectId": 1,
    "issueKey": "BLG-1",
    "keyId": 1,
    "issueType": {
      "id": 2,
      "projectId": 1,
      "name": "Task",
      "color": "#7ea800",
      "displayOrder": 0
    },
    "summary": "first issue",
    "description": "",
    "resolutions": null,
    "priority": {
      "id": 3,
      "name": "Normal"
    },
    "status": {
      "id": 1,
      "projectId": 1,
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
    "milestone": [
      {
        "id": 30,
        "projectId": 1,
        "name": "wait for release",
        "description": "",
        "startDate": null,
        "releaseDueDate": null,
        "archived": false,
        "displayOrder": 0
      }
    ],
    "startDate": null,
    "dueDate": null,
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
    "created": "2012-07-23T06:10:15Z",
    "updatedUser": {
      "id": 1,
      "userId": "admin",
      "name": "admin",
      "roleType": 1,
      "lang": "ja",
      "mailAddress": "eguchi@nulab.example"
    },
    "updated": "2013-02-07T08:09:49Z",
    "customFields": [],
    "attachments": [
      {
        "id": 1,
        "name": "IMGP0088.JPG",
        "size": 85079
      }
    ],
    "sharedFiles": [
      {
        "id": 454403,
        "type": "file",
        "dir": "/userIcon/",
        "name": "01_male clerk.png",
        "size": 2735,
        "createdUser": {
          "id": 5686,
          "userId": "takada",
          "name": "takada",
          "roleType": 2,
          "lang": "ja",
          "mailAddress": "takada@nulab.example"
        },
        "created": "2009-02-27T03:26:15Z",
        "updatedUser": {
          "id": 5686,
          "userId": "takada",
          "name": "takada",
          "roleType": 2,
          "lang": "ja",
          "mailAddress": "takada@nulab.example"
        },
        "updated": "2009-03-03T16:57:47Z"
      }
    ],
    "stars": [
      {
        "id": 10,
        "comment": null,
        "url": "https://xx.backlogtool.com/view/BLG-1",
        "title": "[BLG-1] first issue | Show issue - Backlog",
        "presenter": {
          "id": 2,
          "userId": "eguchi",
          "name": "eguchi",
          "roleType": 2,
          "lang": "ja",
          "mailAddress": "eguchi@nulab.example"
        },
        "created": "2013-07-08T10:24:28Z"
      }
    ]
  }
]`)),
		Header: make(http.Header),
	}
}
