package tests

import (
	"fmt"
	"reflect"

	"github.com/over-engineering/msa/models"
)

func Example() {
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
	status := &models.Status{
		Health:  float32(98),
		Mentals: mentalMap,
	}
	fmt.Println(status.Mentals[models.Ambition])

	r := reflect.ValueOf(status)
	f := reflect.Indirect(r).FieldByName("Mentals")
	fmt.Println(reflect.TypeOf(mentalMap))
	fmt.Println(f)

	e := models.Effect{"Mentals", models.Mentals{
		models.Ambition: &models.Mental{models.Ambition, float32(1)},
	}}

	status.ApplyEffect([]models.Effect{e})
	fmt.Println(status.Mentals[models.Ambition])
	// Output:
	//
}
