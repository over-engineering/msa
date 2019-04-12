package finance

import "github.com/over-engineering/msa/common/models/types"

// Payer represents interface for those who could pay and paid.
type Payer interface {
	Pay(amount Dollars) error
	Paid(amount Dollars) error
	GetName() string
	GetID() types.UID
	// GetAccountNumber() AccountNumber
}

// Banker represents interface for transfering money from Payer to Payer.
type Banker interface {
	Transfer(amount Dollars, from, to Payer) error
	TransferById(amount Dollars, from, to types.UID) error
}

// TaxCollector represents 국세청's functionality.
type TaxCollector interface {
	CalculateTax(annualIncome Dollars) Dollars
	CollectTax(from Payer, amount Dollars) error
}
