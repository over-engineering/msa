package models

import (
	"errors"
	"time"
)

// ConditionType represents all kinds of conditions.
type ConditionType struct {
	Name     string         `json:"type"`
	Duration *time.Duration `json:"duration"`
	Effects  Effects        `json:"effects"`
}

// MakeNewConditionType returns new ConditionType.
func MakeNewConditionType(name string, duration *time.Duration, effects Effects) *ConditionType {
	return &ConditionType{name, duration, effects}
}

// Describe returns the description of the condition.
func (c ConditionType) Describe() string {
	description := c.Name + "은(는) 보통 "
	description += c.Duration.String() + " 동안 나타납니다."
	return description
}

// Condition represents all kinds of conditions.
type Condition struct {
	ConditionType
	StartTime time.Time
	endTime   *time.Time
}

// MakeNewCondition returns new condition.
func MakeNewCondition(c ConditionType, startTime time.Time) *Condition {
	if c.Duration != nil {
		// 영구적이지 않다.
		endTime := startTime.Add(*c.Duration)
		return &Condition{
			ConditionType: c,
			StartTime:     startTime,
			endTime:       &endTime,
		}
	}
	return &Condition{
		ConditionType: c,
		StartTime:     startTime,
	}
}

// Shorten shortens endTime of condition.
func (c *Condition) Shorten(d time.Duration) error {
	if c.endTime != nil {
		*c.endTime = c.endTime.Add(-d)
		return nil
	}
	return errors.New("endTime is not specified")
}

// Extend extends endTime of condition.
func (c *Condition) Extend(d time.Duration) error {
	if c.endTime != nil {
		*c.endTime = c.endTime.Add(d)
		return nil
	}
	return errors.New("endTime is not specified")
}

// Conditions represents []Condition
type Conditions map[string]Condition

// DeleteByTypeName deletes value from map.
func (cs *Conditions) DeleteByTypeName(cName string) error {
	delete(*cs, cName)
	return nil
}

// DeleteByTime deletes conditions which are expired.
func (cs *Conditions) DeleteByTime(t time.Time) {
	for _, c := range *cs {
		if c.endTime.Before(t) {
			cs.DeleteByTypeName(c.Name)
		}
	}
}
