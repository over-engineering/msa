package body

import (
	common "temp/structure/common"
)

type Arm struct {
	bType common.BodyType `json:"b_type"`
	Front FrontArm        `json:"front"`
	Rear  RearArm         `json:"rear"`
	Hand  Hand            `json:"hand"`
}

type FrontArm struct {
	common.Bone
	common.Fat

	Extensor common.Muscle `json:"extensor"`
	Flexor   common.Muscle `json:"flexor"`
}

type RearArm struct {
	common.Bone
	common.Fat

	Triceps    common.Muscle `json:"triceps"`
	Biceps     common.Muscle `json:"biceps"`
	Brachialis common.Muscle `json:"brachialis"`
	Deltoid    common.Muscle `json:"deltoid"`
}

type Hand struct {
	common.Bone
	common.Fat

	Lumbricalis common.Muscle `json:"lumbricalis"`
	Flexor      common.Muscle `json:"flexor"`
	Abductor    common.Muscle `json:"abductor"`
	Opponens    common.Muscle `json:"opponens"`
}

func newArm() *Arm {
	arm := Arm{}
	return &arm
}

func newFrontArm() *FrontArm {
	frontArm := FrontArm{}
	return &frontArm
}

func newRearArn() *RearArm {
	rearArm := RearArm{}
	return &rearArm
}

func newHand() *Hand {
	hand := Hand{}
	return &hand
}

func (a *Arm) GetMuscles() map[string]common.Muscle {
}
