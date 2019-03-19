package models

import (
	"fmt"
)

// PersonalityType consists of 16 different personalities.
// Each personality has its own strength, weakness points.
type PersonalityType Type

const (
	UNKNOWN PersonalityType = iota
	ISTJ
	ISFJ
	INFJ
	INTJ
	ISTP
	ISFP
	INFP
	INTP
	ESTP
	ESFP
	ENFP
	ENTP
	ESTJ
	ESFJ
	ENFJ
	ENTJ
)

const PersonalityNum = 16

func (pType PersonalityType) String() string {
	names := [17]string{
		"UNKNOWN",
		"ISTJ",
		"ISFJ",
		"INFJ",
		"INTJ",
		"ISTP",
		"ISFP",
		"INFP",
		"INTP",
		"ESTP",
		"ESFP",
		"ENFP",
		"ENTP",
		"ESTJ",
		"ESFJ",
		"ENFJ",
		"ENTJ",
	}

	if pType < ISTJ || pType > ENTJ {
		return names[0]
	}

	return names[pType]
}

type PersonalityTable map[PersonalityType]map[interface{}]float32

func (pTable PersonalityTable) AddPersonalityCoe(m [16]map[interface{}]float32) error {
	for i, value := range m {
		if value == nil {
			return fmt.Errorf("Pmap value == nil")
		}

		pTable[PersonalityType(i)] = value
	}

	return nil
}

var PTable = PersonalityTable{}

// Personality have its own value, it may not need struct structure.
type Personality struct {
	// Value float32 `json:"value"` // value range: 0 ~ 1
	Type         PersonalityType `json:"type"`
	Extraversion float32         `json:"extraversion"`
	Sensing      float32         `json:"sensing"`
	Thinking     float32         `json:"thinking"`
	Judging      float32         `json:"judging"`
}

func NewPersonality(values [4]float32) Personality {
	p := Personality{Extraversion: values[0], Sensing: values[1], Thinking: values[2], Judging: values[3]}
	p.FindType()
	return p
}

func (p *Personality) UpdateValue(ps Personality) {
	val := [4]float32{p.Extraversion + ps.Extraversion, p.Sensing + ps.Sensing, p.Thinking + ps.Thinking, p.Judging + ps.Judging}
	pVal := [4]*float32{&p.Extraversion, &p.Sensing, &p.Thinking, &p.Judging}
	for i := 0; i < 4; i++ {
		if val[i] < 0 {
			*pVal[i] = 0
		} else if val[i] > 100 {
			*pVal[i] = 100
		} else {
			*pVal[i] = val[i]
		}
	}
	// if val1 < 0 || val1 > 100 || val2 < 0 || val2 > 100 || val3 < 0 || val3 > 100 || val4 < 0 || val4 > 100 {
	// 	return fmt.Errorf("personalitly < 0 || > 100")
	// }

	// p.Extraversion = val1
	// p.Sensing = val2
	// p.Thinking = val3
	// p.Judging = val4
	p.FindType()
}

// FindType find personality type
func (p *Personality) FindType() error {
	if p.Extraversion < 0 || p.Extraversion > 1 ||
		p.Sensing < 0 || p.Sensing > 1 ||
		p.Thinking < 0 || p.Thinking > 1 ||
		p.Judging < 0 || p.Judging > 1 {
		return fmt.Errorf("Personality value should be 0~1")
	}

	e := p.Extraversion >= 0.5
	s := p.Sensing >= 0.5
	t := p.Thinking >= 0.5
	j := p.Judging >= 0.5

	// missing exp == true
	switch {
	case !e && s && t && j:
		p.Type = ISTJ
	case !e && s && !t && j:
		p.Type = ISFJ
	case !e && !s && !t && j:
		p.Type = INFJ
	case !e && !s && t && j:
		p.Type = INTJ
	case !e && s && t && !j:
		p.Type = ISTP
	case !e && s && !t && !j:
		p.Type = ISFP
	case !e && !s && !t && !j:
		p.Type = INFP
	case !e && !s && t && !j:
		p.Type = INTP
	case e && s && t && !j:
		p.Type = ESTP
	case e && s && !t && !j:
		p.Type = ESFP
	case e && !s && !t && !j:
		p.Type = ENFP
	case e && !s && t && !j:
		p.Type = ENTP
	case e && s && t && j:
		p.Type = ESTJ
	case e && s && !t && j:
		p.Type = ESFJ
	case e && !s && !t && j:
		p.Type = ENFJ
	case e && !s && t && j:
		p.Type = ENTJ
	default:
		return fmt.Errorf("Personality type error")
	}
	return nil
}
