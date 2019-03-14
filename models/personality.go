package models

// PersonalityType consists of 16 different personalities.
// Each personality has its own strength, weakness points.
type PersonalityType Type

// Personality have its own value, it may not need struct structure.
type Personality struct {
	Value float32 `json:"value"` // value range: 0 ~ 1
}
