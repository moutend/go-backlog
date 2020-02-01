package client

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/url"

	. "github.com/moutend/go-backlog/pkg/types"
)

const (
	V2NotificationsPath = "/api/v2/notifications"
)

// GetNotifications returns own notifications.
//
// For more details, see the API document.
//
// https://developer.nulab.com/docs/backlog/api/2/get-notification/#get-notification
func (c *Client) GetNotifications(query url.Values) ([]*Notification, error) {
	return c.GetNotificationsContext(context.Background(), query)
}

// GetNotificationsContext accepts context.
func (c *Client) GetNotificationsContext(ctx context.Context, query url.Values) ([]*Notification, error) {
	path, err := c.root.Parse(V2NotificationsPath)

	if err != nil {
		return nil, err
	}

	res, err := c.getContext(ctx, path, query)

	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		return nil, err
	}

	var notifications []*Notification

	if err := json.Unmarshal(body, &notifications); err != nil {
		return nil, err
	}

	return notifications, nil
}
