package client

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"path"

	. "github.com/moutend/go-backlog/pkg/types"
)

const (
	V2UsersPath = "/api/v2/users"
)

// GetUsers returns list of users in your space.
//
// For more details, see the API document.
//
// https://developer.nulab.com/docs/backlog/api/2/get-user-list/#get-user-list
func (c *Client) GetUsers() ([]*User, error) {
	return c.GetUsersContext(context.Background())
}

// GetUsersContext accepts context.
func (c *Client) GetUsersContext(ctx context.Context) ([]*User, error) {
	path, err := c.root.Parse(V2UsersPath)

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

	var users []*User

	if err := json.Unmarshal(body, &users); err != nil {
		return nil, err
	}

	return users, nil
}

// GetMyself returns own information about user.
//
// For more details, see the API document.
//
// https://developer.nulab.com/docs/backlog/api/2/get-own-user/#get-own-user
func (c *Client) GetMyself() (*User, error) {
	return c.GetMyselfContext(context.Background())
}

// GetMyselfContext accepts context.
func (c *Client) GetMyselfContext(ctx context.Context) (*User, error) {
	path, err := c.root.Parse(path.Join(V2UsersPath, "myself"))

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

	var myself *User

	if err := json.Unmarshal(body, &myself); err != nil {
		return nil, err
	}

	return myself, nil
}
