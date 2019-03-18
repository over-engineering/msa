package models

import "time"

// Condition TBD
type Condition struct {
	Type     string     `json:"type"`
	Duration *time.Time `json:"duration"`
}


