package nats

import (
	"context"

	"github.com/nats-io/nats.go"
	"go.opentelemetry.io/otel/propagation"
)

var propagator = propagation.TraceContext{}

func Inject(ctx context.Context, msg *nats.Msg) {
	if msg.Header == nil {
		msg.Header = nats.Header{}
	}

	carrier := propagation.HeaderCarrier(msg.Header)
	propagator.Inject(ctx, carrier)
}

func Extract(ctx context.Context, msg *nats.Msg) context.Context {
	if msg == nil || msg.Header == nil {
		return ctx
	}

	carrier := propagation.HeaderCarrier(msg.Header)
	return propagator.Extract(ctx, carrier)
}
