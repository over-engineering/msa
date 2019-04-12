package physical

import "github.com/over-engineering/msa/common/models/types"

// JointType is an enum for various joint types
type JointType types.Type

const (
	BALL_AND_SOCKET_JOINT JointType = iota
	ELBOW_JOINT
	GLIDING_JOINT
	HAND_JOINT
	HINGE_JOINT
	HIP_JOINT
	SADDLE_JOINT
	SPINE_JOINT
)

// Joint represents the joint in character's body
type Joint struct {
	Strength     float32 `json:"strength"`  // Strength is the parameter for calculation of wearing out
	Condition    float32 `json:"condition"` // Condition range: 0~1
	MaxCondition float32 `json:"max_condition"`
}
