package tests

import (
	"fmt"

	"github.com/over-engineering/msa/models/world/space"
)

func ExampleFacility() {
	fm := world.FacilityManager(world.Stadium{
		Facility: world.Facility{
			ID: "id",
			Location: world.Location{
				X: 50,
				Y: 50,
			},
		},
		Capacity: 10000,
	})

	fmt.Println(fm.GetInfo(0))
	// Output:
	// 10000
}
