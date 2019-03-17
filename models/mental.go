package models

import (
	"fmt"
)

type MentalType int

const (
	Ambition MentalType = iota
	Boldness
	Aggression
	Predictation
	Composure // 참착성
	Concentration
	Immersion
	Competition
	SelfEsteem
	Confidence
	Attention // 관종도
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

type Mental struct {
	Type              MentalType                  `json:"type"`
	Value             float32                     `json:"value"` // max value is 100
	PersonalityImpact map[PersonalityType]float32 `json:"impact"`
}

// coefficient impact?
// func (m Mental) UpdateMentalValue(amount float32) {
// }

func (m Mental) UpdateMentalValue(p PersonalityType, amount float32) {
	m.Value += amount * m.PersonalityImpact[p]
}

// NewMental: make new mental
func NewMental(t MentalType, values map[string]float32, pImpact map[PersonalityType]float32) (Mental, error) {
	mental := Mental{Type: t, Value: 0, PersonalityImpact: pImpact}
	coefficients := mentalCoefficients[t]

	if coefficients == nil {
		return mental, fmt.Errorf("%v is not mental type", t)
	}

	if len(coefficients) != len(values) {
		return mental, fmt.Errorf("key number is different in mental type %v", t)
	}

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

func NewMentalMap(values map[MentalType]map[string]float32, pImpacts map[MentalType]map[PersonalityType]float32) (map[MentalType]Mental, error) {
	mentalMap := map[MentalType]Mental{}
	var err error
	for key, value := range values {
		if mentalMap[key], err = NewMental(key, value, pImpacts[key]); err != nil {
			return mentalMap, err
		}
	}

	return mentalMap, err
}

// func UpdateMentalMap(mentalMap map[MentalType]Mental, typeToUpdate MentalType, valueToAdd float32) error {
// 	m := mentalMap[typeToUpdate]
// 	if m.Type != typeToUpdate {
// 		return fmt.Errorf("Mental type %v doesn't exist in mental map", typeToUpdate)
// 	}

// 	if m.Value+valueToAdd < 0 || m.Value+valueToAdd > 100 {
// 		return fmt.Errorf("Update value < 0 || > 100")
// 	}

// 	m.Value += valueToAdd

// 	return nil
// }
