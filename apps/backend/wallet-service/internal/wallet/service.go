package wallet

import (
	"context"
	"errors"
)

type EventPublisher interface {
	WalletDebited(ctx context.Context, walletID string, amount int64) error
}

type Service struct {
	repo      Repository
	publisher EventPublisher
}

func NewService(repo Repository, publisher EventPublisher) *Service {
	return &Service{
		repo:      repo,
		publisher: publisher,
	}
}

func (s *Service) Debit(ctx context.Context, walletID string, amount int64) error {

	if amount <= 0 {
		return errors.New("invalid amount")
	}

	w, err := s.repo.GetByID(ctx, walletID)
	if err != nil {
		return err
	}

	if w.Balance < amount {
		return errors.New("insufficient funds")
	}

	w.Balance -= amount

	if err := s.repo.Save(ctx, w); err != nil {
		return err
	}

	// EVENT (after state change)
	return s.publisher.WalletDebited(ctx, walletID, amount)
}
