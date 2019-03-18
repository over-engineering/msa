package models

import "fmt"

type Age struct {
	Age      int `json:"age"`
	LifeSpan int `json:"lifeSpan"`
}

func (a *Age) UpdateValue(age int) error {
	if a.Age+age < 0 {
		return fmt.Errorf("age < 0")
	}

	a.Age += age

	if a.Age >= a.LifeSpan {
		a.Dead()
	}

	return nil
}

func (a Age) Dead() {

}
