package physical

import (
	"fmt"
	"math/rand"

	"github.com/over-engineering/msa/common/models/types"
)

// JointType is an enum for various joint types
type JointType types.Type

const (
	LeftShoulderJoint JointType = iota
	RightShoulderJoint
	LeftElbowJoint
	RightElbowJoint
	LeftHandJoint
	RightHadJoint
	LeftHipJoint
	RightHipJoint
	LeftKneeJoint
	RightKneeJoint
	LeftFootJoint
	RightFootJoint
)

var jointTypeStrings = [12]string{
	"LeftShoulderJoint",
	"RightShoulderJoint",
	"LeftElbowJoint",
	"RightElbowJoint",
	"LeftHandJoint",
	"RightHadJoint",
	"LeftHipJoint",
	"RightHipJoint",
	"LeftKneeJoint",
	"RightKneeJoint",
	"LeftFootJoint",
	"RightFootJoint",
}

func (jt JointType) String() string {
	return jointTypeStrings[jt]
}

// Joint represents the joint in character's body
type Joint struct {
	Type         JointType `json:"type"`
	Strength     float32   `json:"strength"`  // Strength is the parameter for calculation of wearing out
	Condition    float32   `json:"condition"` // Condition range: 0~1
	MaxCondition float32   `json:"max_condition"`
	Sprained     bool      `json:"sprained"`
}

func (j *Joint) String() string {
	return fmt.Sprint(*j)
}

// Joints is a data structure for Joint Map.
type Joints []Joint

// NewJoints returns Joints with intitialized value.
func NewJoints(r *rand.Rand, b, s float32) Joints {
	m := Joints{}
	for i := LeftShoulderJoint; i < RightFootJoint; i++ {
		newJoint := Joint{
			Type:         i,
			Strength:     (r.Float32()-0.5)*s + b,
			Condition:    100,
			MaxCondition: 100,
			Sprained:     false,
		}
		m = append(m, newJoint)
	}
	return m
}
