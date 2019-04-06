package physical

import "github.com/over-engineering/msa/models/types"

type BoneType types.Type

const (
	NeckBone BoneType = iota //0
)

type Bone struct {
	Strength     float32 `json:"strength"`  // Strength is the parameter for calculation of wearing out
	Condition    float32 `json:"condition"` // Condition range: 0~1
	MaxCondition float32 `json:"max_condition"`
	Cracked      bool    `json:"cracked"`
}
