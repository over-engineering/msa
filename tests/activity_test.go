package tests

// import (
// 	"fmt"
// 	"testing"

// 	"github.com/over-engineering/msa/models"
// )

// // func Tmp2() {
// func TestActivity(t *testing.T) {
// 	// func TestActivity(t *testing.T) {
// 	cList := CreateCharacters()
// 	c1 := cList[0]
// 	// c2 := cList[1]
// 	// c3 := cList[2]

// 	PrintCharacter(c1)

// 	apple := models.NewFood(models.AppleID, "Apple", 100, 50, 300, 7)
// 	orange := models.NewFood(models.OrangeID, "Orange", 50, 60, 100, 7)
// 	smartPhone := &models.SmartPhone{
// 		Goods: models.Goods{
// 			ID:        10,
// 			GoodsType: models.PhoneType,
// 			Name:      "Galaxy S10",
// 			Brand:     "Samsung",
// 			Price:     100,
// 			// Duration:    1000, // day
// 			Description: "Samsung Galaxy S10 is a line of Android smartphones manufactured by Samsung Electronics. The Galaxy S10 series is a celebratory series of the 10th anniversary of the Samsung Galaxy S flagship line. Unveiled during a press event, Samsung Galaxy Unpacked 2019 on February 20, 2019, they started shipping on March 8, 2019, and in some regions such as Australia and the United States, they started shipping them on March 6, 2019.",
// 		},
// 		Payment: 10,
// 		Memory:  8,    // GB
// 		Battery: 3400, // mAh
// 	}

// 	drug := CreateDrug()

// 	market := &models.Market{models.UID(15), &models.Location{100, 100}, 100, []models.GoodsHelper{}, 0.2}

// 	fmt.Println("=============================== Market Test =======================")

// 	fmt.Println("================Market Goods==================")
// 	PrintGoods(market.GoodsList)

// 	// register goods
// 	market.AddGoods(apple)
// 	market.AddGoods(orange)
// 	market.AddGoods(smartPhone)
// 	market.AddGoods(drug)
// 	fmt.Println("================Market Goods==================")
// 	PrintGoods(market.GoodsList)

// 	fmt.Println("================Character Goods==================")
// 	fmt.Println(c1.Goods)

// 	// buy goods
// 	// act.MarketService(market, 2)
// 	models.ActVisit(c1, market, []int{0, 0})
// 	fmt.Println("================Market Goods==================")
// 	PrintGoods(market.GoodsList)
// 	models.ActVisit(c1, market, []int{0, 0})
// 	fmt.Println("================Market Goods==================")
// 	PrintGoods(market.GoodsList)
// 	models.ActVisit(c1, market, []int{0, 0})
// 	fmt.Println("================Market Goods==================")
// 	PrintGoods(market.GoodsList)
// 	models.ActVisit(c1, market, []int{0, 0})
// 	fmt.Println("================Market Goods==================")
// 	PrintGoods(market.GoodsList)
// 	fmt.Println("================Character Goods==================")
// 	fmt.Println(c1.Goods)

// 	// act.Eat(act.Character.Goods[], 1)
// 	// PrintCharacter(c1)

// 	// Sell goods
// 	models.ActVisit(c1, market, []int{1, 3, 0})
// 	fmt.Println("================Character Goods==================")
// 	fmt.Println(c1.Goods)
// 	models.ActVisit(c1, market, []int{1, 3, 0})
// 	fmt.Println("================Character Goods==================")
// 	fmt.Println(c1.Goods)
// 	fmt.Println("================Market Goods==================")
// 	PrintGoods(market.GoodsList)

// 	fmt.Println(c1.Status)
// 	// Use drug
// 	fmt.Println("=======================Use Goods===================")
// 	models.ActConsume(c1, models.GetGoods(c1.Goods, models.FoodType, 0), []int{0})
// 	fmt.Println(c1.Status)
// 	fmt.Println("=======================Use Goods===================")
// 	models.ActConsume(c1, models.GetGoods(c1.Goods, models.FoodType, 0), []int{0})
// 	fmt.Println(c1.Status)

// 	fitness := &models.Fitness{10, &models.Location{2, 5}, 100, 50}
// 	models.ActVisit(c1, fitness, []int{0})
// 	PrintCharacter(c1)
// }

// func PrintGoods(g []models.GoodsHelper) {
// 	for _, val := range g {
// 		fmt.Println(val)
// 	}
// }
