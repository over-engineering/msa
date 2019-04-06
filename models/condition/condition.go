package condition

import (
	"errors"
	"time"

	"github.com/over-engineering/msa/models/status"
	"github.com/over-engineering/msa/models/status/physical"
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
func (da DrugAddiction) ApplyEffect(s *status.Status) error {
	s.Charm -= da.Intensity
	return nil
}

// CancelEffect cancels DrugAddiction's effect to character
func (da DrugAddiction) CancelEffect(s *status.Status) error {
	s.Charm += da.Intensity
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

// NeckFracture represents neck bone fracture.
type NeckFracture struct {
	Name      string    `json:"name"`
	Intensity float32   `json:"intensity"`
	StartTime time.Time `json:"start_time"`
	EndTime   time.Time `json:"end_time"`
}

func (nf NeckFracture) ApplyEffect(s *status.Status) error {
	neck, ok := s.Physical.Bones[physical.NeckBone]
	if !ok {
		return errors.New("no neck bone available")
	}
	neck.Cracked = true
	neck.Condition -= nf.Intensity / neck.Strength
	return nil
}

func (nf NeckFracture) CancelEffect(s *status.Status) error {
	neck, ok := s.Physical.Bones[physical.NeckBone]
	if !ok {
		return errors.New("no neck bone available")
	}
	neck.Cracked = false
	neck.Condition += nf.Intensity / neck.Strength
	return nil
}

// Validate returns boolean value of whether this effect still valid.
func (nf NeckFracture) Validate(t time.Time) bool {
	return nf.EndTime.After(t)
}

// Describe returns string value of DrugAddiction's description.
func (nf NeckFracture) Describe() string {
	description := nf.Name
	return description
}
