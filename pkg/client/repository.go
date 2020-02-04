package client

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/url"
	"path"

	. "github.com/moutend/go-backlog/pkg/types"
)

// GetRepositories returns list of Git repositories.
//
// For more details, see the API document.
//
// https://developer.nulab.com/docs/backlog/api/2/get-list-of-git-repositories/#get-list-of-git-repositories
func (c *Client) GetRepositories(projectKeyOrId string, query url.Values) ([]*Repository, error) {
	return c.GetRepositoriesContext(context.Background(), projectKeyOrId, query)
}

// GetRepositoriesContext accepts context.
func (c *Client) GetRepositoriesContext(ctx context.Context, projectKeyOrId string, query url.Values) ([]*Repository, error) {
	path, err := c.root.Parse(path.Join(V2ProjectsPath, projectKeyOrId, "git", "repositories"))

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

	var repositories []*Repository

	if err := json.Unmarshal(body, &repositories); err != nil {
		return nil, err
	}

	return repositories, nil
}

// GetRepository returns Git repository.
//
// For more details, see the API document.
//
// https://developer.nulab.com/docs/backlog/api/2/get-git-repository/
func (c *Client) GetRepository(projectKeyOrId, repositoryNameOrId string) (*Repository, error) {
	return c.GetRepositoryContext(context.Background(), projectKeyOrId, repositoryNameOrId)
}

// getRepositoryContext accepts context.
func (c *Client) GetRepositoryContext(ctx context.Context, projectKeyOrId, repositoryNameOrId string) (*Repository, error) {
	path, err := c.root.Parse(path.Join(
		V2ProjectsPath, projectKeyOrId,
		"git", "repositories", repositoryNameOrId,
	))

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

	var repository *Repository

	if err := json.Unmarshal(body, &repository); err != nil {
		return nil, err
	}

	return repository, nil
}
