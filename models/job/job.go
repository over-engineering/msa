package job

import "github.com/over-engineering/msa/models/types"

// Type is an enum for various job types.
type Type types.Type

const (
	None           Type = iota // 0
	FootballPlayer             // 1
)

func (t Type) String() string {
	switch t {
	case FootballPlayer:
		return "Football Player"
	default:
		return "None"
	}
}
