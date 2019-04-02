package football

import (
	"fmt"

	"github.com/over-engineering/msa/football"
	"github.com/over-engineering/msa/football/types"
)

// func TestAbility2(t *testing.T) {
func Ability() {
	user0 := types.UID("User00")
	ability := &football.Ability{
		ID:          user0,
		AbilityList: []float32{50, 50, 50, 50, 50, 50, 50, 50, 50, 50, 50, 50, 50, 50, 50, 50, 50, 50, 50, 50},
	}

	globalEqList := football.GlobalEquipmentMap{}
	globalEqList[user0] = football.MakeNewEquipmentList()
	// eList := football.EquipmentList{
	// 	ID:         "User00",
	// 	Equipments: map[football.EquipmentType]*football.Equipment{},
	// }

	footballBoot0 := types.UID("FootballBoot00")
	footballBoot := &football.Equipment{
		ID:            footballBoot0,
		EquipmentType: football.EQUIP_FOOT,
		TargetAbility: map[football.AbilityType]float32{football.Shooting: 10, football.Dribbling: 5, football.Crossing: 5, football.Passing: 5},
	}

	fmt.Println(ability, globalEqList[user0])
	footballBoot.Equip(ability, globalEqList[user0])
	// globalEqList[user0][football.EQUIP_LEFT_FOOT] = footballBoot
	fmt.Println(ability, globalEqList[user0])
	footballBoot.UnEquip(ability, globalEqList[user0])
	fmt.Println(ability, globalEqList[user0])

	// tr0 := football.Training{
	// 	ID:            "training00",
	// 	TargetAbility: map[football.AbilityType]float32{football.Shooting: 10, football.Dribbling: 10, football.Crossing: 10, football.Passing: 10},
	// 	// Level:         1,
	// 	// Time:          1,
	// }

	// tr0.TakeTraining(ability)
	fmt.Println(ability)
}
