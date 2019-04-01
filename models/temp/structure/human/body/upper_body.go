package body

import (
	common "temp/structure/common"
)

type UpperBody struct {
	Front FrontBody `json:"front"`
	Back  BackBody  `json:"back"`

	/* Could be replaced with organ system
	Heart     Heart     `json:"heart"`
	Lung      Lung      `json:"lung"`
	Liver     Liver     `json:"liver"`
	Intestine Intestine `json:"intestine"`
	*/
}

type FrontBody struct {
	common.Bone // The lib
	common.Fat

	Pectoralis common.Muscle `json:"pectoralis"`
	Abdominals common.Muscle `json:"abdominals"`
	Oblique    common.Muscle `json:"oblique"`
	Serratus   common.Muscle `json:"serratus"`
}

type BackBody struct {
	common.Bone // The backbone
	common.Fat

	Trapezius     common.Muscle `json:"trapezisu"`
	Latissimus    common.Muscle `json:"latissimus"`
	Teres         common.Muscle `json:"teres"`
	Infraspinatus common.Muscle `json:"infraspinatus"`
}

func newUpperBody() *UpperBody {
	upperBody := UpperBody{}
	return &upperBody
}

func newFrontBody() *FrontBody {
	frontBody := FrontBody{}
	return &frontBody
}

func newBackBody() *BackBody {
	backBody := BackBody{}
	return &backBody
}

func (b *UpperBody) GetMuscles() map[string]common.Muscle {
}

/*
 * TODO
 * What kind of elements wiil be needed?
 */
type Heart struct {
	Rate float32 `json:"rate"` // Rate represents the heart rate
}

type Lung struct {
	health float32 `json:"health"`
}

type Liver struct {
	health float32 `json:"health"`
}

type Intestine struct {
	health float32 `json:"health"`
}
