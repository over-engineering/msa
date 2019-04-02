package football

import (
	"fmt"

	"github.com/over-engineering/msa/football/types"
)

// import "go/types"

// import "github.com/over-engineering/msa/models/types"

// Equipment interface represents interface for equipments.
// All equipments should have Type, Effect and Durability.
// type Equipment interface {
// 	GetType() EquipmentType
// }

// EquipmentType is enum for various equipment types.
type EquipmentType types.Type

const (
	EQUIP_HEAD       EquipmentType = iota // 0
	EQUIP_UPPERBODY                       // 1
	EQUIP_LOWERBODY                       // 2
	EQUIP_LEFT_HAND                       // 3
	EQUIP_RIGHT_HAND                      // 4
	EQUIP_FOOT                            // 5
)

type GlobalEquipmentMap map[types.UID]EquipmentList

var GEquipmentMap = GlobalEquipmentMap{}

func FindEquipmentListByCharacterID(id types.UID) EquipmentList {
	var eList EquipmentList
	if eList, ok := GEquipmentMap[id]; ok {
		return eList
	} else {
		// Read from db
	}

	return eList
}

type EquipmentMap map[types.UID]*Equipment

var EqMap = EquipmentMap{}

type EquipmentList []*Equipment

func MakeNewEquipmentList() EquipmentList {
	return EquipmentList{
		EQUIP_HEAD:       nil,
		EQUIP_UPPERBODY:  nil,
		EQUIP_LOWERBODY:  nil,
		EQUIP_LEFT_HAND:  nil,
		EQUIP_RIGHT_HAND: nil,
		EQUIP_FOOT:       nil,
	}
}

// type EquipmentList struct {
// 	// ID represents character or entity that have this ability
// 	ID         types.UID                    `json:"id"`
// 	Equipments map[EquipmentType]*Equipment `json:"equipments"`
// }
type Equipments struct {
	Equipments []Equipment `json:"equipments"`
}

type Equipment struct {
	ID            types.UID               `json:"id"`
	EquipmentType EquipmentType           `json:"equipment_type"`
	TargetAbility map[AbilityType]float32 `json:"target_ability"`
}

func RegisterEquipments(es []Equipment) {
	for i, e := range es {
		EqMap[e.ID] = &es[i]
		// newEq := &Equipment{ID: e.ID, EquipmentType: e.EquipmentType, TargetAbility: e.TargetAbility}
		// EqMap[e.ID] = newEq
		// fmt.Println(e)
	}

	// TODO: equipment to db server
}

func (e *Equipment) Equip(a *Ability, eList EquipmentList) error {
	// if a.ID != eList.ID {
	// 	return fmt.Errorf("Diffrent Ability, EquipmentList ID")
	// }

	if val := eList[e.EquipmentType]; val != nil {
		return fmt.Errorf("Equip exist")
	}

	a.AddValue(e.TargetAbility)

	eList[e.EquipmentType] = e

	return nil
}

func (e *Equipment) UnEquip(a *Ability, eList EquipmentList) error {
	// if a.ID != eList.ID {
	// 	return fmt.Errorf("Diffrent Ability, EquipmentList ID")
	// }

	if val := eList[e.EquipmentType]; val == nil {
		return fmt.Errorf("Equip doesn't exist")
	}

	a.SubValue(e.TargetAbility)

	eList[e.EquipmentType] = nil
	// delete(eList, e.EquipmentType)

	return nil
}
