package world

import (
	"math"

	"github.com/over-engineering/msa/models/types"
)

type Location struct {
	X int
	Y int
	// We may implement z and affect to cardio :)
}

// func (l *Location) Visit(l2 *Location, st *Status) {
// 	dis := Distance(l, l2)
// 	st.ApplyEffects(
// 		Effects{
// 			Effect{
// 				Target: "Health",
// 				Value:  -30 * dis,
// 			},
// 		})

// 	l2.UpdateDistance(l.X, l.Y)
// }

func Distance(l1 *Location, l2 *Location) float32 {
	return float32(math.Pow(math.Pow(float64(l1.X-l2.X), 2)+math.Pow(float64(l1.Y-l2.Y), 2), 0.5))
}

func (l *Location) UpdateDistance(x int, y int) {
	l.X = x
	l.Y = y
}

type FacilityManager interface {
	GetLocation() *Location
	GetCapacity() int
	// Visit(l *Location, st *Status)
}

type Stadium struct {
	ID       types.UID `json:"id"`
	Location *Location `json:"location"`
	Capacity int       `json:"capacity"`
}

func (s *Stadium) GetLocation() *Location {
	return s.Location
}

func (s *Stadium) GetCapacity() int {
	return s.Capacity
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
