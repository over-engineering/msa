package models

// MuscleType is an enum for various muscle types
type MuscleType Type

// Muscle represents the muscles in character's body
type Muscle struct {
	Mass     float32 `json:"mass"`     // Mass represents the mass of muscle
	RedRatio float32 `json:"redRatio"` // RedRatio represents ratio of red muscle in total muscle
}
