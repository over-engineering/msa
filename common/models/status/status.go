package status

import (
	"math/rand"
	"time"

	"github.com/over-engineering/msa/common/models/status/brain"
	"github.com/over-engineering/msa/common/models/status/life"
	"github.com/over-engineering/msa/common/models/status/needs"
	"github.com/over-engineering/msa/common/models/status/physical"
)

// Status represents all aspects of character's current state.
type Status struct {
	// 나이
	life.LifeCycle `json:"life_cycle"` // range: 0~
	// Physical
	physical.Physical `json:"physical"`
	// Brain
	brain.Brain `json:"brain"`
	// Needs
	needs.Needs    `json:"needs"`
	needs.Hormones `json:"hormones"`
}

// NewStatus returns Status with given information
func NewStatus(
	born, current time.Time,
	generation int,
	pBase, pSigma, bBase, bSigma, nBase, nSigma, hBase, hSigma float32,
	height, weight, fat float32,
) *Status {
	dead := born.Add(80 * 24 * 365 * time.Hour)
	age := current.Year() - born.Year()
	seed := rand.NewSource(time.Now().UnixNano())
	r := rand.New(seed)
	return &Status{
		LifeCycle: life.LifeCycle{
			Alive:      true,
			Born:       born,
			Dead:       dead,
			Age:        age,
			Generation: generation,
		},
		Physical: physical.NewPhysical(
			height,
			weight,
			r,
			pBase,
			pSigma,
			fat,
		),
		Brain:    brain.NewBrain(r, bBase, bSigma),
		Needs:    needs.NewNeeds(r, nBase, nSigma),
		Hormones: needs.NewHormones(r, hBase, hSigma),
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
