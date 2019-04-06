package goods

import (
	"github.com/over-engineering/msa/models/types"
)

// GlobalOwnerMap is a map uses owner's ID as a key, OwnedMap as a val.
type GlobalOwnerMap map[types.UID]OwnedMap

// OwnedMap is a map uses goods' ID as a key, OwnedGoods as a val.
type OwnedMap map[types.UID]OwnedGoods
