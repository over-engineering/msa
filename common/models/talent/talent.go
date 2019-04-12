package talent

import (
	"math/rand"

	"github.com/over-engineering/msa/common/models/types"
)

type TalentType types.Type

const (
	Athletic TalentType = iota
	Music
	Language
	Math
	Visualization
)

var talentTypeString = [5]string{
	"Ahtletic",
	"Music",
	"Language",
	"Math",
	"Visualization",
}

func (mType TalentType) String() string {
	return talentTypeString[mType]
}

type Talents [5]float32

func NewTalents(r *rand.Rand, b, s float32) Talents {
	m := Talents{}
	for i, _ := range m {
		m[i] = (r.Float32()-0.5)*s + b
	}
	return m
}

func (t Talents) UpdateValue(val Talents) {
	for key, v := range val {
		t[key] += v
	}
}
