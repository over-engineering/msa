package life

import (
	"errors"
	"time"
)

// LifeCycle represents the structure dealing with aging.
type LifeCycle struct {
	Alive      bool      `json:"alive"`
	Born       time.Time `json:"born"`
	Dead       time.Time `json:"lifeSpan"`
	Age        int       `json:"age"`
	Generation int       `json:"generation"`
}

// UpdateValues updates values in LifeCycle structure with given time.
func (l *LifeCycle) UpdateValues(currentTime time.Time) error {
	if currentTime.Before(l.Born) {
		return errors.New("current time is before born")
	}
	bornYear, _, _ := l.Born.Date()
	if currentTime.After(l.Dead) {
		l.Die()
		deadYear, _, _ := l.Dead.Date()
		l.Age = deadYear - bornYear
		return nil
	}
	currentYear, _, _ := currentTime.Date()
	l.Age = currentYear - bornYear
	return nil
}

// Die makes Alive to false
func (l *LifeCycle) Die() {
	l.Alive = false
}
