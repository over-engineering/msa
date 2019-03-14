package models

import "errors"

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

// Personality have its own value, it may not need struct structure.
type Personality struct {
	// Value float32 `json:"value"` // value range: 0 ~ 1
	Type         PersonalityType `json:"type"`
	Extraversion float32         `json:"extraversion"`
	Sensing      float32         `json:"sensing"`
	Thinking     float32         `json:"thinking"`
	Judging      float32         `json:"judging"`
}

// FindType find personality type
func (p Personality) FindType() (pType PersonalityType, err error) {
	e := p.Extraversion >= 0.5
	s := p.Sensing >= 0.5
	t := p.Thinking >= 0.5
	j := p.Judging >= 0.5

	// missing exp == true
	switch {
	case !e && s && t && j:
		return ISTJ, nil
	case !e && s && !t && j:
		return ISFJ, nil
	case !e && !s && !t && j:
		return INFJ, nil
	case !e && !s && t && j:
		return INTJ, nil
	case !e && s && t && !j:
		return ISTP, nil
	case !e && s && !t && !j:
		return ISFP, nil
	case !e && !s && !t && !j:
		return INFP, nil
	case !e && !s && t && !j:
		return INTP, nil
	case e && s && t && !j:
		return ESTP, nil
	case e && s && !t && !j:
		return ESFP, nil
	case e && !s && !t && !j:
		return ENFP, nil
	case e && !s && t && !j:
		return ENTP, nil
	case e && s && t && j:
		return ESTJ, nil
	case e && s && !t && j:
		return ESFJ, nil
	case e && !s && !t && j:
		return ENFJ, nil
	case e && !s && t && j:
		return ENTJ, nil
	default:
		return UNKNOWN, errors.New("Personality type error")
	}
}
