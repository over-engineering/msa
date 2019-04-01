package goods

import (
	"github.com/over-engineering/msa/models/character"
	"github.com/over-engineering/msa/models/types"
)

type GoodsManager interface {
	GetID() types.UID
	GetGoodsType() GoodsType
	GetName() string
	GetPrice() float32
	SetPrice(p float32)
	UpdateByDay()
	Use(by *character.Character, option int, args []string) error
}
