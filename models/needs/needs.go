package needs

import "github.com/over-engineering/msa/models/types"

type NeedsType types.Type

const (
	Food        NeedsType = iota // 식욕
	Sleep                        // 수면욕
	Sex                          // 성욕
	Rest                         // 태만(휴식)욕?
	Stability                    // 안정욕
	Affection                    // 애정욕
	Reputation                   // 명예욕
	Achievement                  // 성취욕
)

var NeedsTable = []map[HormoneType]float32{
	Food:        map[HormoneType]float32{Serotonin: 0.1, Norepinephrine: 0.1, Epinephrine: 0.2, Dopamine: 0, Endorphin: 0.1, Oxytocin: 0.1},
	Sleep:       map[HormoneType]float32{Serotonin: 0.1, Norepinephrine: 0.1, Epinephrine: 0.2, Dopamine: 0, Endorphin: 0.1, Oxytocin: 0.1},
	Sex:         map[HormoneType]float32{Serotonin: 0.1, Norepinephrine: 0.1, Epinephrine: 0.2, Dopamine: 0, Endorphin: 0.1, Oxytocin: 0.1},
	Rest:        map[HormoneType]float32{Serotonin: 0.1, Norepinephrine: 0.1, Epinephrine: 0.2, Dopamine: 0, Endorphin: 0.1, Oxytocin: 0.1},
	Stability:   map[HormoneType]float32{Serotonin: 0.1, Norepinephrine: 0.1, Epinephrine: 0.2, Dopamine: 0, Endorphin: 0.1, Oxytocin: 0.1},
	Affection:   map[HormoneType]float32{Serotonin: 0.1, Norepinephrine: 0.1, Epinephrine: 0.2, Dopamine: 0, Endorphin: 0.1, Oxytocin: 0.1},
	Reputation:  map[HormoneType]float32{Serotonin: 0.1, Norepinephrine: 0.1, Epinephrine: 0.2, Dopamine: 0, Endorphin: 0.1, Oxytocin: 0.1},
	Achievement: map[HormoneType]float32{Serotonin: 0.1, Norepinephrine: 0.1, Epinephrine: 0.2, Dopamine: 0, Endorphin: 0.1, Oxytocin: 0.1},
}

const NeedsThreshold = 100

type Needs map[NeedsType]float32

func (ns Needs) UpdateValue(v Needs, hs Hormones) {
	hormoneMap := Hormones{}
	for key, value := range v {
		if value == 0 {
			continue
		}
		ns.UpdateHormoneMap(key, ns[key], value, hormoneMap)
		ns[key] += value
	}

	hs.UpdateValue(hormoneMap)

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
