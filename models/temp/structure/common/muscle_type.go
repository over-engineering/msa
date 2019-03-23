package common

type MuscleType int

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
