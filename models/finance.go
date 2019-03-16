package models

import "errors"

// Finance represents the character's financial states.
// We may implements diverse types of financial states like stocks here.
type Finance struct {
	Balance int `json:"balance"`
}

// AddBalance add value to the character's balance.
func (f *Finance) AddBalance(val int) {
	f.Balance += val
}

// SubBalance subtracts the character's balance only if it is enough.
// We may consider minus value later.
func (f *Finance) SubBalance(val int) error {
	if f.Balance >= val {
		f.Balance -= val
		return nil
	}
	return errors.New("character's balance is not enough")
}
