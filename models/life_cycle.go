package models

import "fmt"

type LifeCycle struct {
	Age      float32 `json:"age"`
	LifeSpan float32 `json:"lifeSpan"`
}

func (l *LifeCycle) UpdateValue(age float32) error {
	if l.Age+age < 0 {
		return fmt.Errorf("age < 0")
	}

	l.Age += age

	if l.Age >= l.LifeSpan {
		l.Dead()
	}

	return nil
}

func (l LifeCycle) Dead() {

}
