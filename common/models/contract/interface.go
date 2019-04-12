package contract

import (
	"github.com/over-engineering/msa/common/models/finance"
	"github.com/over-engineering/msa/common/models/types"
)

// Executant defines implementation functions with contracts.
type Executant interface {
	Implement(b finance.Banker) error
	ImposeOn(vid types.UID, b finance.Banker) error
	GiveBonus(bid types.UID, b finance.Banker) error
}
