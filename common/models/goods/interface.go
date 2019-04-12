package goods

// UniqueInfos represents static, unique information of goods.
type UniqueInfos interface{}

// OwnedGoods represents owned goods by owner.
type OwnedGoods interface {
	UpdateByDay() error
}
