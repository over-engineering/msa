package history

import (
	"time"

	"github.com/over-engineering/msa/common/models/types"
	"github.com/over-engineering/msa/common/models/world/space"
)

var GenesisEntityMap map[types.UID]struct{}

var GlobalEntityChain []EntityBlock

// EntityBlock represents the 살생부
type EntityBlock struct {
	Timeline time.Time   `json:"timeline"`
	Births   []types.UID `json:"births"`
	Deaths   []types.UID `json:"deaths"`
}

var GenesisGoodsMap map[types.CID]struct{}

var GlobalGoodsChain []GoodsBlock

type GoodsBlock struct {
	Timeline time.Time   `json:"timeline"`
	Births   []types.CID `json:"births"`
	Deaths   []types.CID `json:"deaths"`
}

// Target represents the target map, which uses routing-like key value.
type Target map[string]struct {
	// Add represent boolean value if this is adding value or
	// subtracting value
	Add   bool   `json:"add"`
	Value string `json:"value"`
}

// Transition represents all transitions for the specific ID.
type Transition struct {
	ID       types.UID `json:"id"`
	ReasonID types.UID `json:"reason_id"` // Causing event id.
	Targets  Target    `json:"target"`
}

// TransitionBlock represents the data structure consists of timeline and
// corresponding transitions.
type TransitionBlock struct {
	Timeline    time.Time    `json:"timeline"`
	Transitions []Transition `json:"transitions"`
}

var GenesisFacilityMap map[types.UID]*space.Facility

var GlobalFacilityChain []TransitionBlock

var GlobalIndividualChain []TransitionBlock

var GlobalCorporationChain []TransitionBlock

// National, Continental Chain...
