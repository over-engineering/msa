package condition

import (
	"time"

	"github.com/over-engineering/msa/models/character"
)

// DrugAddiction represents addicted condition.
type DrugAddiction struct {
	Name      string    `json:"name"`
	Intensity float32   `json:"intensity"`
	StartTime time.Time `json:"start_time"`
	EndTime   time.Time `json:"end_time"`
}

// NewDrugAddiction retruns DrugAddiction
func NewDrugAddiction(intensity float32, ct time.Time, d time.Duration) DrugAddiction {
	addiction := DrugAddiction{}
	addiction.Name = "약물 중독"
	addiction.Intensity = intensity
	addiction.StartTime = ct
	addiction.EndTime = ct.Add(d)
	return addiction
}

// ApplyEffect applies DrugAddiction's effect to character
func (da DrugAddiction) ApplyEffect(c *character.Character) error {
	c.Status.Charm -= da.Intensity
	return nil
}

// CancelEffect cancels DrugAddiction's effect to character
func (da DrugAddiction) CancelEffect(c *character.Character) error {
	c.Status.Charm += da.Intensity
	return nil
}

// Validate returns boolean value of whether this effect still valid.
func (da DrugAddiction) Validate(t time.Time) bool {
	return da.EndTime.After(t)
}

// Describe returns string value of DrugAddiction's description.
func (da DrugAddiction) Describe() string {
	description := da.Name
	return description
}
