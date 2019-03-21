package models

type GoodsIndex int

const (
	PhoneIndex GoodsIndex = iota
	CarIndex
	HouseIndex
	FoodIndex
)

// Goods interface covers all kinds of goods
// type Goods interface {
type GoodsIt interface {
	GetID() interface{}
	GetName() string
	// GetEffect()
	// GetDuration()
	GetBrand() string
	GetPrice() float32

	// RegisterEffects(s *Status)
	// ApplyEffects()
	RegisterGoods(c *Character)
	UnRegisterGoods(c *Character)
}

// base goods struct
type Goods struct {
	ID    interface{} `json:"id"`
	Name  string      `json:"name"`
	Brand string      `json:"brand"`
	Price float32     `json:"price"`
	// Duration    float32 `json:"duration"`
	Description string  `json:"description"`
	Effects     Effects `json:"effects"`
}

func (g Goods) GetID() interface{} {
	return g.ID
}

func (g Goods) GetName() string {
	return g.Name
}

func (g Goods) GetPrice() float32 {
	return g.Price
}

func (g Goods) GetBrand() string {
	return g.Brand
}

// func (g Goods) ApplyEffects() {
// 	g.Effects.ApplyEffects()
// }

type Camera struct {
	Goods
}

type SmartPhoneID int

const (
	Galaxy SmartPhoneID = iota
)

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

func (s SmartPhone) GetPayment() float32 {
	return s.Payment
}

func (s *SmartPhone) Call(st *Status) {
	st.ApplyEffects(
		Effects{
			Effect{
				Target: "UpdateBalance",
				Value:  50 + s.Payment*0.01,
			},
		})

}

func (s *SmartPhone) RegisterGoods(c *Character) {
	c.Goods[PhoneIndex][s.ID] = s
}

func (s *SmartPhone) UnRegisterGoods(c *Character) {
	c.Goods[PhoneIndex][s.ID] = nil
}

type Car struct {
	Goods
	HousePower float32 `json:house_power`
}

func (car *Car) RegisterGoods(c *Character) {
}

func (car *Car) UnRegisterGoods(cr *Character) {
}

func (car *Car) Ride(st *Status) {
	st.ApplyEffects(
		Effects{
			Effect{
				Target: "UpdateCharm",
				Value:  10 + car.HousePower,
			},
			Effect{
				Target: "UpdateBalance",
				Value:  10 + car.HousePower,
			},
		})
}

type House struct {
	Goods
	// 면적 단위
	Acreage      float32  `json:acreage`
	Location     Location `json:place`
	ParkingSpace bool     `json:parking_space`
}

// type FoodHelper interface {
// 	Eat()
// 	GetExpirationDate()
// }

type FoodID int

const (
	AppleID FoodID = iota
	OrangeID
	EnergyDrinkID
)

type Food struct {
	Goods
	Flavor         float32
	Kcal           float32 // kcal
	Mass           float32 // kg
	ExpirationDate float32
}

const (
	AppleKcalPerMass = 0.0521
	OrangeKcalPerMass
)

func NewFood(id FoodID, name string, price float32, mass float32) *Food {
	return &Food{
		Goods: Goods{
			ID:   id,
			Name: name,
			// Brand: ,
			Price: 100,
			// Description: ,
			// Effects     ,
		},
		Flavor:         50,
		Kcal:           AppleKcalPerMass * mass,
		Mass:           mass,
		ExpirationDate: 100,
	}
}

func (f *Food) Eat(st *Status) {
	st.ApplyEffects(
		Effects{
			Effect{
				Target: "ConsumedKcal",
				Value:  -f.Kcal * st.KcKgTranslationRate,
			},
		})
}

func (f *Food) RegisterGoods(c *Character) {

	// if c.Goods[FoodIndex] == nil {
	// 	c.Goods[FoodIndex] = make(map[interface{}]interface{})
	// }
	c.Goods[FoodIndex][f.ID] = f
}

func (f *Food) UnRegisterGoods(c *Character) {
	c.Goods[FoodIndex][f.ID] = nil
}

func (f *Food) GetExpirationDate() float32 {
	return f.ExpirationDate
}
