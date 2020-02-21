package client

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"

	. "github.com/moutend/go-backlog/pkg/types"
)

var (
	version   = `v0.4.0`
	userAgent = fmt.Sprintf("go-backlog %v (Fore more details, see https://github.com/moutend/go-backlog)", version)
)

// HTTPClient is an interface which performs http.Client.Do method.
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

// Client provides API access.
type Client struct {
	root       *url.URL
	token      string
	httpClient HTTPClient
}

func New(space, token string) (*Client, error) {
	if space == "" {
		return nil, fmt.Errorf("space is empty")
	}
	if token == "" {
		return nil, fmt.Errorf("token is empty")
	}

	root, err := url.Parse("https://" + space)

	if err != nil {
		return nil, err
	}

	client := &Client{
		root:       root,
		token:      token,
		httpClient: &http.Client{},
	}

	return client, nil
}

// SetHTTPClient sets default HTTP client.
func (c *Client) SetHTTPClient(hc HTTPClient) {
	c.httpClient = hc
}

func (c *Client) newRequestContext(ctx context.Context, method string, endpoint *url.URL, query url.Values, payload io.Reader) (*http.Request, error) {
	q := url.Values{}

	q.Add("apiKey", c.token)

	for k, vs := range query {
		for _, v := range vs {
			q.Add(k, v)
		}
	}

	endpoint.RawQuery = q.Encode()

	req, err := http.NewRequest(method, endpoint.String(), payload)

	if err != nil {
	}

	req.Header.Add(`User-Agent`, userAgent)

	req = req.WithContext(ctx)

	return req, nil
}

func validateResponse(res *http.Response) error {
	if res.StatusCode < http.StatusBadRequest {
		return nil
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		return err
	}

	var e Errors

	if err := json.Unmarshal(body, &e); err != nil {
		return err
	}

	return e
}

func (c *Client) getContext(ctx context.Context, endpoint *url.URL, query url.Values) (*http.Response, error) {
	req, err := c.newRequestContext(ctx, http.MethodGet, endpoint, query, nil)

	if err != nil {
		return nil, err
	}

	res, err := c.httpClient.Do(req)

	if err != nil {
		return nil, err
	}
	if err := validateResponse(res); err != nil {
		return nil, err
	}

	return res, nil
}

func (c *Client) patchContext(ctx context.Context, endpoint *url.URL, query url.Values, payload io.Reader) (*http.Response, error) {
	req, err := c.newRequestContext(ctx, http.MethodPatch, endpoint, query, payload)

	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	res, err := c.httpClient.Do(req)

	if err != nil {
		return nil, err
	}
	if err := validateResponse(res); err != nil {
		return nil, err
	}

	return res, nil
}
func (c *Client) postContext(ctx context.Context, endpoint *url.URL, query url.Values, payload io.Reader) (*http.Response, error) {
	req, err := c.newRequestContext(ctx, http.MethodPost, endpoint, query, payload)

	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	res, err := c.httpClient.Do(req)

	if err != nil {
		return nil, err
	}
	if err := validateResponse(res); err != nil {
		return nil, err
	}

	return res, nil
}

func (c *Client) deleteContext(ctx context.Context, endpoint *url.URL, query url.Values) (*http.Response, error) {
	req, err := c.newRequestContext(ctx, http.MethodDelete, endpoint, query, nil)

	if err != nil {
		return nil, err
	}

	res, err := c.httpClient.Do(req)

	if err != nil {
		return nil, err
	}
	if err := validateResponse(res); err != nil {
		return nil, err
	}

	return res, nil
}
