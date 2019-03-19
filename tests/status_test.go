package tests

import (
	"fmt"
	"testing"

	"github.com/over-engineering/msa/models"
)

/*
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
*/

func TestCreateStatus(t *testing.T) {
	// create status, intelligence, personality, mental
	status := models.Status{}

	// intelligence
	intellValue := [8]float32{10, 30, 40, 100, 50, 5, 7, 23}
	intellPeakAge := [8]int{30, 30, 30, 30, 30, 30, 30, 30}
	intellVariation := [8]float32{0.1, 0.1, 0.1, 0.1, 0.1, 0.1, 0.1, 0.1}
	status.Intelligences = models.NewIntelligences(intellValue, intellPeakAge, intellVariation)

	// persoanlity
	status.Personality = models.NewPersonality([4]float32{0.5, 0.4, 0.1, 1.0})

	// mental
	status.Mentals = CreateMentals()
	// personalilty-mentalcoefficient
	pMentalMap := map[interface{}]float32{
		models.Ambition:      1.6,
		models.Boldness:      1.2,
		models.Aggression:    0.9,
		models.Predictation:  1.11,
		models.Composure:     1.3,
		models.Concentration: 1.2,
		models.Immersion:     1.5,
		models.Competition:   0.98,
		models.SelfEsteem:    0.9,
		models.Confidence:    0.8,
		models.Attention:     1.0,
	}
	PMentalList := [16]map[interface{}]float32{}
	for i := 0; i < 16; i++ {
		PMentalList[i] = pMentalMap
	}
	models.PTable.AddPersonalityCoe(PMentalList)

	// PrintStatus(&status)
	// update
	// UpdateStatus(&status)
	// PrintStatus(&status)
}

// func UpdateStatus(status *models.Status) {

// 	es := []models.Effect{
// 		// intelligence effects
// 		models.Effect{"Intelligences", map[models.IntelligenceType]float32{models.Interpersonal: float32(20)}},
// 		models.Effect{"Intelligences", map[models.IntelligenceType]float32{models.Intrapersonal: float32(-20)}},

// 		// personalilty effects
// 		models.Effect{"Personality", models.Personality{Extraversion: 2, Sensing: 3, Thinking: -3, Judging: 20}},

// 		// mental effects
// 		models.Effect{"Mentals", models.Mentals{models.Ambition: &models.Mental{models.Ambition, float32(33)}}},
// 		models.Effect{"Mentals", models.Mentals{models.Competition: &models.Mental{models.Competition, float32(-20)}}},
// 	}
// 	status.ApplyEffect(es)

// }

func PrintStatus(status *models.Status) {
	for i := 0; i < models.IntelligenceNum; i++ {
		fmt.Println("Intelligence: ", status.Intelligences[models.IntelligenceType(i)])
	}

	fmt.Println("Personality: ", status.Personality)

	for i := 0; i < models.MentalNum; i++ {
		fmt.Println("Mental: ", status.Mentals[models.MentalType(i)])
	}
}
