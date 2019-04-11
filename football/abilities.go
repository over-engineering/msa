package football

import (
	"fmt"

	"github.com/over-engineering/msa/football/types"
)

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

func FindAbilityByID(id types.UID) (*Ability, error) {
	a := globalAbilityMap[id]

	if a == nil {
		// TODO: Read from db
	}
	return a, nil
}

func RegisterAbility(ability *Ability) error {
	a, err := FindAbilityByID(ability.ID)
	if err != nil {
		return err
	}
	if a != nil {
		return fmt.Errorf("Ability", ability.ID, "exist")
	}

	ability.AbilityList = []float32{}
	for _, val := range ability.AbilityIvList {
		ability.AbilityList = append(ability.AbilityList, val)
	}

	globalAbilityMap[ability.ID] = ability

	// TODO: Write to db
	return nil
}

func UpdateAbility(ability *Ability) error {
	a, err := FindAbilityByID(ability.ID)
	if err != nil {
		return err
	}
	if a == nil {
		return fmt.Errorf("Ability", ability.ID, "does not exist")
	}

	for i, val := range a.AbilityIvList {
		a.AbilityList[i] = a.AbilityList[i] - val + ability.AbilityIvList[i]
		a.AbilityIvList[i] = ability.AbilityIvList[i]

	}
	// TODO: Write to db

	return nil
}

func UnRegisterAbility(id types.UID) {
	globalAbilityMap[id] = nil
	// TODO: Delete ability in db
}

type Ability struct {
	// ID represents character or entity that have this ability
	ID            types.UID `json:"id"`
	AbilityIvList []float32 `json:"ability_iv_list"`
	AbilityList   []float32 `json:"ability_list"`
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
