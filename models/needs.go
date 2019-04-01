package models

import "fmt"

type NeedsType int

const (
	FoodNeeds        NeedsType = iota // 식욕
	SleepNeeds                        // 수면욕
	SexNeeds                          // 성욕
	RestNeeds                         // 태만(휴식)욕?
	StabilityNeeds                    // 안정욕
	AffectionNeeds                    // 애정욕
	ReputationNeeds                   // 명예욕
	AchievementNeeds                  // 성취욕
)

func (nType NeedsType) String() string {
	names := [8]string{
		"FoodNeeds",
		"SleepNeeds",
		"SexNeeds",
		"RestNeeds",
		"StabilityNeeds",
		"AffectionNeeds",
		"ReputationNeeds",
		"AchievementNeeds",
	}

	if nType < FoodNeeds || nType > AchievementNeeds {
		return "Error"
	}

	return names[nType]
}

// needs가 호르몬에 주는 영향 table
// needs가 threshold보다 낮을 때는 큰 영향을 주지 않음 ex) 0.01,
// needs가 threshold보다 높아지면 호르몬에 큰 영향을 줌 ex) 0.1
var NeedsTable = []map[HormoneType][]float32{
	FoodNeeds:        map[HormoneType][]float32{Serotonin: []float32{0.01, 0.1}, Norepinephrine: []float32{0.01, 0.1}, Epinephrine: []float32{0.01, 0.2}, Dopamine: []float32{0.01, 0}, Endorphin: []float32{0.01, 0.1}, Oxytocin: []float32{0.01, 0.1}},
	SleepNeeds:       map[HormoneType][]float32{Serotonin: []float32{0.01, 0.1}, Norepinephrine: []float32{0.01, 0.1}, Epinephrine: []float32{0.01, 0.2}, Dopamine: []float32{0.01, 0}, Endorphin: []float32{0.01, 0.1}, Oxytocin: []float32{0.01, 0.1}},
	SexNeeds:         map[HormoneType][]float32{Serotonin: []float32{0.01, 0.1}, Norepinephrine: []float32{0.01, 0.1}, Epinephrine: []float32{0.01, 0.2}, Dopamine: []float32{0.01, 0}, Endorphin: []float32{0.01, 0.1}, Oxytocin: []float32{0.01, 0.1}},
	RestNeeds:        map[HormoneType][]float32{Serotonin: []float32{0.01, 0.1}, Norepinephrine: []float32{0.01, 0.1}, Epinephrine: []float32{0.01, 0.2}, Dopamine: []float32{0.01, 0}, Endorphin: []float32{0.01, 0.1}, Oxytocin: []float32{0.01, 0.1}},
	StabilityNeeds:   map[HormoneType][]float32{Serotonin: []float32{0.01, 0.1}, Norepinephrine: []float32{0.01, 0.1}, Epinephrine: []float32{0.01, 0.2}, Dopamine: []float32{0.01, 0}, Endorphin: []float32{0.01, 0.1}, Oxytocin: []float32{0.01, 0.1}},
	AffectionNeeds:   map[HormoneType][]float32{Serotonin: []float32{0.01, 0.1}, Norepinephrine: []float32{0.01, 0.1}, Epinephrine: []float32{0.01, 0.2}, Dopamine: []float32{0.01, 0}, Endorphin: []float32{0.01, 0.1}, Oxytocin: []float32{0.01, 0.1}},
	ReputationNeeds:  map[HormoneType][]float32{Serotonin: []float32{0.01, 0.1}, Norepinephrine: []float32{0.01, 0.1}, Epinephrine: []float32{0.01, 0.2}, Dopamine: []float32{0.01, 0}, Endorphin: []float32{0.01, 0.1}, Oxytocin: []float32{0.01, 0.1}},
	AchievementNeeds: map[HormoneType][]float32{Serotonin: []float32{0.01, 0.1}, Norepinephrine: []float32{0.01, 0.1}, Epinephrine: []float32{0.01, 0.2}, Dopamine: []float32{0.01, 0}, Endorphin: []float32{0.01, 0.1}, Oxytocin: []float32{0.01, 0.1}},
}

const NeedsThreshold = 100

type Needs map[NeedsType]float32

const (
	FoodNeedsRisingVal  = 50
	SleepNeedsRisingVal = 50
	SexNeedsRisingVal   = 5
)

// needs 업데이트하고 needs table에 따라 호르몬의 영향을 계산하여 호르몬 업데이트
func (ns Needs) UpdateValue(v Needs, hs Hormones) {
	hormoneMap := Hormones{}
	for key, value := range v {
		if value == 0 {
			continue
		}
		ns.UpdateHormoneMap(key, ns[key], value, hormoneMap)
		ns[key] += value

		if ns[key] < 0 {
			ns[key] = 0
		}
	}

	hs.UpdateValue(hormoneMap)
	fmt.Println("Update Needs", v, hs, ns)
}

func (ns Needs) UpdateHormoneMap(t NeedsType, orgVal float32, updateVal float32, hormoneMap Hormones) Hormones {
	effVal := [2]float32{0, 0}

	// 증가하는 경우
	if updateVal >= 0 {
		if orgVal <= NeedsThreshold {
			if orgVal+updateVal <= NeedsThreshold {
				effVal[0] = updateVal
			} else {
				effVal[0] = NeedsThreshold - orgVal
				effVal[1] = updateVal - effVal[0]
			}
		} else {
			effVal[1] = updateVal
		}

	} else { // 감소하는 경우
		if orgVal > NeedsThreshold {
			if orgVal+updateVal < NeedsThreshold {
				effVal[1] = NeedsThreshold - orgVal
				effVal[0] = updateVal - effVal[1]
			} else {
				effVal[1] = updateVal
			}
		} else {
			effVal[0] = updateVal
		}
	}

	for key, list := range NeedsTable[t] {
		// *hormoneMap[key] += val * effVal
		for i, val := range list {
			hormoneMap[key] += val * effVal[i]
		}
	}

	return hormoneMap
}
