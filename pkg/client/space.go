package client

import (
	"context"
	"encoding/json"
	"io/ioutil"

	. "github.com/moutend/go-backlog/pkg/types"
)

// GetSpace returns information about your space.
//
// For more details, see the API document.
//
// https://developer.nulab.com/docs/backlog/api/2/get-space/#get-space
func (c *Client) GetSpace() (*Space, error) {
	return c.GetSpaceContext(context.Background())
}

// GetSpaceContext accepts context.
func (c *Client) GetSpaceContext(ctx context.Context) (*Space, error) {
	path, err := c.root.Parse(V2SpacePath)

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

	var space *Space

	if err := json.Unmarshal(body, &space); err != nil {
		return nil, err
	}

	return space, nil
}

// GetSpaceDiskUsage returns information about space disk usage.
//
// For more details, see the API document.
//
// https://developer.nulab.com/docs/backlog/api/2/get-space-disk-usage/#get-space-disk-usage
func (c *Client) GetSpaceDiskUsage() (*TotalDiskUsage, error) {
	return c.GetSpaceDiskUsageContext(context.Background())
}

// GetSpaceDiskUsageContext accepts context.
func (c *Client) GetSpaceDiskUsageContext(ctx context.Context) (*TotalDiskUsage, error) {
	path, err := c.root.Parse(V2SpaceDiskUsagePath)

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

	var totalDiskUsage *TotalDiskUsage

	if err := json.Unmarshal(body, &totalDiskUsage); err != nil {
		return nil, err
	}

	return totalDiskUsage, nil
}
