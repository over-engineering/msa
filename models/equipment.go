package models

// Equipment interface represents interface for equipments.
// All equipments should have Type, Effect and Durability.
type Equipment interface {
	Type() EquipmentType
	Effect()
	Durability()
}

// EquipmentType is enum for various equipment types.
type EquipmentType Type

const (
	EQUIP_HEAD       EquipmentType = iota // 0
	EQUIP_UPPERBODY                       // 1
	EQUIP_LOWERBODY                       // 2
	EQUIP_LEFT_HAND                       // 3
	EQUIP_RIGHT_HAND                      // 4
	EQUIP_LEFT_FOOT                       // 5
	EQUIP_RIGHT_FOOT                      // 6
)
