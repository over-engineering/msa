package physical

import "math/rand"

type Physical struct {
	Health      float32 `json:"health"`
	Height      float32 `json:"height"`
	Weight      float32 `json:"weight"`
	Fat         float32 `json:"fat"` // 체지방률
	Resilience  float32 `json:"resilience"`
	Cardio      float32 `json:"cardio"`
	Charm       float32 `json:"charm"`
	Digestion   float32 `json:"digestion"` // 소화력 = 소화 흡수율 (당장은)
	ProducedCal float32 `json:"produced_cal"`
	ConsumedCal float32 `json:"consumed_cal"`

	Bones   `json:"bone"`
	Joints  `json:"joints"`
	Muscles `json:"muscles"`
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
		Joints:      NewJoints(r, b, s),
		Muscles:     NewMuscles(r, b, s),
		Fat:         fat,
		Resilience:  (r.Float32()-0.5)*s + b,
		Cardio:      (r.Float32()-0.5)*s + b,
		Charm:       (r.Float32()-0.5)*s + b,
		Digestion:   (r.Float32()-0.5)*s + b,
		ProducedCal: 0,
		ConsumedCal: 0,
	}
}
