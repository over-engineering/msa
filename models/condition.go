package models

// Condition TBD
type Condition struct {
	Type     string  `json:"type"`
	Duration int     `json:"duration"`
	Effects  Effects `json:"effects"`
}

func (c *Condition) SubDuration(time int) {
	c.Duration -= time
}

func (c *Condition) AddDuration(time int) {
	c.Duration += time
}
