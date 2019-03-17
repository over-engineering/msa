package models

type Location struct {
	x int
	y int
	// We may implement z and affect to cardio :)
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
