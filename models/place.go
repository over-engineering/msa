package models

type Location struct {
	x int
	y int
	// We may implement z and affect to cardio :)
}

type Facility interface {
	Location() Location
}
