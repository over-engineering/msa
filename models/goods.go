package models

// Goods interface covers all kinds of goods
type Goods interface {
	GetID() UID
	GetName() string
	GetEffect()
	GetDuration()
	GetPrice() float32
}

type SmartPhone struct {
	ID      UID     `json:"id"`
	Name    string  `json:"name"`
	Brand   string  `json:"brand"`
	Price   float32 `json:"price"`
	Payment float32 `json:"plan"` // 월 요금제
	// Duration
	// Effect
	// TODO: Apps, 통화, 메시지?
}

func (s SmartPhone) GetID() UID {
	return s.ID
}

func (s SmartPhone) GetName() string {
	return s.Name
}

func (s SmartPhone) GetPrice() float32 {
	return s.Price
}

func (s SmartPhone) GetPayment() float32 {
	return s.Payment
}
