package nats

import (
	"context"
	"encoding/json"
	"log"

	"github.com/nats-io/nats.go"
)

type HandlerFunc func(ctx context.Context, msg *Message) error

type Message struct {
	Subject string
	Data    []byte
	Headers map[string]string
}

func (c *Client) Subscribe(subject string, handler HandlerFunc) error {

	_, err := c.conn.Subscribe(subject, func(msg *nats.Msg) {

		ctx := context.Background()

		var payload map[string]any
		if err := json.Unmarshal(msg.Data, &payload); err != nil {
			log.Printf("[NATS] invalid payload: %v", err)
			return
		}

		err := handler(ctx, &Message{
			Subject: msg.Subject,
			Data:    msg.Data,
		})

		if err != nil {
			log.Printf("[NATS] handler error: %v", err)
		}
	})

	return err
}
