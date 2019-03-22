package models

type GoodsType UID

const (
	PhoneType GoodsType = iota
	CarType
	HouseType
	FoodType
)

// Goods interface covers all kinds of goods
// type Goods interface {
type GoodsHelper interface {
	GetID() UID
	GetGoodsType() GoodsType
	GetName() string
	GetPrice() float32
	SetPrice(p float32)
}

func GetGoods(goodsList [][]GoodsHelper, goodsType GoodsType, goodsIndex int) GoodsHelper {
	return goodsList[goodsType][goodsIndex]
}

func RegisterGoods(goodsList [][]GoodsHelper, g GoodsHelper) {
	goodsList[g.GetGoodsType()] = append(goodsList[g.GetGoodsType()], g)
}

func UnRegisterGoods(goodsList [][]GoodsHelper, goodsType GoodsType, index int) {
	goodsList[goodsType] = append(goodsList[goodsType][:index], goodsList[goodsType][index+1:]...)
}

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

// type SmartPhoneID int

const (
	Galaxy UID = 100
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

type Car struct {
	Goods
	HousePower float32 `json:house_power`
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

const (
	AppleID UID = 1000
	OrangeID
	EnergyDrinkID
)

type Food struct {
	*Goods
	Flavor         float32
	Kcal           float32 // kcal
	Mass           float32 // kg
	ExpirationDate int
}

const (
	AppleKcalPerMass = 0.0521
	OrangeKcalPerMass
)

func NewFood(id UID, name string, price float32, flavor float32, mass float32) *Food {
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

func (f *Food) GetExpirationDate() int {
	return f.ExpirationDate
}
