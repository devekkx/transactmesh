package events

import (
	"context"

	nats "github.com/devekkx/transactmesh/packages/nats"
)

type Publisher struct {
	bus *nats.Client
}

func NewPublisher(bus *nats.Client) *Publisher {
	return &Publisher{bus: bus}
}

func (p *Publisher) WalletDebited(ctx context.Context, walletID string, amount int64) error {

	return p.bus.Publish(ctx, "wallet.debited", map[string]any{
		"wallet_id": walletID,
		"amount":    amount,
	})
}
