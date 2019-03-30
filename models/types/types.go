package types

// Type represents the enum type for various things in this game.
type Type int

// UID represents the unique ID for various types.
type UID string

// JobType is an enum for various job types.
type JobType Type

// EquipmentType is an enum for equipment types.
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
