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

	// inject tracing + correlation
	Inject(ctx, msg)

	_, err := c.js.PublishMsg(msg)
	if err != nil {
		return err
	}

	return nil
}
