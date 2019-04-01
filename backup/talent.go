package models

type TalentType int

const (
	Athletic TalentType = iota
	Music
	Language
	Math
	Visualization
)

func (mType TalentType) String() string {
	names := [5]string{
		"Ahtletic",
		"Music",
		"Language",
		"Math",
		"Visualization",
	}

	if mType < Athletic || mType > Visualization {
		return "Error"
	}

	return names[mType]
}

type Talents []float32

func NewTalents(val []float32) Talents {
	m := []float32{}
	for _, v := range val {
		m = append(m, v)
	}
	return m
}

func (t Talents) UpdateValue(val Talents) {
	for key, v := range val {
		t[key] += v
	}
}
