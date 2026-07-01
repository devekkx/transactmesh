package nats

import (
	"context"
	"encoding/json"

	"github.com/nats-io/nats.go"
	"go.opentelemetry.io/otel/trace"
)

type Publisher struct {
	conn *nats.Conn
}

func New(conn *nats.Conn) *Publisher {
	return &Publisher{conn: conn}
}

func (p *Publisher) Publish(ctx context.Context, subject string, payload any) error {
	span := trace.SpanFromContext(ctx)

	msg := map[string]any{
		"trace_id": span.SpanContext().TraceID().String(),
		"payload":  payload,
	}

	data, err := json.Marshal(msg)
	if err != nil {
		return err
	}

	return p.conn.Publish(subject, data)
}
