package models

// IntelligenceType is an enum for various intelligence types accouding to
// Howard Gardner's multiple intelligence theory.
type IntelligenceType Type

// Intelligence only have its value right now, but we may introduce some
// aspects like increase/decrease level.
type Intelligence struct {
	Value float32 `json:"value"`
}
