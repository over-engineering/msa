package models

// Goods interface covers all kinds of goods
// type Goods interface {
type GoodsIt interface {
	GetID() UID
	GetName() string
	// GetEffect()
	GetDuration()
	GetPrice() float32
	RegisterEffects(c *Character)
	ApplyEffects()
}

// base goods struct
type Goods struct {
	ID          UID     `json:"id"`
	Name        string  `json:"name"`
	Brand       string  `json:"brand"`
	Price       float32 `json:"price"`
	Duration    float32 `json:"duration"`
	Description string  `json:"description"`
	Effects     Effects `json:"effects"`
}

func (g *Goods) GetID() UID {
	return g.ID
}

func (g *Goods) GetName() string {
	return g.Name
}

func (g *Goods) GetPrice() float32 {
	return g.Price
}

func (g *Goods) GetDuration() {

}

func (g *Goods) ApplyEffects() {
	g.Effects.ApplyEffects()
}

type Camera struct {
	Goods
}

type SmartPhone struct {
	Goods           // embedded struct
	Payment float32 `json:"plan"` // 월 요금제
	Camera  *Camera `json:"camera`
	Memory  float32 `json:memory`
	Battery float32 `json:battery`
	// Duration
	// Effect
	// TODO: Apps, 통화, 메시지?
}

// func (s SmartPhone) GetID() UID {
// 	return s.ID
// }

// func (s SmartPhone) GetName() string {
// 	return s.Name
// }

// func (s SmartPhone) GetPrice() float32 {
// 	return s.Price
// }

func (s SmartPhone) GetPayment() float32 {
	return s.Payment
}

func (s *SmartPhone) RegisterEffects(c *Character) {
	s.Goods.Effects = Effects{
		Effect{
			Target: "UpdateMentals",
			Value:  c.Status.Mentals.UpdateValue,
			Parameters: []interface{}{
				Mentals{
					Confidence: &Mental{Confidence, float32(20)},
					Ambition:   &Mental{Ambition, float32(5)},
				},
				c.Status.Personality,
			},
		},
		Effect{
			Target:     "AddAllFriendships",
			Value:      AddAllFriendships,
			Parameters: []interface{}{c, float32(10)},
		},
	}
	// c.Friendships
}

type Car struct {
	Goods
	HousePower float32 `json:house_power`
}

type House struct {
	Goods
	// 면적 단위
	Acreage      float32  `json:acreage`
	Location     Location `json:place`
	ParkingSpace bool     `json:parking_space`
}
