package models

// Payer represents interface for those who could pay and paid.
type Payer interface {
	Pay(amount float32) error
	Paid(amount float32) error
	GetName() string
	GetID() UID
	// GetAccountNumber() AccountNumber
}
