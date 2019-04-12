package body

import (
	common "temp/structure/common"
)

type Bust struct {
	Head Head `json:"head"`
	Neck Neck `json:"Neck"`
}

type Head struct {
	common.Bone
	common.Fat

	Muscle1 common.Muscle
	Muscle2 common.Muscle

	Brain Brain `json:"brain"`

	LeftEye  Eye   `json:"left_eye"`
	RightEye Eye   `json:"right_eye"`
	Nose     Nose  `json:"nose"`
	Mouse    Mouse `json:"mouse"`

	Hair float32 `json:"hair"` // Hair represents the ratio of hair to area of head
}

type Neck struct {
	common.Bone
	common.Fat

	Muscle1 common.Muscle

	VocalCords float32 `json:"vocal_cords"`
}

func newBust() *Bust {
	bust := Bust{}
	return &bust
}

func newHead() *Head {
	head := Head{}
	return &head
}

func newNeck() *Neck {
	neck := Neck{}
	return &neck
}

func (h *Head) GetMuscles() map[string]common.Muscle {
}

/*
 * TODO: Make the element of brain affect the abilities
 */
type Brain struct {
	Frontal   float32 `json:"frontal"`
	Parietal  float32 `json:"parietal"`
	Occipital float32 `json:"occipital"`
	Temporal  float32 `json:"temporal"`
}

type Eye struct {
	Vision float32 `json:"vision"`
}

type Nose struct {
	Olfactory  float32 `json:"olfactory"`
	Congestion float32 `json:"congestion"` // Congestion represents the ratio which the nose hole is congested
}

/*
 * TODO
 * Calculate taste ability using taste cells
 * Calculate speech ability using muscle of tungue (with elements of brain)
 * Make teeth and tungue status affects the digest ability
 */
type Mouse struct {
}
