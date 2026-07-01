package events

import (
	"context"
	"encoding/json"

	"github.com/nats-io/nats.go"
	"go.opentelemetry.io/otel/trace"
)

type Publisher struct {
	conn *nats.Conn
}

func NewPublisher(conn *nats.Conn) *Publisher {
	return &Publisher{conn: conn}
}

func (p *Publisher) Publish(ctx context.Context, subject string, payload []byte) error {
	span := trace.SpanFromContext(ctx)
	traceID := span.SpanContext().TraceID().String()

	msg := map[string]any{
		"trace_id": traceID,
		"payload":  payload,
	}

	data, _ := json.Marshal(msg)

	return p.conn.Publish(subject, data)
}
