package testutil

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"testing"
)

type RoundTripFunc func(req *http.Request) *http.Response

func (fn RoundTripFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	return fn(req), nil
}

func NewTestClient(fn RoundTripFunc) *http.Client {
	return &http.Client{
		Transport: RoundTripFunc(fn),
	}
}

func NewFakeClient(t *testing.T) *http.Client {
	t.Helper()

	return NewTestClient(func(req *http.Request) *http.Response {
		t.Logf("%s %s\n", req.Method, req.URL.Path)

		switch req.Method + " " + req.URL.Path {
		case "POST /api/v2/space/attachment":
			return postSpaceAttachment()
		case "GET /api/v2/wikis":
			return getWikis()
		case "POST /api/v2/wikis":
			return postWikis()
		case "GET /api/v2/wikis/12345":
			return getWikis12345()
		case "PATCH /api/v2/wikis/12345":
			return patchWikis12345()
		case "DELETE /api/v2/wikis/12345":
			return deleteWikis12345()
		case "GET /api/v2/wikis/count":
			return getWikisCount()
		case "GET /api/v2/wikis/tags":
			return getWikisTags()
		case "GET /api/v2/wikis/12345/attachments":
			return getWikis12345Attachments()
		case "DELETE /api/v2/wikis/12345/attachments/67890":
			return deleteWikis12345Attachments67890()
		case "POST /api/v2/wikis/12345/attachments":
			return postWikis12345Attachments()
		case "GET /api/v2/wikis/12345/attachments/67890":
			return getWikis12345Attachments67890()
		case "GET /api/v2/issues/12345/comments":
			return getIssuesComments()
		case "GET /api/v2/projects/12345/git/repositories/67890/pullRequests/123/comments":
			return getPullRequestsComments()
		case "POST /api/v2/issues":
			return postIssues()
		case "GET /api/v2/issues/12345":
			return getIssues12345()
		case "PATCH /api/v2/issues/12345":
			return patchIssues12345()
		case "DELETE /api/v2/issues/12345":
			return deleteIssues12345()
		case "GET /api/v2/projects/12345/issueTypes":
			return getIssueTypes()
		case "GET /api/v2/issues/count":
			return getIssuesCount()
		case "GET /api/v2/issues":
			return getIssues()
		case "GET /api/v2/notifications":
			return getNotifications()
		case "GET /api/v2/priorities":
			return getPriorities()
		case "GET /api/v2/projects/12345/statuses":
			return getProject12345Statuses()
		case "GET /api/v2/projects/12345":
			return getProject12345()
		case "GET /api/v2/projects":
			return getProjects()
		case "POST /api/v2/projects/12345/git/repositories/67890/pullRequests":
			return postPullRequest()
		case "GET /api/v2/projects/12345/git/repositories/67890/pullRequests/123":
			return getPullRequest()
		case "PATCH /api/v2/projects/12345/git/repositories/67890/pullRequests/123":
			return patchPullRequest()
		case "GET /api/v2/projects/12345/git/repositories/67890/pullRequests/count":
			return getPullRequestsCount()
		case "GET /api/v2/projects/12345/git/repositories/67890/pullRequests":
			return getPullRequests()
		default:
			break
		}

		return defaultResponse()
	})
}

func defaultResponse() *http.Response {
	return &http.Response{
		StatusCode: http.StatusBadRequest,
		Body: ioutil.NopCloser(bytes.NewBufferString(`{
  "errors": [
    {
      "message": "",
      "code": 999,
      "moreInfo": ""
    }
  ]
}`)),
		Header: make(http.Header),
	}
}
