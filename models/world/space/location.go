package space

import (
	"math"
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

func Distance(l1 Location, l2 Location) float32 {
	return float32(math.Pow(math.Pow(float64(l1.X-l2.X), 2)+math.Pow(float64(l1.Y-l2.Y), 2), 0.5))
}

func (l *Location) UpdateDistance(x int, y int) {
	l.X = x
	l.Y = y
}
