package client

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/url"
	"path"

	. "github.com/moutend/go-backlog/pkg/types"
)

// GetProjectStatuses returns list of status in the project.
//
// For more details, see the API document.
//
// https://developer.nulab.com/docs/backlog/api/2/get-status-list-of-project/#get-status-list-of-project
func (c *Client) GetProjectStatuses(projectIdOrKey string) ([]*ProjectStatus, error) {
	return c.GetProjectStatusesContext(context.Background(), projectIdOrKey)
}

// GetProjectStatuses accepts context.
func (c *Client) GetProjectStatusesContext(ctx context.Context, projectIdOrKey string) ([]*ProjectStatus, error) {
	path, err := c.root.Parse(path.Join(V2ProjectsPath, projectIdOrKey, "statuses"))

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

	var projectStatuses []*ProjectStatus

	if err := json.Unmarshal(body, &projectStatuses); err != nil {
		return nil, err
	}

	return projectStatuses, nil
}

// GetProject returns information about project.
//
// For more details, see the API document.
//
// https://developer.nulab.com/docs/backlog/api/2/get-project/#get-project
func (c *Client) GetProject(projectKeyOrId string) (*Project, error) {
	return c.GetProjectContext(context.Background(), projectKeyOrId)
}

// GetProjectContext accepts context.
func (c *Client) GetProjectContext(ctx context.Context, projectKeyOrId string) (*Project, error) {
	path, err := c.root.Parse(path.Join(V2ProjectsPath, projectKeyOrId))

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

	var project Project

	if err := json.Unmarshal(body, &project); err != nil {
		return nil, err
	}

	return &project, nil
}

// GetProjects returns list of projects.
//
// For more details, see the API document.
//
// https://developer.nulab.com/docs/backlog/api/2/get-project-list/#get-project-list
func (c *Client) GetProjects(query url.Values) ([]*Project, error) {
	return c.GetProjectsContext(context.Background(), query)
}

// GetProjectsContext accepts context.
func (c *Client) GetProjectsContext(ctx context.Context, query url.Values) ([]*Project, error) {
	path, err := c.root.Parse(V2ProjectsPath)

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

	var projects []*Project

	if err := json.Unmarshal(body, &projects); err != nil {
		return nil, err
	}

	return projects, nil
}
