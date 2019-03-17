package models

import "fmt"

// IntelligenceType is an enum for various intelligence types accouding to
// Howard Gardner's multiple intelligence theory.
type IntelligenceType int

const (
	MusicalRhythmic = iota
	VisualSpatial
	VerbalLinguistic
	LogicalMathematical
	BodilyKinesthetic
	Interpersonal
	Intrapersonal
	Naturalistic
)

// Intelligence only have its value right now, but we may introduce some
// aspects like increase/decrease level.
type Intelligence struct {
	Type      IntelligenceType `json:"type"`
	Value     float32          `json:"value"`
	PeakAge   int              `json:"peakAge`
	Variation float32          `json:"variation"`
}

// 나이에 따라 감소하는 등 intelligence update rule 정의해야 함
func (i Intelligence) UpdateIntelligenceAsRule(age int) {
	if age <= i.PeakAge {
		i.Value += i.Variation
	} else {
		i.Value -= i.Variation
	}
}

func NewIntelligence(t IntelligenceType, value float32, peakAge int, variation float32) Intelligence {
	return Intelligence{Type: t, Value: value, PeakAge: peakAge, Variation: variation}
}

func NewIntelligenceMap(initialValues map[IntelligenceType]float32, peakAge int, variation float32) (map[IntelligenceType]Intelligence, error) {
	intelligenceMap := map[IntelligenceType]Intelligence{}
	for i := MusicalRhythmic; i < Naturalistic; i++ {
		value, ok := initialValues[IntelligenceType(i)]
		if !ok {
			return nil, fmt.Errorf("%v doesn't exist in values", IntelligenceType(i))
		}

		intelligenceMap[IntelligenceType(i)] = NewIntelligence(IntelligenceType(i), value, peakAge, variation)
	}

	return intelligenceMap, nil
}

// func UpdateIntelligenceMap(age int, intelligenceMap map[IntelligenceType]Intelligence) {
// 	for _, value := range intelligenceMap {
// 		ApplyIntelligenceRule(age, value)
// 	}
// }
