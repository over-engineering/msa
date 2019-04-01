package models

type GoodsType UID

const (
	PhoneType GoodsType = iota
	CarType
	HouseType
	FoodType
)

const (
	Galaxy        UID = 100
	AppleID       UID = 1000
	OrangeID          = 1001
	EnergyDrinkID     = 1002
	DrugID            = 1003
)

// base goods struct
type Goods struct {
	ID          UID       `json:"id"`
	GoodsType   GoodsType `json:"goods_type"`
	Name        string    `json:"name"`
	Brand       string    `json:"brand"`
	Price       float32   `json:"price"`
	Description string    `json:"description"`
}

func (g *Goods) GetID() UID {
	return g.ID
}

func (g *Goods) GetGoodsType() GoodsType {
	return g.GoodsType
}

func (g *Goods) GetName() string {
	return g.Name
}

func (g *Goods) GetPrice() float32 {
	return g.Price
}

func (g *Goods) SetPrice(p float32) {
	g.Price = p
}

func (g *Goods) GetBrand() string {
	return g.Brand
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

func (s *SmartPhone) UpdateByDay() {
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

type Car struct {
	Goods
	HousePower float32 `json:house_power`
}

func (c *Car) UpdateByDay() {
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

func (f *House) UpdateByDay() {
}

func (h *House) Rest(st *Status) {
	st.ApplyEffects(
		Effects{
			Effect{
				Target: "Health",
				Value:  30,
			},
		})
}

type Food struct {
	*Goods
	Flavor         float32 `json:"flavor"`
	Kcal           float32 `json:"kcal"` // kcal
	Mass           float32 `json:"mass"` // kg
	ExpirationDate int     `json:"expiration_date"`
}

const (
	foodNeedsCoe  = 0.1
	foodFlavorCoe = 0.1
)

var KcalPerMass = map[UID]float32{
	AppleID:       0.0521,
	OrangeID:      0.04,
	EnergyDrinkID: 0.1,
	DrugID:        0.07,
}

func NewFood(id UID, name string, price float32, flavor float32, mass float32, expDate int) *Food {
	return &Food{
		Goods: &Goods{
			ID:        id,
			GoodsType: FoodType,
			Name:      name,
			// Brand: ,
			Price: price,
			// Description: ,
		},
		Flavor:         flavor,
		Kcal:           KcalPerMass[id] * mass,
		Mass:           mass,
		ExpirationDate: expDate,
	}
}

func (f *Food) UpdateByDay() {
	// f.Price *= (float32(f.ExpirationDate) - 1) / float32(f.ExpirationDate)
	if f.ExpirationDate < 0 {
		// f.Price = 0
		f.Flavor *= foodFlavorCoe
	}
	f.ExpirationDate -= 1
}

func (f *Food) Eat(st *Status) {
	if f.ExpirationDate < 0 {
		// TODO: 컨디션 저하
	}

	es := Effects{
		Effect{
			Target: "ConsumedKcal",
			Value:  -f.Kcal * st.KcKgTranslationRate,
		},
		Effect{
			Target: "Needs",
			Value:  Needs{FoodNeeds: -f.Flavor * f.Kcal * foodNeedsCoe},
		},
	}
	st.ApplyEffects(es)
}

func (f *Food) GetExpirationDate() int {
	return f.ExpirationDate
}

type Drug struct {
	*Goods
	Strength       float32 `json:"strength"`
	Kcal           float32 `json:"kcal"` // kcal
	Mass           float32 `json:"mass"` // kg
	ExpirationDate int     `json:"expiration_date"`
}

func NewDrug(id UID, name string, price float32, strength float32, mass float32, expDate int) *Drug {
	return &Drug{
		Goods: &Goods{
			ID:        id,
			GoodsType: FoodType,
			Name:      name,
			// Brand: ,
			Price: price,
			// Description: ,
		},
		Strength:       strength,
		Kcal:           KcalPerMass[id] * mass,
		Mass:           mass,
		ExpirationDate: expDate,
	}
}

func (d *Drug) UpdateByDay() {

}

func (d *Drug) Eat(st *Status) {
	st.ApplyEffects(
		Effects{
			Effect{
				Target: "ConsumedKcal",
				Value:  -d.Kcal * st.KcKgTranslationRate,
			},
			Effect{
				Target: "Hormones",
				// Value:  Hormones{Dopamine: &d.Strength},
				Value: Hormones{Dopamine: d.Strength},
			},
		})
}
