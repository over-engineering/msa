package models

type Brain struct {
	AthleticBrain SubBrain `json:"athletic_brain`
	MusicBrain    SubBrain `json:"music_brain`
	LanguageBrain SubBrain `json:"language_brain`
	MathBrain     SubBrain `json:"math_brain`
	VisualBrain   SubBrain `json:"visual_brain`
}

type BrainType int

const (
	// 운동
	MuscleControl BrainType = iota
	Balance
	Posture
	BodyMovement
	BodyAwareness
	// 감각
	Sensation
	Smell
	Hearing
	Taste
	// 수리
	MathmaticalCalculation
	MathmaticalApprox
	// 언어
	Speech
	LanguageComprehension
	MeaningOfWord
	//
	Concentration
	Planning
	ProblemSolving
	// 기억
	ShortTermMemory
	LongTermMemory
	// 공간
	VisualSpatialAttention
	Sight
	// 인지
	MotionPerception
	SpeedPerception
	ColorPerception
	FacialRecognition
	// 성격, 감정
	// Personality
	// Emotion
)

func (mType BrainType) String() string {
	names := [25]string{
		"MuscleControl",
		"Balance",
		"Posture",
		"BodyMovement",
		"BodyAwareness",
		"Sensation",
		"Smell",
		"Hearing",
		"Taste",
		"MathmaticalCalculation",
		"MathmaticalApprox",
		"Speech",
		"LanguageComprehension",
		"MeaningOfWord",
		"Concentration",
		"Planning",
		"ProblemSolving",
		"ShortTermMemory",
		"LongTermMemory",
		"VisualSpatialAttention",
		"Sight",
		"MotionPerception",
		"SpeedPerception",
		"ColorPerception",
		"FacialRecognition",
	}

	if mType < MuscleControl || mType > FacialRecognition {
		return "Error"
	}

	return names[mType]
}

// peak까지는 iSlope만큼 증가, peak 이후부터는 Dslope만큼 brain 값이 감소
const peak, iSlope, dSlope = 30, 1.1, 0.9

type SubBrain map[BrainType]float32

func NewSubBrain(index int, val []float32) SubBrain {
	m := SubBrain{}

	var i, j BrainType
	switch index {
	case 0:
		i = MuscleControl
		j = BodyAwareness
	case 1:
		i = Sensation
		j = Taste
	case 2:
		i = Speech
		j = MeaningOfWord
	case 3:
		i = MathmaticalCalculation
		j = MathmaticalApprox
	case 4:
		i = VisualSpatialAttention
		j = Sight
	default:
	}

	for i <= j {
		m[i] = val[i-MuscleControl]
		i++
	}
	return m
}

func (b SubBrain) UpdateValue(vMap SubBrain, talent float32) {
	for key, val := range vMap {
		b[key] += val * (100 + talent) / 100
	}
}

func (b SubBrain) UpdateByAge(age int) {
	if age <= peak {
		for _, val := range b {
			val *= iSlope
		}
	} else {
		for _, val := range b {
			val *= dSlope
		}
	}
}

// type AthleticBrain map[BrainType]float32

// func NewAthleticBrain(val []float32) AthleticBrain {
// 	m := map[BrainType]float32{}
// 	for i := MuscleControl; i <= BodyAwareness; i++ {
// 		m[BrainType(i)] = val[i-MuscleControl]
// 	}
// 	return m
// }

// func (ab AthleticBrain) UpdateValue(vMap AthleticBrain, tal Talents) {
// 	for key, val := range vMap {
// 		ab[key] += val * (100 + tal[Athletic]) / 100
// 	}
// }

// func (ab AthleticBrain) UpdateByAge(age int) {
// 	if age <= peak {
// 		for _, val := range ab {
// 			val *= iSlope
// 		}
// 	} else {
// 		for _, val := range ab {
// 			val *= dSlope
// 		}
// 	}
// }

// type MusicBrain map[BrainType]float32

// func NewMusicBrain(val []float32) MusicBrain {
// 	m := map[BrainType]float32{}
// 	for i := Sensation; i <= Taste; i++ {
// 		m[BrainType(i)] = val[i-Sensation]
// 	}
// 	return m
// }

// func (mb MusicBrain) UpdateValue(vMap MusicBrain, tal Talents) {
// 	for key, val := range vMap {
// 		mb[key] += val * (100 + tal[Music]) / 100
// 	}
// }

// func (mb MusicBrain) UpdateByAge(age int) {
// 	if age <= peak {
// 		for _, val := range mb {
// 			val *= iSlope
// 		}
// 	} else {
// 		for _, val := range mb {
// 			val *= dSlope
// 		}
// 	}
// }

// type LanguageBrain map[BrainType]float32

// func NewLanguageBrain(val []float32) LanguageBrain {
// 	m := map[BrainType]float32{}
// 	for i := Speech; i <= MeaningOfWord; i++ {
// 		m[BrainType(i)] = val[i-Speech]
// 	}
// 	return m
// }

// func (lb LanguageBrain) UpdateValue(vMap LanguageBrain, tal Talents) {
// 	for key, val := range vMap {
// 		lb[key] += val * (100 + tal[Language]) / 100
// 	}
// }

// func (lb LanguageBrain) UpdateByAge(age int) {
// 	if age <= peak {
// 		for _, val := range lb {
// 			val *= iSlope
// 		}
// 	} else {
// 		for _, val := range lb {
// 			val *= dSlope
// 		}
// 	}
// }

// type MathBrain map[BrainType]float32

// func NewMathBrain(val []float32) MathBrain {
// 	m := map[BrainType]float32{}
// 	for i := MathmaticalCalculation; i <= MathmaticalApprox; i++ {
// 		m[BrainType(i)] = val[i-MathmaticalCalculation]
// 	}
// 	return m
// }

// func (mb MathBrain) UpdateValue(vMap MathBrain, tal Talents) {
// 	for key, val := range vMap {
// 		mb[key] += val * (100 + tal[Math]) / 100
// 	}
// }

// func (mb MathBrain) UpdateByAge(age int) {
// 	if age <= peak {
// 		for _, val := range mb {
// 			val *= iSlope
// 		}
// 	} else {
// 		for _, val := range mb {
// 			val *= dSlope
// 		}
// 	}
// }

// type VisualBrain map[BrainType]float32

// func NewVisualBrain(val []float32) VisualBrain {
// 	m := map[BrainType]float32{}
// 	for i := VisualSpatialAttention; i <= Sight; i++ {
// 		m[BrainType(i)] = val[i-VisualSpatialAttention]
// 	}
// 	return m
// }

// func (vb VisualBrain) UpdateValue(vMap VisualBrain, tal Talents) {
// 	for key, val := range vMap {
// 		vb[key] += val * (100 + tal[Visualization]) / 100
// 	}
// }

// func (vb VisualBrain) UpdateByAge(age int) {
// 	if age <= peak {
// 		for _, val := range vb {
// 			val *= iSlope
// 		}
// 	} else {
// 		for _, val := range vb {
// 			val *= dSlope
// 		}
// 	}
// }
