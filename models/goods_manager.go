package models

// Goods interface covers all kinds of goods
// type Goods interface {
type GoodsHelper interface {
	GetID() UID
	GetGoodsType() GoodsType
	GetName() string
	GetPrice() float32
	SetPrice(p float32)
}

func GetGoods(goodsList [][]GoodsHelper, goodsType GoodsType, goodsIndex int) GoodsHelper {
	return goodsList[goodsType][goodsIndex]
}

func RegisterGoods(goodsList [][]GoodsHelper, g GoodsHelper) {
	goodsList[g.GetGoodsType()] = append(goodsList[g.GetGoodsType()], g)
}

func UnRegisterGoods(goodsList [][]GoodsHelper, goodsType GoodsType, index int) {
	goodsList[goodsType] = append(goodsList[goodsType][:index], goodsList[goodsType][index+1:]...)
}

type FoodManager interface {
	Eat(st *Status)
}
