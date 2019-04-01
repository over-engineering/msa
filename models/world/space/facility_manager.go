package world

import (
	"github.com/over-engineering/msa/models/character"
	"github.com/over-engineering/msa/models/types"
)

type FacilityManager interface {
	GetID() types.UID
	GetLocation() Location

	GetInfo(option int) string
	UseFacility(by *character.Character, option int, args []string) error
	// Visit(l *Location, st *Status)
}
