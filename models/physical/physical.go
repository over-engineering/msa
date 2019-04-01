package physical

type Physical struct {
	Health float32 `json:"health"`
	Height float32 `json:"height"`
	Weight float32 `json:"weight"`

	Bones   map[BoneType]*Bone     `json:"bone"`
	Joints  map[JointType]*Joint   `json:"joints"`
	Muscles map[MuscleType]*Muscle `json:"muscles"`
	Fat     map[FatType]*Fat       `json:"fats"`

	Resilience float32 `json:"resilience"`
	Cardio     float32 `json:"cardio"`
	Charm      float32 `json:"charm"`
	Digestion  float32 `json:"digestion"` // 소화력 = 소화 흡수율 (당장은)

	ProducedCal float32 `json:"produced_cal"`
	ComsumedCal float32 `json:"consumed_cal"`
}
