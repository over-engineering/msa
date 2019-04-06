package activity

import (
	"time"

	"github.com/over-engineering/msa/models/goods"
	"github.com/over-engineering/msa/models/types"
	"github.com/over-engineering/msa/models/user"
	"github.com/over-engineering/msa/models/world/space"
)

var gom goods.GlobalOwnerMap = goods.GlobalOwnerMap{}
var ggm goods.GlobalGoodsMap = goods.GlobalGoodsMap{}

func Visit(c *user.Character, to space.Location, vehicleId types.UID) error {
	current := c.Location
	distance := space.Distance(current, to)
	if vehicleId == "" {
		// On foot
		// Vehicle에 따라 헬스 깎이는게 다르게
	}
	c.Status.Physical.Health -= distance * 0.001
	c.CurrentTime = c.CurrentTime.Add(time.Duration(4 / distance))
	c.Location = to
	return nil
}
