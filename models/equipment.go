package models

// Equipment interface represents interface for equipments.
// All equipments should have Type, Effect and Durability.
type Equipment interface {
	GetType() EquipmentType
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

type HeadGear struct {
	ID      UID      `json:"id"`
	Name    string   `json:"name"`
	Price   float32  `json:"price"`
	Effects []Effect `json:"effect"`
	// Duration
}

func (h *HeadGear) GetType() EquipmentType {
	return EQUIP_HEAD
}

func (h *HeadGear) GetID() UID {
	return h.ID
}

func (h *HeadGear) GetName() string {
	return h.Name
}

func (h *HeadGear) GetPrice() float32 {
	return h.Price
}

// 고칠 수 있는 것 vs 없는 것 먹는거 장비 뭐 이런것들
// 엄청 구체적인 예시들을 각각 모델링하고나서 생각해보는 걸로
