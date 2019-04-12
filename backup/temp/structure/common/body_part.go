package common

type BodyPart interface {
	GetMuscles() map[string]Muscle
	GetFat() int32
}
