package status

import (
	"fmt"
)

func ExampleNewStatus() {
	ns := NewStatus()
	fmt.Println(ns)
	ns.Physical.Muscles = map[MuscleType]*Muscle{
		Biceps: &Muscle{
			Mass:     5,
			RedRatio: 0.5,
			Fatigue:  1,
		},
	}
	ns.Physical.Muscles[Biceps].Fatigue -= 0.1
	fmt.Println(ns)
	fmt.Println(ns.Physical.Muscles[Biceps].Fatigue)
	delete(ns.Physical.Muscles, Biceps)
	ns.Physical.Muscles[CalfMuscle] = &Muscle{
		Mass:     5,
		RedRatio: 0.5,
		Fatigue:  1,
	}
	fmt.Println(ns)
	// Output:
	//
}
