package wallet

import "context"

type Repository interface {
	GetByID(ctx context.Context, id string) (*Wallet, error)
	Save(ctx context.Context, w *Wallet) error
}
