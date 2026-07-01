package nats

import (
	"context"

	"github.com/nats-io/nats.go"
)

type EventBus struct {
	client *Client
}

func NewEventBus(c *Client) *EventBus {
	return &EventBus{client: c}
}

func (e *EventBus) Publish(ctx context.Context, subject string, payload []byte) error {
	return publish(ctx, e.client, subject, payload)
}

func (e *EventBus) Subscribe(subject string, handler HandlerFunc) (*nats.Subscription, error) {
	return subscribe(e.client, subject, handler)
}
