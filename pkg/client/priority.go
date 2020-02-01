package client

import (
	"context"
	"encoding/json"
	"io/ioutil"

	. "github.com/moutend/go-backlog/pkg/types"
)

const (
	V2PrioritiesPath = "/api/v2/priorities"
)

// GetPriorities returns list of priorities.
//
// For more details, see the API document.
//
// https://developer.nulab.com/docs/backlog/api/2/get-priority-list/#get-priority-list
func (c *Client) GetPriorities() ([]*Priority, error) {
	return c.GetPrioritiesContext(context.Background())
}

// GetPrioritiesContext accepts context.
func (c *Client) GetPrioritiesContext(ctx context.Context) ([]*Priority, error) {
	path, err := c.root.Parse(V2PrioritiesPath)

	if err != nil {
		return nil, err
	}

	res, err := c.getContext(ctx, path, nil)

	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		return nil, err
	}

	var priorities []*Priority

	if err := json.Unmarshal(body, &priorities); err != nil {
		return nil, err
	}

	return priorities, nil
}
