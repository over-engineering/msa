package world

import (
	"errors"
	"fmt"

	"github.com/over-engineering/msa/models/character"
	"github.com/over-engineering/msa/models/types"
)

type Facility struct {
	ID       types.UID `json:"id"`
	Location Location  `json:"location"`
}

func (f Facility) GetID() types.UID {
	return f.ID
}

func (f Facility) GetLocation() Location {
	return f.Location
}

type Stadium struct {
	Facility `json:"facility"`

	Capacity int `json:"capacity"`
}

// GetInfo returns the specific information in string.
// option 0: capacity
func (s Stadium) GetInfo(option int) string {
	switch option {
	case 0:
		cap := fmt.Sprintf("%d", s.Capacity)
		return cap
	default:
		return "No option matching"
	}
}

func (s Stadium) UseFacility(by *character.Character, option int, args []string) error {
	switch option {
	case 0:
		return nil
	default:
		return errors.New("no option matching")
	}
}

// func (s *Stadium) Visit(l *Location, st *Status) {
// 	s.Location.Visit(l, st)
// }

type Market struct {
	ID       types.UID `json:"id"`
	Location *Location `json:"location"`
	Capacity int       `json:"capacity"`
	// GoodsList    []GoodsHelper `json:"goods_list"`
	DiscountRate float32 `json:"discount_rate"` // 0~1
}

func (m *Market) GetLocation() *Location {
	return m.Location
}

func (m *Market) GetCapacity() int {
	return m.Capacity
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
