package tests

import (
	"fmt"

	"github.com/over-engineering/msa/models"
)

func Example_mentals() {
	testSet := map[models.MentalType]map[string]float32{
		models.Ambition: map[string]float32{
			"Example1": 50,
			"Example2": 40,
			"Example3": 10,
		},
		models.Boldness: map[string]float32{
			"Example1": 10,
			"Example2": 2,
			"Example3": 3,
		},
		models.Aggression: map[string]float32{
			"Example1": 100,
			"Example2": 100,
			"Example3": 100,
		},
		models.Predictation: map[string]float32{
			"Example1": 100,
			"Example2": 100,
			"Example3": 100,
		},
		models.Composure: map[string]float32{
			"Example1": 100,
			"Example2": 100,
			"Example3": 100,
		},
		models.Concentration: map[string]float32{
			"Example1": 100,
			"Example2": 100,
			"Example3": 100,
		},
		models.Immersion: map[string]float32{
			"Example1": 100,
			"Example2": 100,
			"Example3": 100,
		},
		models.Competition: map[string]float32{
			"Example1": 100,
			"Example2": 100,
			"Example3": 100,
		},
		models.SelfEsteem: map[string]float32{
			"Example1": 100,
			"Example2": 100,
			"Example3": 100,
		},
		models.Confidence: map[string]float32{
			"Example1": 100,
			"Example2": 100,
			"Example3": 100,
		},
		models.Attention: map[string]float32{
			"Example1": 100,
			"Example2": 100,
			"Example3": 100,
		},
	}

	mentalMap, _ := models.NewMentalMap(testSet)

	fmt.Println(mentalMap)
	// Output:
	// map[Ambition:{Ambition 26} Boldness:{Boldness 3.3} Aggression:{Aggression 100} Predictation:{Predictation 100} Composure:{Composure 100} Concentration:{Concentration 100} Immersion:{Immersion 100} Competition:{Competition 100} SelfEsteem:{SelfEsteem 100} Confidence:{Confidence 100} Attention:{Attention 100}]
}
