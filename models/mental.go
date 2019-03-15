package models

// package main

import (
	"fmt"
)

type MentalType int

const (
	Ambition MentalType = iota
	Boldness
	Aggression
	Predictation
	Composure
	Concentration
	Immersion
	Competition
	SelfEsteem
	Confidence
	Attention
)

func (mType MentalType) String() string {
	names := [11]string{
		"Ambition",
		"Boldness",
		"Aggression",
		"Predictation",
		"Composure",
		"Concentration",
		"Immersion",
		"Competition",
		"SelfEsteem",
		"Confidence",
		"Attention",
	}

	if mType < Ambition || mType > Attention {
		return names[0]
	}

	return names[mType]
}

// define coefficient of each mental as slice(unexported)
var mentalCoefficients = map[MentalType]map[string]float32{
	Ambition: map[string]float32{
		"Example1": 0.1,
		"Example2": 0.4,
		"Example3": 0.5,
	},
	Boldness: map[string]float32{
		"Example1": 0.1,
		"Example2": 0.4,
		"Example3": 0.5,
	},
	Aggression: map[string]float32{
		"Example1": 0.1,
		"Example2": 0.4,
		"Example3": 0.5,
	},
	Predictation: map[string]float32{
		"Example1": 0.1,
		"Example2": 0.4,
		"Example3": 0.5,
	},
	Composure: map[string]float32{
		"Example1": 0.1,
		"Example2": 0.4,
		"Example3": 0.5,
	},
	Concentration: map[string]float32{
		"Example1": 0.1,
		"Example2": 0.4,
		"Example3": 0.5,
	},
	Immersion: map[string]float32{
		"Example1": 0.1,
		"Example2": 0.4,
		"Example3": 0.5,
	},
	Competition: map[string]float32{
		"Example1": 0.1,
		"Example2": 0.4,
		"Example3": 0.5,
	},
	SelfEsteem: map[string]float32{
		"Example1": 0.1,
		"Example2": 0.4,
		"Example3": 0.5,
	},
	Confidence: map[string]float32{
		"Example1": 0.1,
		"Example2": 0.4,
		"Example3": 0.5,
	},
	Attention: map[string]float32{
		"Example1": 0.1,
		"Example2": 0.4,
		"Example3": 0.5,
	},
}

// var (
// 	ambitionCoe = map[string]float32{
// 		"Example1": 0.1,
// 		"Example2": 0.4,
// 		"Example3": 0.5,
// 	}
// 	boldnessCoe = map[string]float32{
// 		"Example1": 0.1,
// 		"Example2": 0.4,
// 		"Example3": 0.5,
// 	}
// 	aggressionCoe = map[string]float32{
// 		"Example1": 0.1,
// 		"Example2": 0.4,
// 		"Example3": 0.5,
// 	}
// 	predictationCoe = map[string]float32{
// 		"Example1": 0.1,
// 		"Example2": 0.4,
// 		"Example3": 0.5,
// 	}
// 	composureCoe = map[string]float32{
// 		"Example1": 0.1,
// 		"Example2": 0.4,
// 		"Example3": 0.5,
// 	}
// 	concentrationCoe = map[string]float32{
// 		"Example1": 0.1,
// 		"Example2": 0.4,
// 		"Example3": 0.5,
// 	}
// 	immersionCoe = map[string]float32{
// 		"Example1": 0.1,
// 		"Example2": 0.4,
// 		"Example3": 0.5,
// 	}
// 	competitionCoe = map[string]float32{
// 		"Example1": 0.1,
// 		"Example2": 0.4,
// 		"Example3": 0.5,
// 	}
// 	selfEsteemCoe = map[string]float32{
// 		"Example1": 0.1,
// 		"Example2": 0.4,
// 		"Example3": 0.5,
// 	}
// 	confidenceCoe = map[string]float32{
// 		"Example1": 0.1,
// 		"Example2": 0.4,
// 		"Example3": 0.5,
// 	}
// 	attentionCoe = map[string]float32{
// 		"Example1": 0.1,
// 		"Example2": 0.4,
// 		"Example3": 0.5,
// 	}
// )

// MentalCoefficient: unexported struct
// type mentalCoefficient struct {
// 	ambitionCoe      []float32
// 	boldnessCoe      []float32
// 	aggressionCoe    []float32
// 	predictationCoe  []float32
// 	composureCoe     []float32 // 참착성
// 	concentrationCoe []float32
// 	immersionCoe     []float32
// 	competitionCoe   []float32
// 	selfEsteemCoe    []float32
// 	confidenceCoe    []float32
// 	attentionCoe     []float32 // 관종도
// }

type Mental struct {
	Type  MentalType `json:"type"`
	Value float32    `json:"value"` // max value is 100
	// Coefficients []float32  `json:"coefficient"`
	// Ambition      float32 `json:"ambition"`
	// Boldness      float32 `json:"boldness"`
	// Aggression    float32 `json:"aggression"`
	// Predictation  float32 `json:"predictation"`
	// Composure     float32 `json:"composure"` // 참착성
	// Concentration float32 `json:"concentration"`
	// Immersion     float32 `json:"immersion"`
	// Competition   float32 `json:"competition"`
	// SelfEsteem    float32 `json:"selfEsteem"`
	// Confidence    float32 `json:"confidence"`
	// Attention     float32 `json:"attention"` // 관종도
}

// NewMental: make new mental
func NewMental(t MentalType, values map[string]float32) (Mental, error) {

	coefficients := mentalCoefficients[t]
	mental := Mental{Type: t, Value: 0}

	if coefficients == nil {
		return mental, fmt.Errorf("%v is not mental type", t)
	}

	if len(coefficients) != len(values) {
		return mental, fmt.Errorf("key number is different in mental type %v", t)
	}
	// switch t {
	// case Ambition:
	// 	coefficients = ambitionCoe
	// case Boldness:
	// 	coefficients = boldnessCoe
	// case Aggression:
	// 	coefficients = aggressionCoe
	// case Predictation:
	// 	coefficients = predictationCoe
	// case Composure:
	// 	coefficients = composureCoe
	// case Concentration:
	// 	coefficients = concentrationCoe
	// case Immersion:
	// 	coefficients = immersionCoe
	// case Competition:
	// 	coefficients = competitionCoe
	// case SelfEsteem:
	// 	coefficients = selfEsteemCoe
	// case Confidence:
	// 	coefficients = confidenceCoe
	// case Attention:
	// 	coefficients = attentionCoe
	// default:
	// }

	var totalCoe float32
	for key, value := range values {
		if value < 0 || value > 100 {
			return mental, fmt.Errorf("%v < 0 or %v > 100 in mental type of %v", value, value, t)
		}

		if _, ok := coefficients[key]; !ok {
			return mental, fmt.Errorf("%v doesn't exist in mental type of %v", key, t)
		}

		mental.Value += coefficients[key] * value
		totalCoe += coefficients[key]
	}

	if totalCoe != 1 {
		return mental, fmt.Errorf("total coe != 1 in mental type of %v", t)
	}

	return mental, nil
}

func NewMentalMap(values map[MentalType]map[string]float32) (map[MentalType]Mental, error) {
	mentalMap := map[MentalType]Mental{}
	var err error
	for key, value := range values {
		if mentalMap[key], err = NewMental(key, value); err != nil {
			return mentalMap, err
		}
	}

	return mentalMap, err
}

func UpdateMentalMap(mentalMap map[MentalType]Mental, typeToUpdate MentalType, valueToAdd float32) error {
	m := mentalMap[typeToUpdate]
	if m.Type != typeToUpdate {
		return fmt.Errorf("Mental type %v doesn't exist in mental map", typeToUpdate)
	}

	if m.Value+valueToAdd < 0 || m.Value+valueToAdd > 100 {
		return fmt.Errorf("Update value < 0 || > 100")
	}

	m.Value += valueToAdd

	return nil
}

func test() {
	testSet := map[MentalType]map[string]float32{
		Ambition: map[string]float32{
			"Example1": 0,
			"Example2": 40,
			"Example3": 10,
		},
		Boldness: map[string]float32{
			"Example1": 10,
			"Example2": 2,
			"Example3": 3,
		},
		Aggression: map[string]float32{
			"Example1": 100,
			"Example2": 100,
			"Example3": 100,
		},
		Predictation: map[string]float32{
			"Example1": 100,
			"Example2": 100,
			"Example3": 100,
		},
		Composure: map[string]float32{
			"Example1": 100,
			"Example2": 100,
			"Example3": 100,
		},
		Concentration: map[string]float32{
			"Example1": 100,
			"Example2": 100,
			"Example3": 100,
		},
		Immersion: map[string]float32{
			"Example1": 100,
			"Example2": 100,
			"Example3": 100,
		},
		Competition: map[string]float32{
			"Example1": 100,
			"Example2": 100,
			"Example3": 100,
		},
		SelfEsteem: map[string]float32{
			"Example1": 100,
			"Example2": 100,
			"Example3": 100,
		},
		Confidence: map[string]float32{
			"Example1": 100,
			"Example2": 100,
			"Example3": 100,
		},
		Attention: map[string]float32{
			"Example1": 100,
			"Example2": 100,
			"Example3": 100,
		},
	}

	mentalMap, err := NewMentalMap(testSet)

	// fmt.Printf(mentalMap, err)
	fmt.Println(mentalMap, err)

}

// func main() {
// 	test()
// }
