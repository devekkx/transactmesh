package wallet

import (
	"context"
	"errors"
	"sync"
)

type InMemoryRepo struct {
	mu      sync.RWMutex
	storage map[string]*Wallet
}

func NewInMemoryRepo() Repository {
	return &InMemoryRepo{
		storage: make(map[string]*Wallet),
	}
}

func (r *InMemoryRepo) GetByID(ctx context.Context, id string) (*Wallet, error) {

	r.mu.RLock()
	defer r.mu.RUnlock()

	w, ok := r.storage[id]
	if !ok {
		return nil, errors.New("wallet not found")
	}

	// return copy (avoid mutation leaks)
	copy := *w
	return &copy, nil
}

func (r *InMemoryRepo) Save(ctx context.Context, w *Wallet) error {

	r.mu.Lock()
	defer r.mu.Unlock()

	// store copy (avoid shared memory mutation issues)
	copy := *w
	r.storage[w.ID] = &copy

	return nil
}
