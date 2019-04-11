package ability

import (
	"github.com/over-engineering/msa/models/character"
	"github.com/over-engineering/msa/models/types"
)

// CommonAbility represents common abilities that could be used through
// all kinds of job.
type FootballAbilityType types.Type

const (
	Balance        FootballAbilityType = iota // 0
	Jumping                                   // 1
	Pace                                      // 2
	Acceleration                              // 3
	Power                                     // 4
	Agility                                   // 5
	Shooting                                  // 6
	Technique                                 // 7
	Dribbling                                 // 8
	LongShots                                 // 9
	Corners                                   // 10
	FreeKickTaking                            // 11
	Throws                                    // 12
	Crossing                                  // 13
	Passing                                   // 14
	Tackling                                  // 15
	Marking                                   // 16
	Heading                                   // 17
	PenaltyTaking                             // 18
)

func (mType FootballAbilityType) String() string {
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

type FootballAbility struct {
	// ID represents character or entity that have this ability
	ID            types.UID `json:"id"`
	AbilityIvList []float32 `json:"ability_iv_list"`
}

func UpdateFootballAbility(id types.UID, s *character.Status) FootballAbility {
	// TODO: Update football ability depending on character status
	return FootballAbility{
		ID:            id,
		AbilityIvList: []float32{50, 50, 50, 50, 50, 50, 50, 50, 50, 50, 50, 50, 50, 50, 50, 50, 50, 50, 50},
	}
}
