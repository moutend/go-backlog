package client

import (
	"context"
	"encoding/json"
	"io/ioutil"

	. "github.com/moutend/go-backlog/pkg/types"
)

const (
	V2SpacePath             = "/api/v2/space"
	V2SpaceActivitiesPath   = "/api/v2/space/activities"
	V2SpaceImagePath        = "/api/v2/space/image"
	V2SpaceNotificationPath = "/api/v2/space/notification"
	V2SpaceDiskUsagePath    = "/api/v2/space/diskUsage"
	V2SpaceAttachmentPath   = "/api/v2/space/attachment"
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
