package nats

import "github.com/nats-io/nats.go"

func (c *Client) JetStream() nats.JetStreamContext {
	return c.js
}
