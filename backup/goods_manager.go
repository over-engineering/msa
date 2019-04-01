package models

// Goods interface covers all kinds of goods
// type Goods interface {
type GoodsHelper interface {
	GetID() UID
	GetGoodsType() GoodsType
	GetName() string
	GetPrice() float32
	SetPrice(p float32)
	UpdateByDay()
}

type GoodsList [][]GoodsHelper

func (gl GoodsList) UpdateByDay() {
	for _, gs := range gl {
		for _, g := range gs {
			g.UpdateByDay()
		}
	}
}

func (gl GoodsList) GetGoods(goodsType GoodsType, goodsIndex int) GoodsHelper {
	return gl[goodsType][goodsIndex]
}

func (gl GoodsList) RegisterGoods(g GoodsHelper) {
	gl[g.GetGoodsType()] = append(gl[g.GetGoodsType()], g)
}

func (gl GoodsList) UnRegisterGoods(goodsType GoodsType, index int) {
	gl[goodsType] = append(gl[goodsType][:index], gl[goodsType][index+1:]...)
}

type FoodManager interface {
	Eat(st *Status)
}
