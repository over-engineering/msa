package physical

import (
	"fmt"
	"math/rand"

	"github.com/over-engineering/msa/common/models/types"
)

// BoneType represents the type of bone
type BoneType types.Type

const (
	Skull BoneType = iota // 0
	JawBone
	NeckBone     // 목 (척추골의 상단이기도 함.)
	LeftClavicle // 쇄골
	RightClavicle
	LeftRib
	RightRib
	Vertabra // 척추골
	LeftUpperArmBone
	RightUpperArmBone
	LeftForeArmBone
	RightForeArmBone
	LeftFingerBone
	RightFingerBone
	Pelvis // 골반뼈
	LeftThighBone
	RightThighBone
	LeftCalfBone
	RightCalfBone
	LeftToeBone
	RightToeBone
)

var boneTypeStrings = [21]string{
	"Skull",
	"JawBone",
	"NeckBone",
	"Left Clavicle",
	"Right Clavicle",
	"Left Rib",
	"Right Rib",
	"Vertabra",
	"Left UpperArmBone",
	"Right UpperArmBone",
	"Left ForeArmBone",
	"Right ForeArmBone",
	"Left FingerBone",
	"Right FingerBone",
	"Pelvis",
	"Left ThighBone",
	"Right ThighBone",
	"Left CalfBone",
	"Right CalfBone",
	"Left ToeBone",
	"Right ToeBone",
}

func (bt BoneType) String() string {
	return boneTypeStrings[bt]
}

// Bone represents the strength, condition, cracked info
// about the bone.
type Bone struct {
	Type         BoneType `json:"type"`
	Strength     float32  `json:"strength"`  // Strength is the parameter for calculation of wearing out
	Condition    float32  `json:"condition"` // Condition range: 0~1
	MaxCondition float32  `json:"max_condition"`
	Cracked      bool     `json:"cracked"`
}

func (b *Bone) String() string {
	return fmt.Sprint(*b)
}

// Bones is a data structure for Bone Map.
type Bones []Bone

// NewBones returns Bones with intialized value.
func NewBones(r *rand.Rand, b, s float32) Bones {
	m := Bones{}
	for i := Skull; i < RightToeBone; i++ {
		newBone := Bone{
			Type:         i,
			Strength:     (r.Float32()-0.5)*s + b,
			Condition:    100,
			MaxCondition: 100,
			Cracked:      false,
		}
		m = append(m, newBone)
	}
	return m
}
