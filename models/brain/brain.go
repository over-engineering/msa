package brain

import (
	"github.com/over-engineering/msa/models/types"
)

type Brain struct {
	AthleticBrain SubBrain `json:"athletic_brain`
	MusicBrain    SubBrain `json:"music_brain`
	LanguageBrain SubBrain `json:"language_brain`
	MathBrain     SubBrain `json:"math_brain`
	VisualBrain   SubBrain `json:"visual_brain`
}

func NewBrain(val float32) Brain {
	return Brain{
		AthleticBrain: NewSubBrain(0, []float32{val, val, val, val, val}),
		MusicBrain:    NewSubBrain(0, []float32{val, val, val, val, val}),
		LanguageBrain: NewSubBrain(0, []float32{val, val, val, val, val}),
		MathBrain:     NewSubBrain(0, []float32{val, val, val, val, val}),
		VisualBrain:   NewSubBrain(0, []float32{val, val, val, val, val}),
	}
}

type BrainType types.Type

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
