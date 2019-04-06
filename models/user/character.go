package user

import (
	"time"

	"github.com/over-engineering/msa/models/entity"
)

type User struct{}

// Character represents the game character belongs to user
// Things could be referencing with unique IDs.
type Character struct {
	User              `json:"user"`
	entity.Individual `json:"individual"`

	CurrentTime time.Time `json:"current_time"`
}

// NewCharacter returns new Character object
func NewCharacter() *Character {
	return &Character{}
}
