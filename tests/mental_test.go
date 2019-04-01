package tests

import (
	"github.com/over-engineering/msa/models"
)

func CreateMentals() models.Mentals {
	mentals := models.Mentals{
		models.Ambition:     &models.Mental{Type: models.Ambition, Value: 50},
		models.Boldness:     &models.Mental{Type: models.Boldness, Value: 50},
		models.Aggression:   &models.Mental{Type: models.Aggression, Value: 50},
		models.Predictation: &models.Mental{Type: models.Predictation, Value: 50},
		models.Composure:    &models.Mental{Type: models.Composure, Value: 50},
		// models.Concentration: &models.Mental{Type: models.Concentration, Value: 50},
		models.Immersion:   &models.Mental{Type: models.Immersion, Value: 50},
		models.Competition: &models.Mental{Type: models.Competition, Value: 50},
		models.SelfEsteem:  &models.Mental{Type: models.SelfEsteem, Value: 50},
		models.Confidence:  &models.Mental{Type: models.Confidence, Value: 50},
		models.Attention:   &models.Mental{Type: models.Attention, Value: 50},
	}
	return mentals
}

// func CreateMentalMap(l []float32) [16]map[interface{}]float32 {

// 	mentalMap := map[interface{}]float32{
// 		models.Ambition:      l[0],
// 		models.Boldness:      l[1],
// 		models.Aggression:    l[2],
// 		models.Predictation:  l[3],
// 		models.Composure:     l[4],
// 		models.Concentration: l[5],
// 		models.Immersion:     l[6],
// 		models.Competition:   l[7],
// 		models.SelfEsteem:    l[8],
// 		models.Confidence:    l[9],
// 		models.Attention:     l[10],
// 	}

// 	list := [16]map[interface{}]float32{}
// 	for i := 0; i < 16; i++ {
// 		list[i] = mentalMap
// 	}

// 	return list
// 	// var mentalMap models.PersonalityTable
// 	// var mentalMap map[models.PersonalityType]map[models.MentalType]float32
// 	// var mentalMaps = models.PersonalityTable{}

// 	// for pType := models.ISTJ; pType <= models.ENTJ; pType++ {
// 	// 	mentalMaps[pType] = mentalMap
// 	// 	// fmt.Println("Example_metntal2", mentalMap[pType])
// 	// 	// Output:=
// 	// }
// 	// // fmt.Println("Example_metntal22", mentalMap)
// 	// //Output:
// 	// fmt.Println(mentalMaps)
// 	// // Output

// 	// pTable := models.CreatePersonalityTable(mentalMaps)
// 	// fmt.Println(pTable)
// 	//Output:

// }

// func Example_mentals() {
// 	testSet := map[models.MentalType]map[string]float32{
// 		models.Ambition: map[string]float32{
// 			"Example1": 50,
// 			"Example2": 40,
// 			"Example3": 10,
// 		},
// 		models.Boldness: map[string]float32{
// 			"Example1": 10,
// 			"Example2": 2,
// 			"Example3": 3,
// 		},
// 		models.Aggression: map[string]float32{
// 			"Example1": 100,
// 			"Example2": 100,
// 			"Example3": 100,
// 		},
// 		models.Predictation: map[string]float32{
// 			"Example1": 100,
// 			"Example2": 100,
// 			"Example3": 100,
// 		},
// 		models.Composure: map[string]float32{
// 			"Example1": 100,
// 			"Example2": 100,
// 			"Example3": 100,
// 		},
// 		models.Concentration: map[string]float32{
// 			"Example1": 100,
// 			"Example2": 100,
// 			"Example3": 100,
// 		},
// 		models.Immersion: map[string]float32{
// 			"Example1": 100,
// 			"Example2": 100,
// 			"Example3": 100,
// 		},
// 		models.Competition: map[string]float32{
// 			"Example1": 100,
// 			"Example2": 100,
// 			"Example3": 100,
// 		},
// 		models.SelfEsteem: map[string]float32{
// 			"Example1": 100,
// 			"Example2": 100,
// 			"Example3": 100,
// 		},
// 		models.Confidence: map[string]float32{
// 			"Example1": 100,
// 			"Example2": 100,
// 			"Example3": 100,
// 		},
// 		models.Attention: map[string]float32{
// 			"Example1": 100,
// 			"Example2": 100,
// 			"Example3": 100,
// 		},
// 	}

// 	mentalMap, _ := models.NewMentalMap(testSet)

// 	fmt.Println(mentalMap)
// 	// Output:
// 	// map[Ambition:{Ambition 26} Boldness:{Boldness 3.3} Aggression:{Aggression 100} Predictation:{Predictation 100} Composure:{Composure 100} Concentration:{Concentration 100} Immersion:{Immersion 100} Competition:{Competition 100} SelfEsteem:{SelfEsteem 100} Confidence:{Confidence 100} Attention:{Attention 100}]
// }
