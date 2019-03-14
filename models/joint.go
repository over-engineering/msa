package models

// JointType is an enum for various joint types
type JointType Type

// Joint represents the joint in character's body
type Joint struct {
	Strength     int     `json:"strength"`  // Strength is the parameter for calculation of wearing out
	Condition    float32 `json:"condition"` // Condition range: 0~1
	MaxCondition float32 `json:"maxCondition"`
}
