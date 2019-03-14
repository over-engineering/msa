package models

// Finance TBD
type Finance struct {
}

// Contract TBD
type Contract struct {
}

// Character represents the game character belongs to user
type Character struct {
	ID        int64   `json:"id,string"`
	FirstName string  `json:"firstName"`
	LastName  string  `json:"lastName"`
	Job       JobType `json:"jobType"`
	// TODO: user struct

	Status     Status                      `json:"status"`
	Conditions []Condition                 `json:"conditions"`
	Equipments map[EquipmentType]Equipment `json:"equipments"`

	Finance         Finance  `json:"finance"` // Or maybe Account
	CurrentContract Contract `json:"currentContract"`
}

// NewCharacter returns new Character object
func NewCharacter() *Character {
	return &Character{}
}
