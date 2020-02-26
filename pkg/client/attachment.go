package client

import (
	"context"
	"encoding/json"
	"io/ioutil"

	. "github.com/moutend/go-backlog/pkg/types"
)

// PostAttachment posts an attachment file for issue or wiki. Returns id of the attachment file.
//
// For more details, see the API document.
//
// https://developer.nulab.com/docs/backlog/api/2/post-attachment-file/#post-attachment-file
func (c *Client) PostAttachment(filepath string) (*Attachment, error) {
	return c.PostAttachmentContext(context.Background(), filepath)
}

// PostAttachmentContext accepts context.
func (c *Client) PostAttachmentContext(ctx context.Context, filepath string) (*Attachment, error) {
	path, err := c.root.Parse(V2SpaceAttachmentPath)

	if err != nil {
		return nil, err
	}

	res, err := c.postFileContext(ctx, path, nil, filepath)

	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		return nil, err
	}

	var attachment *Attachment

	if err := json.Unmarshal(body, &attachment); err != nil {
		return nil, err
	}

	return attachment, nil
}
