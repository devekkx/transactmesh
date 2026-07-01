package nats

import (
	"context"

	"github.com/nats-io/nats.go"
)

func subscribe(c *Client, subject string, handler HandlerFunc) (*nats.Subscription, error) {

	sub, err := c.js.Subscribe(
		subject,
		func(msg *nats.Msg) {

			ctx := Extract(context.Background(), msg)

			m := &Message{
				Subject: msg.Subject,
				Data:    msg.Data,
			}

			if err := handler(ctx, m); err != nil {
				_ = msg.Nak()
				return
			}

			_ = msg.Ack()
		},
		nats.Durable("wallet-service-"+subject),
		nats.ManualAck(),
		nats.AckExplicit(),
	)

	if err != nil {
		return nil, err
	}

	return sub, nil
}
