package football

import "github.com/over-engineering/msa/football/types"

type AbilityType types.Type

const (
	Balance        AbilityType = iota // 0
	Jumping                           // 1
	Pace                              // 2
	Acceleration                      // 3
	Power                             // 4
	Agility                           // 5
	Shooting                          // 6
	Technique                         // 7
	Dribbling                         // 8
	LongShots                         // 9
	Corners                           // 10
	FreeKickTaking                    // 11
	Throws                            // 12
	Crossing                          // 13
	Passing                           // 14
	Tackling                          // 15
	Marking                           // 16
	Heading                           // 17
	PenaltyTaking                     // 18
)

func (mType AbilityType) String() string {
	names := [25]string{
		"Balance",        // 0
		"Jumping",        // 1
		"Pace",           // 2
		"Acceleration",   // 3
		"Power",          // 4
		"Agility",        // 5
		"Shooting",       // 6
		"Technique",      // 7
		"Dribbling",      // 8
		"LongShots",      // 9
		"Corners",        // 10
		"FreeKickTaking", // 11
		"Throws",         // 12
		"Crossing",       // 13
		"Passing",        // 14
		"Tackling",       // 15
		"Marking",        // 16
		"Heading",        // 17
		"PenaltyTaking",  // 18
	}

	if mType < Balance || mType > PenaltyTaking {
		return "Error"
	}

	return names[mType]
}

type GlobalAbilitytMap map[types.UID]*Ability

var globalAbilityMap = GlobalAbilitytMap{}

type Ability struct {
	// ID represents character or entity that have this ability
	ID          types.UID `json:"id"`
	AbilityList []float32 `json:"ability_list"`
}

func NewAbility(id types.UID, vList []float32) *Ability {
	aList := []float32{}
	for _, val := range vList {
		aList = append(aList, val)
	}

	ability := &Ability{ID: id, AbilityList: aList}
	globalAbilityMap[id] = ability

	return ability
}

func (a *Ability) AddValue(vMap map[AbilityType]float32) {
	for key, val := range vMap {
		a.AbilityList[key] += val
	}
}

func (a *Ability) SubValue(vMap map[AbilityType]float32) {
	for key, val := range vMap {
		a.AbilityList[key] -= val
	}
}
