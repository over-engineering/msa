package ability

import "github.com/over-engineering/msa/models/types"

// CommonAbility represents common abilities that could be used through
// all kinds of job.
type CommonAbility struct {
	// ID represents character or entity that have this ability
	ID types.UID `json:"id"`
}
