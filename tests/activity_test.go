package tests

import (
	"fmt"
	"testing"

	"github.com/over-engineering/msa/models"
)

// func Tmp2() {
func TestActivity(t *testing.T) {
	// func TestActivity(t *testing.T) {
	cList := CreateCharacters()
	c1 := cList[0]
	// c2 := cList[1]
	// c3 := cList[2]

	act := models.ActivityManager{
		Character: c1,
		JobManager: &models.SoccerPlayer{},
		ActPoint:  100,
		// Facility:
		// Job:
	}

	PrintCharacter(c1)

	apple := models.NewFood(models.AppleID, "Apple", 100, 50, 300)
	orange := models.NewFood(models.OrangeID, "Orange", 50, 60, 100)
	smartPhone := &models.SmartPhone{
		Goods: models.Goods{
			ID:        10,
			GoodsType: models.PhoneType,
			Name:      "Galaxy S10",
			Brand:     "Samsung",
			Price:     100,
			// Duration:    1000, // day
			Description: "Samsung Galaxy S10 is a line of Android smartphones manufactured by Samsung Electronics. The Galaxy S10 series is a celebratory series of the 10th anniversary of the Samsung Galaxy S flagship line. Unveiled during a press event, Samsung Galaxy Unpacked 2019 on February 20, 2019, they started shipping on March 8, 2019, and in some regions such as Australia and the United States, they started shipping them on March 6, 2019.",
		},
		Payment: 10,
		Memory:  8,    // GB
		Battery: 3400, // mAh
	}

	market := &models.Market{models.UID(15), models.Location{100, 100}, 100, []models.GoodsHelper{}, 0.2}

	fmt.Println("=============================== Market Test =======================")

	fmt.Println("================Market Goods==================")
	PrintGoods(market.GoodsList)

	// register goods
	market.AddGoods(apple)
	market.AddGoods(orange)
	market.AddGoods(smartPhone)
	fmt.Println("================Market Goods==================")
	PrintGoods(market.GoodsList)

	fmt.Println("================Character Goods==================")
	fmt.Println(act.Character.Goods)

	// buy goods
	act.MarketService(market, 2)
	fmt.Println("================Market Goods==================")
	PrintGoods(market.GoodsList)
	act.MarketService(market, 2)
	fmt.Println("================Market Goods==================")
	PrintGoods(market.GoodsList)
	fmt.Println("================Character Goods==================")
	fmt.Println(act.Character.Goods)

	// act.Eat(act.Character.Goods[], 1)
	// PrintCharacter(c1)

	// sell goods
	act.MarketService(market, 3)
	fmt.Println("================Character Goods==================")
	fmt.Println(act.Character.Goods)
	act.MarketService(market, 3)
	fmt.Println("================Character Goods==================")
	fmt.Println(act.Character.Goods)
	fmt.Println("================Market Goods==================")
	PrintGoods(market.GoodsList)

	fitness := &models.Fitness{10, models.Location{2, 5}, 100, 50}
	act.FitService(fitness, 1)
	PrintCharacter(c1)
}

func PrintGoods(g []models.GoodsHelper) {
	for _, val := range g {
		fmt.Println(val)
	}
}
