package character

import (
	"fmt"

	"github.com/over-engineering/msa/models/physical"
)

func ExampleNewStatus() {
	ns := NewStatus(float32(50))
	fmt.Println(ns)
	ns.Physical.Muscles = map[physical.MuscleType]*physical.Muscle{
		physical.Biceps: &physical.Muscle{
			Mass:     5,
			RedRatio: 0.5,
			Fatigue:  1,
		},
	}
	ns.Physical.Muscles[physical.Biceps].Fatigue -= 0.1
	fmt.Println(ns)
	fmt.Println(ns.Physical.Muscles[physical.Biceps].Fatigue)
	delete(ns.Physical.Muscles, physical.Biceps)
	ns.Physical.Muscles[physical.CalfMuscle] = &physical.Muscle{
		Mass:     5,
		RedRatio: 0.5,
		Fatigue:  1,
	}
	fmt.Println(ns)
	// Output:
	//
}
