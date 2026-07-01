package nats

import "context"

type HandlerFunc func(ctx context.Context, msg *Message) error

type Message struct {
	Subject       string
	Data          []byte
	Headers       map[string]string
	CorrelationID string
	TraceID       string
}
