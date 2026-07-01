package wallet

import "time"

type Wallet struct {
	ID        string
	Balance   int64
	Currency  string
	CreatedAt time.Time
	UpdatedAt time.Time
}
