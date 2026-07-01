package wallet

import (
	"context"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
)

type Service struct{}

func NewService() *Service {
	return &Service{}
}

func (s *Service) Debit(ctx context.Context, walletID string, amount int64) error {

	tracer := otel.Tracer("wallet-service")

	ctx, span := tracer.Start(ctx, "wallet.debit")
	defer span.End()

	span.SetAttributes(
		attribute.String("wallet.id", walletID),
		attribute.Int64("wallet.amount", amount),
	)

	// business logic here

	return nil
}
