package models

import "strings"

// Character represents the game character belongs to user
// Things could be referencing with unique IDs.
type Character struct {
	ID        int64   `json:"id,string"`
	FirstName string  `json:"first_name"`
	LastName  string  `json:"last_name"`
	Job       JobType `json:"job_type"`
	// TODO: Nationality
	// TODO: user struct

	Status     Status                      `json:"status"`
	Conditions []Condition                 `json:"conditions"`
	Equipments map[EquipmentType]Equipment `json:"equipments"`
	// TODO: 논의 Equipments map[EquipmentType]UUID

	Goods    map[interface{}]interface{} `json:"goods"`
	Finance  Finance                     `json:"finance"` // Or maybe Account
	Contract Contract                    `json:"contract"`
	FanInfo  FanInfo                     `json:"fan_info"`

	// Character would not have a team. Also, team information
	// could be shared with other players.
	Team *Team `json:"team"`
}

// NewCharacter returns new Character object
func NewCharacter() *Character {
	return &Character{}
}

// CurrentJob returns character's current job.
// From this result we could adjust UI, events etc.
func (c *Character) CurrentJob() JobType {
	return c.Job
}

// CurrentContract returns character's current contract.
func (c *Character) CurrentContract() Contract {
	return c.Contract
}

// CurrentTeam returns character's current team.
func (c *Character) CurrentTeam() *Team {
	return c.Team
}

// FullName returns character's full name. This function have to
// consider user's linguistic setting.
func (c *Character) FullName() string {
	// TODO: user's  linguistic setting check
	return strings.Join([]string{c.FirstName, c.LastName}, " ")
}
