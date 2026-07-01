package events

import (
	"context"

	"github.com/devekkx/transactmesh/packages/nats"
)

type Publisher struct {
	client *nats.Client
}

func NewPublisher(client *nats.Client) *Publisher {
	return &Publisher{client: client}
}

func (p *Publisher) PublishWalletDebited(ctx context.Context, walletID string, amount int64) error {
	return p.client.Publish(ctx, "wallet.debited", map[string]any{
		"wallet_id": walletID,
		"amount":    amount,
	})
}
