package space

import (
	"time"

	"github.com/over-engineering/msa/models/finance"

	"github.com/over-engineering/msa/models/types"
)

type FacilityStatusType types.Type

const (
	Available FacilityStatusType = iota // 0
	UnderConstruction
	TemporarySuspension
	Vacation
)

func (fs FacilityStatusType) String() string {
	switch fs {
	case Available:
		return "available"
	case UnderConstruction:
		return "under construction"
	case TemporarySuspension:
		return "temporary suspension"
	case Vacation:
		return "vacation"
	default:
		return "facility status: not defined status type"
	}
}

type Facility struct {
	ID          types.UID          `json:"id"`
	Name        string             `json:"name"`
	Owner       types.UID          `json:"owner"`
	Location    Location           `json:"location"`
	Constructed time.Time          `json:"constructed"`
	Status      FacilityStatusType `json:"status"`
	Value       finance.Dollars    `json:"value"`
}

type Building struct {
	Facility `json:"facility"`

	Floor      int     `json:"floor"`
	Horizontal float32 `json:"horizontal"`
	Vertical   float32 `json:"vertical"`
}

type Stadium struct {
	Facility `json:"facility"`

	TeamID      types.UID   `json:"team_id"`
	Capacity    int         `json:"capacity"`
	Trainings   []types.UID `json:"trainings"`
	AvailEvents []types.UID `json:"avail_events"`
}

type Market struct {
	Facility `json:"facility"`

	Goods map[types.CID]int `json:"goods_ids"`
	// SellerIDs    []types.UID `json:"seller_ids"`
	DiscountRate float32       `json:"discount_rate"` // 0~1
	ShoppingTime time.Duration `json:"shopping_time"`
}

// func (m *Market) Visit(l *Location, st *Status) {
// 	m.Location.Visit(l, st)
// }

// func (m *Market) GetGoodsList() []GoodsHelper {
// 	return m.GoodsList
// }

// func (m *Market) BuyGoods(index int) GoodsHelper {
// 	goods := m.GoodsList[index]
// 	m.SubGoods(index)
// 	return goods
// }

// func (m *Market) SellGoods(g GoodsHelper) {
// 	m.AddGoods(g)
// }

// func (m *Market) AddGoods(g GoodsHelper) {
// 	g.SetPrice(g.GetPrice() * (1 - m.DiscountRate))
// 	m.GoodsList = append(m.GoodsList, g)
// }

// func (m *Market) SubGoods(index int) {
// 	m.GoodsList = append(m.GoodsList[:index], m.GoodsList[index+1:]...)
// }

type Hospital struct {
	ID       types.UID `json:"id"`
	Location *Location `json:"location"`
	Capacity int       `json:"capacity"`
	Quality  float32   `json:"quality"`
}

func (h *Hospital) GetLocation() *Location {
	return h.Location
}

func (h *Hospital) GetCapacity() int {
	return h.Capacity
}

// func (h *Hospital) Visit(l *Location, st *Status) {
// 	h.Location.Visit(l, st)
// }

// func (h *Hospital) Treat(st *Status, cond Conditions) {

// }

// func (h *Hospital) Surgery(st *Status, cond Conditions) {

// }

type Fitness struct {
	ID       types.UID `json:"id"`
	Location *Location `json:"location"`
	Capacity int       `json:"capacity"`
	Quality  float32   `json:"quality"`
}

func (f *Fitness) GetLocation() *Location {
	return f.Location
}

func (f *Fitness) GetCapacity() int {
	return f.Capacity
}

// func (f *Fitness) Visit(l *Location, st *Status) {
// 	f.Location.Visit(l, st)
// }

// func (f *Fitness) Exercise(st *Status) {
// 	st.ApplyEffects(
// 		Effects{
// 			Effect{
// 				Target: "Health",
// 				Value:  float32(-10),
// 			},
// 			Effect{
// 				Target: "ConsumedKcal",
// 				Value:  float32(200) * f.Quality / 100,
// 			},
// 		})
// }
