package models

import "math"

type Location struct {
	x int
	y int
	// We may implement z and affect to cardio :)
}

func Distance(l1 Location, l2 Location) float32 {
	return float32(math.Pow(math.Pow(float64(l1.x-l2.x), 2)+math.Pow(float64(l1.x-l2.x), 2), 0.5))
}

type Facility interface {
	GetLocation() Location
	GetCapacity() int
}

type Stadium struct {
	ID       UID      `json:"id"`
	Location Location `json:"location"`
	Capacity int      `json:"capacity"`
}

func (s *Stadium) GetLocation() Location {
	return s.Location
}

func (s *Stadium) GetCapacity() int {
	return s.Capacity
}

type Hospital struct {
	ID       UID      `json:"id"`
	Location Location `json:"location"`
	Capacity int      `json:"capacity"`
	Quality  int      `json:"quality"`
}

func (h *Hospital) GetLocation() Location {
	return h.Location
}

func (h *Hospital) GetCapacity() int {
	return h.Capacity
}

func (h *Hospital) Treat(st *Status, cond Conditions) {

}

func (h *Hospital) Surgery(st *Status, cond Conditions) {

}

type Fitness struct {
	ID       UID      `json:"id"`
	Location Location `json:"location"`
	Capacity int      `json:"capacity"`
	Quality  int      `json:"quality"`
}

func (f *Fitness) GetLocation() Location {
	return f.Location
}

func (f *Fitness) GetCapacity() int {
	return f.Capacity
}
