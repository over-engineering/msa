package models

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

const IntelligenceNum = 8

func (iType IntelligenceType) String() string {
	names := [IntelligenceNum]string{
		"MusicalRhythmic",
		"VisualSpatial",
		"VerbalLinguistic",
		"LogicalMathematical",
		"BodilyKinesthetic",
		"Interpersonal",
		"Intrapersonal",
		"Naturalistic",
	}

	if iType < MusicalRhythmic || iType > Naturalistic {
		return names[0]
	}

	return names[iType]
}

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

type Intelligences map[IntelligenceType]*Intelligence

func (is Intelligences) UpdateValue(m map[IntelligenceType]float32) {
	for key, value := range m {
		val := is[key].Value + value
		if val < 0 {
			is[key].Value = 0
		} else if val > 100 {
			is[key].Value = 100
		} else {
			is[key].Value = val
		}

		// is[key].Value += value
	}
}

func NewIntelligence(t IntelligenceType, value float32, peakAge int, variation float32) *Intelligence {
	return &Intelligence{Type: t, Value: value, PeakAge: peakAge, Variation: variation}
}

func NewIntelligences(initialValues [IntelligenceNum]float32, peakAge [IntelligenceNum]int, variation [IntelligenceNum]float32) Intelligences {
	intelligences := Intelligences{}
	for i := MusicalRhythmic; i <= Naturalistic; i++ {
		// value, ok := initialValues[IntelligenceType(i)]
		// if !ok {
		// 	return nil, fmt.Errorf("%v doesn't exist in values", IntelligenceType(i))
		// }

		intelligences[IntelligenceType(i)] = NewIntelligence(IntelligenceType(i), initialValues[i], peakAge[i], variation[i])
	}

	return intelligences
}

// func UpdateIntelligenceMap(age int, intelligenceMap map[IntelligenceType]Intelligence) {
// 	for _, value := range intelligenceMap {
// 		ApplyIntelligenceRule(age, value)
// 	}
// }
