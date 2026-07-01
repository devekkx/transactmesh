package nats

import (
	"context"
	"encoding/json"

	"github.com/nats-io/nats.go"
)

func (c *Client) Publish(ctx context.Context, subject string, payload any) error {

	_ = ctx // reserved for future OTEL injection

	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	msg := &nats.Msg{
		Subject: subject,
		Data:    data,
	}

	return c.conn.PublishMsg(msg)
}
