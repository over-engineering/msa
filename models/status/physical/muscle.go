package physical

import "github.com/over-engineering/msa/models/types"

// MuscleType is an enum for various muscle types
type MuscleType types.Type

const (
	// Front Arm
	Extensor MuscleType = iota
	Flexor

	// Rear Arm
	Triceps
	Biceps
	Brachialis
	Deltoid

	// Hand
	Lumbricalis
	// Flexor
	Abductor
	Opponens

	// Upper Leg
	Gluteals
	Quadriceps
	TFL
	Hamstring

	// Lower Leg
	CalfMuscle
	AchillesTendon

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

// Muscle represents the muscles in character's body
type Muscle struct {
	Mass     float32 `json:"mass"`      // Mass represents the mass of muscle
	RedRatio float32 `json:"red_ratio"` // RedRatio represents ratio of red muscle in total muscle
	Fatigue  float32 `json:"fatigue"`
}
