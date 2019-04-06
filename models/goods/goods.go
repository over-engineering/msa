package goods

import (
	"errors"
	"time"

	"github.com/over-engineering/msa/models/types"
	"github.com/over-engineering/msa/models/user"
	"github.com/over-engineering/msa/models/world/space"
)

type GlobalGoodsMap map[types.CID]*Goods

type GlobalUniqueInfoMap map[types.CID]UniqueInfos

// Goods represents the basic structure of all kinds of goods.
type Goods struct {
	ID    types.CID `json:"id"`
	Name  string    `json:"name"`
	Brand string    `json:"brand"`
	Price float32   `json:"price"`
}

// InnerCameraInfos represents the pixel infos of camera.
type InnerCameraInfos struct {
	ID         types.CID `json:"id"`
	Horizontal int       `json:"horizontal"`
	Vertical   int       `json:"vertical"`
}

// SmartPhoneInfos represents unique information of smartphone.
type SmartPhoneInfos struct {
	ID               types.CID `json:"id"`
	Width            float32   `json:"width"`
	Height           float32   `json:"height"`
	Memory           float32   `json:"memory"`
	InnerCameraInfos `json:"inner_camera_infos"`
}

// SmartPhone represents dynamic informations about smartphone object.
type SmartPhone struct {
	ID          types.UID   `json:"id"`
	CommonID    types.CID   `json:"common_id"`
	ContractID  types.UID   `json:"contract_id"`
	Bought      time.Time   `json:"bought"`
	Durability  float32     `json:"durability"`
	Battery     float32     `json:"battery"`
	ContactList []types.UID `json:"contact_list"`
}

// Call ...
func (s *SmartPhone) Call() {
}

// CarInfos represents unique informations about Car object.
type CarInfos struct {
	FuelEfficiency float32 `json:"fuel_efficiency"`
	MaxFuelLevel   float32 `json:"max_fuel_level"`
}

// Car represents dynamic informations about car object.
type Car struct {
	ID             types.UID `json:"id"`
	CommonID       types.CID `json:"common_id"`
	ContractID     types.UID `json:"contract_id"`
	Bought         time.Time `json:"bought"`
	DrivenDistance float32   `json:"driven_distance"`
	Durability     float32   `json:"durability"`
	FuelLevel      float32   `json:"fuel_level"`     // 임시
	MaxFuelLevel   float32   `json:"max_fuel_level"` // 임시
	FuelEfficiency float32   `json:"fuel_efficiency"`
}

// Ride functino updates car's fuel level.
func (c *Car) Ride(distance float32) error {
	fc := distance / c.FuelEfficiency
	if c.FuelLevel <= fc {
		return errors.New("not enough fuel level")
	}
	c.FuelLevel -= fc
	c.DrivenDistance += distance
	// TODO: Maybe update logic should be here.
	return nil
}

// HouseInfos represents unique informations about House object.
type HouseInfos struct {
	ID types.CID `json:"id"`
	// 면적 단위
	Acreage        float32        `json:"acreage"`
	Location       space.Location `json:"place"`
	ParkingNumber  int            `json:"parking_number"`
	RestEfficiency float32        `json:"rest_efficiency"`
}

// House represents dynamic informations about House object.
type House struct {
	ID         types.UID `json:"uid"`
	CommonID   types.CID `json:"common_id"`
	ContractID types.UID `json:"contract_id"`
	Bought     time.Time `json:"bought"`
}

// Rest updates character's health with efficiency and hours.
func (h House) Rest(hi HouseInfos, c *user.Character, hours float32) {
	c.Status.Health += hi.RestEfficiency * float32(hours) * c.Status.Physical.Resilience
}

type FoodInfos struct {
	ID             types.CID `json:"id"`
	Flavor         float32   `json:"flavor"`
	KcalPerMass    float32   `json:"kcal_per_mass"`
	Expiration     float32   `json:"expiration"`
	ExpirationRate float32   `json:"expiration_rate"`
}

type Food struct {
	ID             types.UID `json:"id"`
	CommonID       types.CID `json:"common_id"`
	Kcal           float32   `json:"kcal"` // kcal
	Mass           float32   `json:"mass"` // kg
	Expiration     float32   `json:"expiration"`
	ExpirationRate float32   `json:"expiration_rate"` // 임시
	Refrigerated   float32   `json:"refrigerated"`
}

func (f *Food) UpdateByDay() error {
	f.Expiration -= f.ExpirationRate * f.Refrigerated
	return nil
}

// func (f *Food) Eat() {

// }

// Drug represents the dynamic informations about the drug object.
type Drug struct {
	ID         types.UID `json:"id"`
	CommonID   types.CID `json:"common_id"`
	Strength   float32   `json:"strength"`
	Kcal       float32   `json:"kcal"` // kcal
	Mass       float32   `json:"mass"` // kg
	Expiration float32   `json:"expiration"`
}

// func NewDrug(id types.UID, name string, price float32, strength float32, mass float32, expDate int) *Drug {
// 	return &Drug{
// 		Goods: &Goods{
// 			ID:        id,
// 			GoodsType: FoodType,
// 			Name:      name,
// 			// Brand: ,
// 			Price: price,
// 			// Description: ,
// 		},
// 		Strength:       strength,
// 		Kcal:           KcalPerMass[string(id)] * mass,
// 		Mass:           mass,
// 		ExpirationDate: expDate,
// 	}
// }

// func (d *Drug) Eat() {

// }
