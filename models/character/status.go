package character

import (
	"github.com/over-engineering/msa/models/brain"
	"github.com/over-engineering/msa/models/lifecycle"
	"github.com/over-engineering/msa/models/needs"
	"github.com/over-engineering/msa/models/physical"
)

// Status represents all aspects of character's current state.
type Status struct {
	// 나이
	lifecycle.LifeCycle `json:"life_cycle"` // range: 0~
	// Physical
	physical.Physical `json:"physical"`
	// Brain
	brain.Brain `json:"brain"`
	// Needs
	needs.Needs    `json:"needs"`
	needs.Hormones `json:"hormones"`
}

// NewStatus returns Status with given information
func NewStatus(val float32) Status {
	return Status{
		LifeCycle: lifecycle.LifeCycle{
			Alive: true,
		},
		Physical: physical.Physical{
			Health: 100,
			Height: float32(180.0),
			Weight: float32(78.3),
			Charm:  50,
		},
		Brain:    brain.NewBrain(val),
		Needs:    needs.NewNeeds(val),
		Hormones: needs.NewHormones(val),
	}
}

// ApplyEffects apply effects to status one by one.
// func (s *Status) ApplyEffects(es Effects) {
// 	for _, e := range es {
// 		switch {
// 		case e.Target == "Intelligences":
// 			s.Intelligences.UpdateValue(e.Value.(map[IntelligenceType]float32))
// 		case e.Target == "Personality":
// 			s.Personality.UpdateValue(e.Value.(Personality))
// 		case e.Target == "Mentals":
// 			s.Mentals.UpdateValue(e.Value.(Mentals), s.Personality)
// 		case e.Target == "Needs":
// 			s.Needs.UpdateValue(e.Value.(Needs), s.Hormones)
// 		case e.Target == "Hormones":
// 			s.Hormones.UpdateValue(e.Value.(Hormones))
// 		// case e.Target == "Intelligence":
// 		// 	s.Intelligence.UpdateValue(Value.(Intelligence))
// 		default:

// 			// f := reflect.ValueOf(s).Elem().FieldByName(e.Target)
// 			// f.SetFloat(f.Float() + e.Value.(float64))
// 			v := reflect.ValueOf(s).Elem().FieldByName(e.Target)
// 			u := v.Interface().(float32)
// 			v.SetFloat(float64(u + e.Value.(float32)))

// 		}
// 	}
// }
