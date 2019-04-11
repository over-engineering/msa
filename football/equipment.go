package football

import (
	"fmt"

	"github.com/over-engineering/msa/football/types"
)

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

type GlobalEquipmentMap map[types.UID]EquipList

var GEquipmentMap = GlobalEquipmentMap{}

func FindEquipListByCharacterID(id types.UID) (EquipList, error) {
	var eList EquipList
	if eList, ok := GEquipmentMap[id]; ok {
		return eList, nil
	} else {
		// Read from db
	}

	return eList, nil
}

type EquipmentMap map[types.UID]*Equipment

var EqMap = EquipmentMap{}

func FindEquipmentByID(id types.UID) (*Equipment, error) {
	eq := EqMap[id]
	if eq == nil {
		// TODO: Read from db
	}
	return eq, nil
}

type EquipList []*Equipment

func RegisterEquipList(id types.UID) {
	eList := EquipList{
		EQUIP_HEAD:       nil,
		EQUIP_UPPERBODY:  nil,
		EQUIP_LOWERBODY:  nil,
		EQUIP_LEFT_HAND:  nil,
		EQUIP_RIGHT_HAND: nil,
		EQUIP_FOOT:       nil,
	}
	GEquipmentMap[id] = eList
	// TODO: Write to db
}

func UnRegisterEquipList(id types.UID) {
	GEquipmentMap[id] = nil
	// TODO: Delete in db
}

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

func (e *Equipment) Equip(a *Ability, eList EquipList) error {
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

func (e *Equipment) UnEquip(a *Ability, eList EquipList) error {
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
