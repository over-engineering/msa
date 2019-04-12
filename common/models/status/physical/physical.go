package physical

import "math/rand"

type Physical struct {
	Health float32 `json:"health"`
	Height float32 `json:"height"`
	Weight float32 `json:"weight"`

	Bones   `json:"bone"`
	Joints  map[JointType]*Joint   `json:"joints"`
	Muscles map[MuscleType]*Muscle `json:"muscles"`
	Fat     `json:"fats"`          // 체지방률

	Resilience float32 `json:"resilience"`
	Cardio     float32 `json:"cardio"`
	Charm      float32 `json:"charm"`
	Digestion  float32 `json:"digestion"` // 소화력 = 소화 흡수율 (당장은)

	ProducedCal float32 `json:"produced_cal"`
	ConsumedCal float32 `json:"consumed_cal"`
}

func NewPhysical(
	height, weight float32,
	r *rand.Rand,
	b, s float32,
	fat float32,
) Physical {
	return Physical{
		Health:      100,
		Height:      height,
		Weight:      weight,
		Bones:       NewBones(r, b, s),
		Joints:      map[JointType]*Joint{},
		Muscles:     map[MuscleType]*Muscle{},
		Fat:         Fat{Mass: fat},
		Resilience:  (r.Float32()-0.5)*s + b,
		Cardio:      (r.Float32()-0.5)*s + b,
		Charm:       (r.Float32()-0.5)*s + b,
		Digestion:   (r.Float32()-0.5)*s + b,
		ProducedCal: 0,
		ConsumedCal: 0,
	}
}
