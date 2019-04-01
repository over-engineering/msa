package goods

import (
	"github.com/over-engineering/msa/models/types"
	"github.com/over-engineering/msa/models/world/space"
)

type GoodsType types.Type

const (
	SmartPhoneType = iota + 1 // 1
	CarType                   // 2
	FoodType                  // 3
	HouseType                 // 4
)

type OwnedList map[GoodsType][]GoodsManager

// Goods represents the basic structure of all kinds of goods.
type Goods struct {
	ID          types.UID `json:"id"`
	GoodsType   GoodsType `json:"goods_type"`
	Name        string    `json:"name"`
	Brand       string    `json:"brand"`
	Price       float32   `json:"price"`
	Description string    `json:"description"`
}

func (g *Goods) GetID() types.UID {
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
	Goods `json:"goods"` // embedded struct

	ContractID types.UID `json:"contract_id"`
	Camera     *Camera   `json:"camera"`
	Memory     float32   `json:"memory"`
	Battery    float32   `json:"battery"`
	// Duration
	// Effect
	// TODO: Apps, 통화, 메시지?
}

func (s *SmartPhone) Call() {
}

type Car struct {
	Goods
	HorsePower float32 `json:"horse_power"`
}

func (car *Car) Ride() {

}

type House struct {
	Goods
	// 면적 단위
	Acreage      float32        `json:acreage`
	Location     world.Location `json:place`
	ParkingSpace bool           `json:parking_space`
}

func (h *House) Rest() {

}

type Food struct {
	*Goods
	Flavor         float32 `json:"flavor"`
	Kcal           float32 `json:"kcal"` // kcal
	Mass           float32 `json:"mass"` // kg
	ExpirationDate int     `json:"expiration_date"`
}

var KcalPerMass = map[string]float32{
	"AppleID":       0.0521,
	"OrangeID":      0.04,
	"EnergyDrinkID": 0.1,
	"DrugID":        0.07,
}

func NewFood(id types.UID, name string, price float32, flavor float32, mass float32, expDate int) *Food {
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
		Kcal:           KcalPerMass[string(id)] * mass,
		Mass:           mass,
		ExpirationDate: expDate,
	}
}

func (f *Food) Eat() {

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

func NewDrug(id types.UID, name string, price float32, strength float32, mass float32, expDate int) *Drug {
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
		Kcal:           KcalPerMass[string(id)] * mass,
		Mass:           mass,
		ExpirationDate: expDate,
	}
}

func (d *Drug) Eat() {

}
