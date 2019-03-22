package models

// Condition TBD
type Condition struct {
	Type     string  `json:"type"`
	Duration int     `json:"duration"`
	Effects  Effects `json:"effects"`
}

<<<<<<< HEAD
func (c *Condition) SubDuration(time int) {
	c.Duration -= time
}

func (c *Condition) AddDuration(time int) {
	c.Duration += time
}
=======
type Conditions []Condition
>>>>>>> 29b16bf37fe6f6e2531ca625565009f4ab0c505f
