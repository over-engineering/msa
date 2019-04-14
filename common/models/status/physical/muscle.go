package physical

import (
	"fmt"
	"math/rand"

	"github.com/over-engineering/msa/common/models/types"
)

// MuscleType is an enum for various muscle types
type MuscleType types.Type

const (
	// Front Arm
	LeftExtensor MuscleType = iota
	RightExtensor
	LeftFlexor
	RightFlexor

	// Rear Arm
	LeftTriceps
	RightTriceps
	LeftBiceps
	RightBiceps
	LeftBrachialis
	RighttBrachialis
	LeftDeltoid
	RightDeltoid

	// Hand
	LeftLumbricalis
	RightLumbricalis
	// Flexor
	// Abductor
	// Opponens

	// Upper Leg
	LeftGluteals
	LeftQuadriceps
	LeftTFL
	LeftHamstring
	RightGluteals
	RightQuadriceps
	RightTFL
	RightHamstring

	// Lower Leg
	LeftCalfMuscle
	LeftAchillesTendon
	RightCalfMuscle
	RightAchillesTendon

	// Front Body
	Pectoralis
	Abdominals
	Oblique
	Serratus

	// Back Body
	Trapezius
	Latissimus
	Teres
	Infraspinatus
)

var muscleTypeString = [34]string{
	"LeftExtensor",
	"RightExtensor",
	"LeftFlexor",
	"RightFlexor",

	// Rear Arm
	"LeftTriceps",
	"RightTriceps",
	"LeftBiceps",
	"RightBiceps",
	"LeftBrachialis",
	"RighttBrachialis",
	"LeftDeltoid",
	"RightDeltoid",

	// Hand
	"LeftLumbricalis",
	"RightLumbricalis",
	// Flexor
	// Abductor
	// Opponens

	// Upper Leg
	"LeftGluteals",
	"LeftQuadriceps",
	"LeftTFL",
	"LeftHamstring",
	"RightGluteals",
	"RightQuadriceps",
	"RightTFL",
	"RightHamstring",

	// Lower Leg
	"LeftCalfMuscle",
	"LeftAchillesTendon",
	"RightCalfMuscle",
	"RightAchillesTendon",

	// Front Body
	"Pectoralis",
	"Abdominals",
	"Oblique",
	"Serratus",

	// Back Body
	"Trapezius",
	"Latissimus",
	"Teres",
	"Infraspinatus",
}

func (mt MuscleType) String() string {
	return muscleTypeString[mt]
}

// Muscle represents the muscles in character's body
type Muscle struct {
	Type         MuscleType `json:"type"`
	Mass         float32    `json:"mass"`      // Mass represents the mass of muscle
	RedRatio     float32    `json:"red_ratio"` // RedRatio represents ratio of red muscle in total muscle
	Condition    float32    `json:"condition"` // Condition range: 0~1
	MaxCondition float32    `json:"max_condition"`
	Ruptured     bool       `json:"ruptured"`
}

func (m *Muscle) String() string {
	return fmt.Sprint(*m)
}

type Muscles []Muscle

func NewMuscles(r *rand.Rand, b, s float32) Muscles {
	m := Muscles{}
	for i := LeftExtensor; i < Infraspinatus; i++ {
		newMuscle := Muscle{
			Type:         i,
			Mass:         (r.Float32()-0.5)*s + b,
			RedRatio:     0.5,
			Condition:    100,
			MaxCondition: 100,
			Ruptured:     false,
		}
		m = append(m, newMuscle)
	}
	return m
}
