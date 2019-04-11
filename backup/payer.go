package models

import "github.com/over-engineering/msa/models/types"

// Payer represents interface for those who could pay and paid.
type Payer interface {
	Pay(amount float32) error
	Paid(amount float32) error
	GetName() string
	GetID() types.UID
	// GetAccountNumber() AccountNumber
}
