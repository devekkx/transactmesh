package nats

import (
	"context"
	"encoding/json"
)

func (c *Client) Publish(ctx context.Context, subject string, payload any) error {

	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	return publish(ctx, c, subject, data)
}
