// client provides API client for backlog.com.
package client

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"

	. "github.com/moutend/go-backlog/pkg/types"
)

var (
	version   = `v0.6.1`
	userAgent = fmt.Sprintf("go-backlog %v (Fore more details, see https://github.com/moutend/go-backlog)", version)
)

// OptionFunc modifies API client.
type OptionFunc func(*Client)

// OptionHTTPClient sets the given HTTP client.
func OptionHTTPClient(client *http.Client) OptionFunc {
	return func(c *Client) {
		if client == nil {
			return
		}

		c.httpClient = client
	}
}

// Client provides API access.
type Client struct {
	root       *url.URL
	space      string
	token      string
	httpClient *http.Client
}

// New returns API client.
//
// Note: space is your backlog address.
//
// // OK
// client, _ := client.New("example.backlog.com", "xxxxxxxx")
//
// // OK
// client, _ := client.New("https://example.backlog.com", "xxxxxxxx")
func New(space, token string, opts ...OptionFunc) (*Client, error) {
	if space == "" {
		return nil, fmt.Errorf("space is required")
	}
	if token == "" {
		return nil, fmt.Errorf("token is required")
	}
	if strings.HasPrefix(space, "https://") {
		space = strings.TrimPrefix(space, "https://")
	}

	root, err := url.Parse("https://" + space)

	if err != nil {
		return nil, err
	}

	client := &Client{
		root:       root,
		space:      space,
		token:      token,
		httpClient: &http.Client{},
	}

	for _, opt := range opts {
		opt(client)
	}

	return client, nil
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

func (c *Client) postFileContext(ctx context.Context, endpoint *url.URL, query url.Values, path string) (*http.Response, error) {
	file, err := os.Open(path)

	if err != nil {
		return nil, err
	}

	defer file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	defer writer.Close()

	part, err := writer.CreateFormFile(filepath.Base(path), filepath.Base(path))

	if err != nil {
		return nil, err
	}
	if _, err := io.Copy(part, file); err != nil {
		return nil, err
	}

	req, err := c.newRequestContext(ctx, http.MethodPost, endpoint, query, body)

	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", writer.FormDataContentType())

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
