package finance

import (
	"errors"

	"github.com/over-engineering/msa/models/types"
)

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
	TransferById(amount Dollars, from, to types.UID)
}

// TaxCollector represents 국세청's functionality.
type TaxCollector interface {
	CalculateTax(annualIncome Dollars) Dollars
	CollectTax(from Payer, amount Dollars) error
}

// Dollars represents the currency unit in the game.
type Dollars float32

// Finance represents the character's financial states.
// We may implements diverse types of financial states like stocks here.
type Finance struct {
	Balance Dollars `json:"balance"`
	TaxInfo TaxInfo `json:"tax_info"`
}

// AddBalance add value to the character's balance.
func (f *Finance) AddBalance(val Dollars) {
	f.Balance += val
}

// SubBalance subtracts the character's balance only if it is enough.
// We may consider minus value later.
func (f *Finance) SubBalance(val Dollars) error {
	if f.Balance >= val {
		f.Balance -= val
		return nil
	}
	return errors.New("character's balance is not enough")
}

// TaxInfo represents the tax related infos charged to the character.
// We may divide incomes to income records. And give the summary
// to user when tax is charged.
type TaxInfo struct {
	AnnualIncome Dollars      `json:"annual_income"` // TODO: 월급으로 계산 or 지속적으로 돈이 벌릴때 마다 update
	TaxCollector TaxCollector `json:"tax_collector"` // 국적에 따라 scale은 업데이트 되어야 함.
}

// CalculateTax returns tax value according to the TaxInfo members.
func (t *TaxInfo) CalculateTax() Dollars {
	return t.TaxCollector.CalculateTax(t.AnnualIncome)
}

// RevenueA is an example of tax rate of country A.
type RevenueA struct {
	Balance Dollars `json:"balance"`
}

// CalculateTax returns how much tax should payer pay for given
// amount of money.
func (ra *RevenueA) CalculateTax(a Dollars) Dollars {
	switch {
	case a < Dollars(12000):
		return a * Dollars(0.06)
	case a < Dollars(46000):
		return a * Dollars(0.15)
	default:
		return a * Dollars(0.24)
	}
}

// CollectTax forces to pay the amount of money from payer and
// increases revenue's balance.
func (ra *RevenueA) CollectTax(from Payer, amount Dollars) error {
	from.Pay(amount)
	ra.Balance += amount
	return nil
}

// 매년 말 그동안 쌓아온 AnnualIncome에 대해 세금을 부과한다.
