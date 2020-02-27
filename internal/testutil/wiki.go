package testutil

import (
	"bytes"
	"io/ioutil"
	"net/http"
)

func getWikis() *http.Response {
	// http.MethodGet+" "+"/api/v2/wikis"
	return &http.Response{
		StatusCode: http.StatusOK,
		Body: ioutil.NopCloser(bytes.NewBufferString(`[
  {
    "id": 112,
    "projectId": 103,
    "name": "Home",
    "tags": [
      {
        "id": 12,
        "name": "proceedings"
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

func postWikis() *http.Response {
	// http.MethodPost+" "+"/api/v2/wikis"
	return &http.Response{
		StatusCode: http.StatusCreated,
		Body: ioutil.NopCloser(bytes.NewBufferString(`{
  "id": 1,
  "projectId": 1,
  "name": "Home",
  "content": "test",
  "tags": [
    {
      "id": 12,
      "name": "proceedings"
    }
  ],
  "attachments": [],
  "sharedFiles": [],
  "stars": [],
  "createdUser": {
    "id": 1,
    "userId": "admin",
    "name": "admin",
    "roleType": 1,
    "lang": "ja",
    "mailAddress": "eguchi@nulab.example"
  },
  "created": "2012-07-23T06:09:48Z",
  "updatedUser": {
    "id": 1,
    "userId": "admin",
    "name": "admin",
    "roleType": 1,
    "lang": "ja",
    "mailAddress": "eguchi@nulab.example"
  },
  "updated": "2012-07-23T06:09:48Z"
}`)),
		Header: make(http.Header),
	}
}

func getWikis12345() *http.Response {
	// http.MethodGet+" "+"/api/v2/wikis/12345"
	return &http.Response{
		StatusCode: http.StatusOK,
		Body: ioutil.NopCloser(bytes.NewBufferString(`{
  "id": 1,
  "projectId": 1,
  "name": "Home",
  "content": "test",
  "tags": [
    {
      "id": 12,
      "name": "proceedings"
    }
  ],
  "attachments": [
    {
      "id": 1,
      "name": "test.json",
      "size": 8857,
      "createdUser": {
        "id": 1,
        "userId": "admin",
        "name": "admin",
        "roleType": 1,
        "lang": "ja",
        "mailAddress": "eguchi@nulab.example"
      },
      "created": "2014-01-06T11:10:45Z"
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
  "stars": [],
  "createdUser": {
    "id": 1,
    "userId": "admin",
    "name": "admin",
    "roleType": 1,
    "lang": "ja",
    "mailAddress": "eguchi@nulab.example"
  },
  "created": "2012-07-23T06:09:48Z",
  "updatedUser": {
    "id": 1,
    "userId": "admin",
    "name": "admin",
    "roleType": 1,
    "lang": "ja",
    "mailAddress": "eguchi@nulab.example"
  },
  "updated": "2012-07-23T06:09:48Z"
}`)),
		Header: make(http.Header),
	}
}

func patchWikis12345() *http.Response {
	// http.MethodPatch+" "+"/api/v2/wikis/12345"
	return &http.Response{
		StatusCode: http.StatusOK,
		Body: ioutil.NopCloser(bytes.NewBufferString(`{
  "id": 1,
  "projectId": 1,
  "name": "Home",
  "content": "test",
  "tags": [
    {
      "id": 12,
      "name": "proceedings"
    }
  ],
  "attachments": [
    {
      "id": 1,
      "name": "test.json",
      "size": 8857,
      "createdUser": {
        "id": 1,
        "name": "admin",
        "roleType": 1,
        "lang": "ja",
        "mailAddress": "eguchi@nulab.example"
      },
      "created": "2014-01-06T11:10:45Z"
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
        "name": "takada",
        "roleType": 2,
        "lang": "ja",
        "mailAddress": "takada@nulab.example"
      },
      "created": "2009-02-27T03:26:15Z",
      "updatedUser": {
        "id": 5686,
        "name": "takada",
        "roleType": 2,
        "lang": "ja",
        "mailAddress": "takada@nulab.example"
      },
      "updated": "2009-03-03T16:57:47Z"
    }
  ],
  "stars": [],
  "createdUser": {
    "id": 1,
    "userId": "admin",
    "name": "admin",
    "roleType": 1,
    "lang": "ja",
    "mailAddress": "eguchi@nulab.example"
  },
  "created": "2012-07-23T06:09:48Z",
  "updatedUser": {
    "id": 1,
    "userId": "admin",
    "name": "admin",
    "roleType": 1,
    "lang": "ja",
    "mailAddress": "eguchi@nulab.example"
  },
  "updated": "2012-07-23T06:09:48Z"
}`)),
		Header: make(http.Header),
	}
}

func deleteWikis12345() *http.Response {
	// http.MethodDelete+" "+"/api/v2/wikis/12345"
	return &http.Response{
		StatusCode: http.StatusOK,
		Body: ioutil.NopCloser(bytes.NewBufferString(`{
  "id": 1,
  "projectId": 1,
  "name": "Home",
  "content": "test",
  "tags": [
    {
      "id": 12,
      "name": "proceedings"
    }
  ],
  "attachments": [
    {
      "id": 1,
      "name": "test.json",
      "size": 8857,
      "createdUser": {
        "id": 1,
        "userId": "admin",
        "name": "admin",
        "roleType": 1,
        "lang": "ja",
        "mailAddress": "eguchi@nulab.example"
      },
      "created": "2014-01-06T11:10:45Z"
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
  "stars": [],
  "createdUser": {
    "id": 1,
    "userId": "admin",
    "name": "admin",
    "roleType": 1,
    "lang": "ja",
    "mailAddress": "eguchi@nulab.example"
  },
  "created": "2012-07-23T06:09:48Z",
  "updatedUser": {
    "id": 1,
    "userId": "admin",
    "name": "admin",
    "roleType": 1,
    "lang": "ja",
    "mailAddress": "eguchi@nulab.example"
  },
  "updated": "2012-07-23T06:09:48Z"
}`)),
		Header: make(http.Header),
	}
}

func getWikisCount() *http.Response {
	// http.MethodGet+" "+"/api/v2/wikis/count"
	return &http.Response{
		StatusCode: http.StatusOK,
		Body: ioutil.NopCloser(bytes.NewBufferString(`{
  "count": 10
}`)),
		Header: make(http.Header),
	}
}

func getWikisTags() *http.Response {
	// http.MethodGet+" "+"/api/v2/wikis/tags"
	return &http.Response{
		StatusCode: http.StatusOK,
		Body: ioutil.NopCloser(bytes.NewBufferString(`[
  {
    "id": 1,
    "name": "test"
  }
]`)),
		Header: make(http.Header),
	}
}

func getWikis12345Attachments() *http.Response {
	// http.MethodGet+" "+"/api/v2/wikis/12345/attachments"
	return &http.Response{
		StatusCode: http.StatusOK,
		Body: ioutil.NopCloser(bytes.NewBufferString(`[
  {
    "id": 1,
    "name": "IMGP0088.JPG",
    "size": 85079
  }
]`)),
		Header: make(http.Header),
	}
}

func deleteWikis12345Attachments67890() *http.Response {
	// http.MethodDelete+" "+"/api/v2/wikis/12345/attachments/67890"
	return &http.Response{
		StatusCode: http.StatusOK,
		Body: ioutil.NopCloser(bytes.NewBufferString(`{
  "id": 2,
  "name": "Duke.png",
  "size": 196186,
  "createdUser": {
    "id": 1,
    "userId": "admin",
    "name": "admin",
    "roleType": 1,
    "lang": null,
    "mailAddress": "eguchi@nulab.example"
  },
  "created": "2014-07-11T06:26:05Z"
}`)),
		Header: make(http.Header),
	}
}

func postWikis12345Attachments() *http.Response {
	// http.MethodPost+" "+"/api/v2/wikis/12345/attachments"
	return &http.Response{
		StatusCode: http.StatusCreated,
		Body: ioutil.NopCloser(bytes.NewBufferString(`[
  {
    "id": 2,
    "name": "Duke.png",
    "size": 196186,
    "createdUser": {
      "id": 1,
      "userId": "admin",
      "name": "admin",
      "roleType": 1,
      "lang": null,
      "mailAddress": "eguchi@nulab.example"
    },
    "created": "2014-07-11T06:26:05Z"
  }
]`)),
		Header: make(http.Header),
	}
}
