package character

import (
	"errors"
	"strings"
	"time"

	"github.com/over-engineering/msa/models/fan"
	"github.com/over-engineering/msa/models/finance"
	"github.com/over-engineering/msa/models/talent"
	"github.com/over-engineering/msa/models/types"
)

// ConditionUpdater updates character values.
type ConditionUpdater interface {
	ApplyEffect(c *Character) error
	CancelEffect(c *Character) error
	Validate(t time.Time) bool
	Describe() string
}

// Wearer updates character values.
type Wearer interface {
	GetEquipmentType() types.EquipmentType
	ApplyEffect(c *Character) error
	CancelEffect(c *Character) error
	Describe() string
}

// Character represents the game character belongs to user
// Things could be referencing with unique IDs.
type Character struct {
	ID        types.UID     `json:"id,string"`
	FirstName string        `json:"first_name"`
	LastName  string        `json:"last_name"`
	JobType   types.JobType `json:"job_type"`
	FanInfo   fan.Info      `json:"fan_info"`

	// TODO: Nationality
	// TODO: user struct

	// Location *Location `json:"location"`

	Status     Status             `json:"status"`
	Conditions []ConditionUpdater `json:"conditions"`
	Contracts  []types.UID        `json:"contracts"`
	// Equipments map[types.EquipmentType]*Wearer `json:"equipments"` this goes to game server
	talent.Talents `json:"talents"`

	finance.Finance `json:"finance"`
}

func FindCharacterByID(id types.UID) *Character {
	// TODO:
	return &Character{}
}

// NewCharacter returns new Character object
func NewCharacter(
	uid types.UID,
	fName, lName string,
	status Status,
) *Character {
	return &Character{
		ID:         uid,
		FirstName:  fName,
		LastName:   lName,
		Status:     status,
		Conditions: []ConditionUpdater{},
		Contracts:  []types.UID{},
		// Equipments: map[types.EquipmentType]*Wearer{},
		Finance: finance.Finance{
			Balance: 0,
		},
		FanInfo: fan.Info{
			Fan:  2,
			Anti: 0,
		},
	}
}

// AddCondition adds updater interface to Conditions
func (c *Character) AddCondition(e ConditionUpdater) error {
	c.Conditions = append(c.Conditions, e)
	e.ApplyEffect(c)
	return nil
}

// RemoveCondition removes updater from conditions
func (c *Character) RemoveCondition(i int) error {
	if i >= len(c.Conditions) {
		return errors.New("index out of range")
	}
	c.Conditions[i].CancelEffect(c)
	c.Conditions = append(c.Conditions[:i], c.Conditions[i+1:]...)
	return nil
}

// UpdateCondition updates conditions with given time
func (c *Character) UpdateCondition(ct time.Time) error {
	for i, condition := range c.Conditions {
		if !condition.Validate(ct) {
			err := c.RemoveCondition(i)
			return err
		}
	}
	return nil
}

// CurrentJob returns character's current job.
// From this result we could adjust UI, events etc.
func (c *Character) CurrentJob() types.JobType {
	return c.JobType
}

// func (c *Character) CurrentLocation() *Location {
// 	return c.Location
// }

// CurrentContractIDs returns character's current contracts' ids.
func (c *Character) CurrentContractIDs() []types.UID {
	return c.Contracts
}

// GetName returns character's full name. This function have to
// consider user's linguistic setting.
func (c Character) GetName() string {
	// TODO: user's  linguistic setting check
	return strings.Join([]string{c.FirstName, c.LastName}, " ")
}

// Pay pays the amount of money from its balance
func (c *Character) Pay(amount finance.Dollars) error {
	if c.Balance < amount {
		return errors.New("not enough balance for this entity")
	}
	c.Balance -= amount
	return nil
}

// Paid increases the amount of money in entity's balance.
func (c *Character) Paid(amount finance.Dollars) error {
	c.Balance += amount
	return nil
}

// GetID returns Entity's ID.
func (c Character) GetID() types.UID {
	return c.ID
}

// Wear wears the Wearer object and take off previous one if exists.
// func (c *Character) Wear(w *Wearer) error {
// 	et := (*w).GetEquipmentType()
// 	if _, ok := c.Equipments[et]; !ok {
// 		c.Equipments[et] = w
// 		(*w).ApplyEffect(c)
// 		return nil
// 	}
// 	prev := c.Equipments[et]
// 	(*prev).CancelEffect(c)
// 	c.Equipments[et] = w
// 	(*w).ApplyEffect(c)
// 	return nil
// }

// // TakeOff takes off the Wearer from character.
// func (c *Character) TakeOff(w *Wearer) error {
// 	delete(c.Equipments, (*w).GetEquipmentType())
// 	(*w).CancelEffect(c)
// 	return nil
// }
