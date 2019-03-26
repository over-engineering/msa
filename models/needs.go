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

var NeedsTable = []map[HormoneType]float32{
	FoodNeeds:        map[HormoneType]float32{Serotonin: 0.1, Norepinephrine: 0.1, Epinephrine: 0.2, Dopamine: 0, Endorphin: 0.1, Oxytocin: 0.1},
	SleepNeeds:       map[HormoneType]float32{Serotonin: 0.1, Norepinephrine: 0.1, Epinephrine: 0.2, Dopamine: 0, Endorphin: 0.1, Oxytocin: 0.1},
	SexNeeds:         map[HormoneType]float32{Serotonin: 0.1, Norepinephrine: 0.1, Epinephrine: 0.2, Dopamine: 0, Endorphin: 0.1, Oxytocin: 0.1},
	RestNeeds:        map[HormoneType]float32{Serotonin: 0.1, Norepinephrine: 0.1, Epinephrine: 0.2, Dopamine: 0, Endorphin: 0.1, Oxytocin: 0.1},
	StabilityNeeds:   map[HormoneType]float32{Serotonin: 0.1, Norepinephrine: 0.1, Epinephrine: 0.2, Dopamine: 0, Endorphin: 0.1, Oxytocin: 0.1},
	AffectionNeeds:   map[HormoneType]float32{Serotonin: 0.1, Norepinephrine: 0.1, Epinephrine: 0.2, Dopamine: 0, Endorphin: 0.1, Oxytocin: 0.1},
	ReputationNeeds:  map[HormoneType]float32{Serotonin: 0.1, Norepinephrine: 0.1, Epinephrine: 0.2, Dopamine: 0, Endorphin: 0.1, Oxytocin: 0.1},
	AchievementNeeds: map[HormoneType]float32{Serotonin: 0.1, Norepinephrine: 0.1, Epinephrine: 0.2, Dopamine: 0, Endorphin: 0.1, Oxytocin: 0.1},
}

const NeedsThreshold = 100

type Needs map[NeedsType]float32

const (
	FoodNeedsRisingVal  = 50
	SleepNeedsRisingVal = 50
	SexNeedsRisingVal   = 5
)

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
	var effVal float32
	if updateVal >= 0 {
		if resVal := orgVal + updateVal; resVal > NeedsThreshold {
			effVal = updateVal
			if orgVal < NeedsThreshold {
				effVal = resVal - NeedsThreshold
			}
		}
	} else {
		if orgVal > NeedsThreshold {
			effVal = updateVal
			if resVal := orgVal + updateVal; resVal < NeedsThreshold {
				effVal = NeedsThreshold - orgVal
			}

		}
	}

	for key, val := range NeedsTable[t] {
		// *hormoneMap[key] += val * effVal
		hormoneMap[key] += val * effVal
	}
	return hormoneMap
}
