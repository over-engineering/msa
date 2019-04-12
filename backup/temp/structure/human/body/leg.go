package body

import (
	common "temp/structure/common"
)

type Leg struct {
	Upper UpperLeg `json:"upper"`
	Lower LowerLeg `json:"lower"`
	Foot  Foot     `json:"foot"`
}

type UpperLeg struct {
	common.Bone
	common.Fat

	Gluteals   common.Muscle `json:"gluteals"`
	Quadriceps common.Muscle `json:"quadriceps"`
	TFL        common.Muscle `json:"tfl"`
	Hamstring  common.Muscle `json:"hastring"`
}

type LowerLeg struct {
	common.Bone
	common.Fat

	CalfMuscle     common.Muscle `json:"calf_muscle"`
	AchillesTendon common.Muscle `json:"achille_tendon"`
}

type Foot struct {
	common.Bon
	common.Fat

	Muscle1 common.Muscle `json:"muscle_1"`
	Muscle2 common.Muscle `json:"muscle_2"`
}

func newLeg() *Leg {
	leg := Leg{}
	return &leg
}

func newUpperLeg() *UpperLeg {
	upperLeg := UpperLeg{}
	return &upperLeg
}

func newLowerLeg() *LowerLeg {
	lowerLeg := LowerLeg{}
	return &lowerLeg
}

func newFoot() *Foot {
	foot := Foot{}
	return &foot
}

func (l *Leg) GetMuscles() map[string]common.Muscle {
}
