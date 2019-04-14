package user

import (
	"time"

	"github.com/over-engineering/msa/common/models/entity"
	"github.com/over-engineering/msa/common/models/types"
)

type User struct {
	ID types.UID `json:"id"`
}

// Character represents the game character belongs to user
// Things could be referencing with unique IDs.
type Character struct {
	User              `json:"user"`
	entity.Individual `json:"individual"`

	CurrentTime time.Time `json:"current_time"`
}

// NewCharacter returns new Character object
func NewCharacter(
	userId types.UID,
	individual entity.Individual,
	gameTime time.Time,
) *Character {
	return &Character{
		User:        User{ID: userId},
		Individual:  individual,
		CurrentTime: gameTime,
	}
}
