package nats

import (
	"context"

	"github.com/nats-io/nats.go"
)

func publish(ctx context.Context, c *Client, subject string, data []byte) error {

	msg := &nats.Msg{
		Subject: subject,
		Data:    data,
		Header:  nats.Header{},
	}

	Inject(ctx, msg)

	_, err := c.js.PublishMsg(msg)
	return err
}
