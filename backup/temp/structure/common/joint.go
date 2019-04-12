package common

type Joint struct {
	Strength     int     `json:"strength"`
	Condition    float32 `json:"condition"`
	MaxCondition float32 `json:"max_condition"`
}
