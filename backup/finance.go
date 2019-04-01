package models

import "errors"

// Finance represents the character's financial states.
// We may implements diverse types of financial states like stocks here.
type Finance struct {
	Balance float32 `json:"balance"`
	TaxInfo TaxInfo `json:"tax_info"`
}

// AddBalance add value to the character's balance.
func (f *Finance) AddBalance(val float32) {
	f.Balance += val
}

// SubBalance subtracts the character's balance only if it is enough.
// We may consider minus value later.
func (f *Finance) SubBalance(val float32) error {
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
	AnnualIncome float32 `json:"annual_income"` // TODO: 월급으로 계산 or 지속적으로 돈이 벌릴때 마다 update
	Scale        Scale   `json:"scale"`         // 국적에 따라 scale은 업데이트 되어야 함.
}

// CalculateTax returns tax value according to the TaxInfo members.
func (t *TaxInfo) CalculateTax() float32 {
	return t.Scale.calculateTax(t.AnnualIncome)
}

// Scale represents the tax rate related interface
type Scale interface {
	calculateTax(annualIncome float32) float32
}

// ScaleA is an example of tax rate of country A.
type ScaleA struct{}

func (sa *ScaleA) calculateTax(a float32) float32 {
	switch {
	case a < 12000:
		return a * 0.06
	case a >= 12000 && a < 46000:
		return a * 0.15
	default:
		return a * 0.24
	}
}

// 매년 말 그동안 쌓아온 AnnualIncome에 대해 세금을 부과한다.
