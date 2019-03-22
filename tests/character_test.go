package tests

import (
	"fmt"
	"reflect"

	"github.com/over-engineering/msa/models"
)

// func TestCraracter(t *testing.T) {
func Tmp2() {
	CreateMentalPTable()

	cList := CreateCharacters()
	c1 := cList[0]
	c2 := cList[1]
	c3 := cList[2]

	// add freindship
	models.AddFriendships(c1, c2, 10)
	models.AddFriendships(c1, c2, 20)
	models.AddFriendships(c1, c2, -10)
	models.AddFriendships(c2, c3, 5)
	models.AddFriendships(c2, c3, -5)
	models.AddFriendships(c2, c3, 3)
	models.AddFriendships(c3, c1, 11.5)

	// print
	PrintCharacter(c1)
	PrintCharacter(c2)
	PrintCharacter(c3)

	// buy goods
	BuyGoods(c1)

	// print
	PrintCharacter(c1)
	PrintCharacter(c2)
	PrintCharacter(c3)
}

func CreateCharacters() []*models.Character {
	c1 := &models.Character{
		ID:        0,
		FirstName: "chan",
		LastName:  "park",
		Job:       models.JOB_NONE,
		Status:    models.StarterStatus(),
		// Conditions: ,
		// Equipments: ,
		// Goods:   map[models.GoodsHelper]models.GoodsHelper{},
		Goods: [][]models.GoodsHelper{
			models.PhoneType: []models.GoodsHelper{},
			models.CarType:   []models.GoodsHelper{},
			models.HouseType: []models.GoodsHelper{},
			models.FoodType:  []models.GoodsHelper{},
		},
		Finance: models.Finance{Balance: 500, TaxInfo: models.TaxInfo{AnnualIncome: 0}},
		// Contract: ,
		// FanInfo: ,
		Friendships: models.Friendships{},
		// Team: ,
	}

	c2 := &models.Character{
		ID:        1,
		FirstName: "junmo",
		LastName:  "park",
		Job:       models.JOB_SOCCER_PLAYER,
		Status:    models.StarterStatus(),
		// Conditions: ,
		// Equipments: ,
		// Goods:   map[models.GoodsHelper]models.GoodsHelper{},
		Goods: [][]models.GoodsHelper{
			models.PhoneType: []models.GoodsHelper{},
			models.CarType:   []models.GoodsHelper{},
			models.HouseType: []models.GoodsHelper{},
			models.FoodType:  []models.GoodsHelper{},
		},
		Finance: models.Finance{Balance: 500, TaxInfo: models.TaxInfo{AnnualIncome: 0}},
		// Contract: ,
		// FanInfo: ,
		Friendships: models.Friendships{},
		// Team: ,
	}

	c3 := &models.Character{
		ID:        2,
		FirstName: "changhyun",
		LastName:  "park",
		Job:       models.JOB_BASEBALL_PLAYER,
		Status:    models.StarterStatus(),
		// Conditions: ,
		// Equipments: ,
		// Goods:   map[models.GoodsHelper]models.GoodsHelper{},
		Goods: [][]models.GoodsHelper{
			models.PhoneType: []models.GoodsHelper{},
			models.CarType:   []models.GoodsHelper{},
			models.HouseType: []models.GoodsHelper{},
			models.FoodType:  []models.GoodsHelper{},
		},
		Finance: models.Finance{Balance: 500, TaxInfo: models.TaxInfo{AnnualIncome: 0}}, // FanInfo: ,
		// Contract: ,
		// FanInfo: ,
		Friendships: models.Friendships{},
		// Team: ,
	}

	return []*models.Character{c1, c2, c3}
}

func PrintCharacter(c *models.Character) {
	fmt.Println(c.FirstName, c.LastName)

	// v := reflect.ValueOf(*c)
	fields := reflect.TypeOf(*c)
	values := reflect.ValueOf(*c)
	for i := 0; i < fields.NumField(); i++ {
		field := fields.Field(i)
		value := values.Field(i)
		fmt.Print(field.Name, ": ", value, "\n")
	}

	fmt.Println("Status:")
	PrintStatus(c.Status)

	fmt.Println("Friendships of ID", c.ID)
	for key, value := range c.Friendships {
		fmt.Println("ID", key, "freindship:", *value)
	}

}

func BuyGoods(c *models.Character) {
	// smartPhone := &models.SmartPhone{
	// 	Goods: models.Goods{
	// 		ID:    10,
	// 		Name:  "Galaxy S10",
	// 		Brand: "Samsung",
	// 		Price: 100,
	// 		// Duration:    1000, // day
	// 		Description: "Samsung Galaxy S10 is a line of Android smartphones manufactured by Samsung Electronics. The Galaxy S10 series is a celebratory series of the 10th anniversary of the Samsung Galaxy S flagship line. Unveiled during a press event, Samsung Galaxy Unpacked 2019 on February 20, 2019, they started shipping on March 8, 2019, and in some regions such as Australia and the United States, they started shipping them on March 6, 2019.",
	// 		// Effects: models.Effects{
	// 		// 	models.Effect{"Mentals", models.Mentals{
	// 		// 		models.Confidence: &models.Mental{models.Confidence, float32(20)},
	// 		// 		models.Ambition:   &models.Mental{models.Ambition, float32(5)},
	// 		// 	},
	// 		// 		models.Effect{"Friendship", models.AddFriendships()},
	// 		// 	}},
	// 	},
	// 	Payment: 10,
	// 	Memory:  8,    // GB
	// 	Battery: 3400, // mAh
	// }

	// act := models.ActivityManager{
	// 	Character: c,
	// 	// Facility:
	// 	// Job:
	// }
	// act.BuyGoods(smartPhone)
	// fmt.Println("My goods map: ", c.Goods[smartPhone])
	// smartPhone.ApplyEffects()

}

func CreateMentalPTable() {
	pMentalMap := map[interface{}]float32{
		models.Ambition:      1.6,
		models.Boldness:      1.2,
		models.Aggression:    0.9,
		models.Predictation:  1.11,
		models.Composure:     1.3,
		models.Concentration: 1.2,
		models.Immersion:     1.5,
		models.Competition:   0.98,
		models.SelfEsteem:    0.9,
		models.Confidence:    0.8,
		models.Attention:     1.0,
	}
	PMentalList := [16]map[interface{}]float32{}
	for i := 0; i < 16; i++ {
		PMentalList[i] = pMentalMap
	}
	models.PTable.AddPersonalityCoe(PMentalList)

	// fmt.Println(pTable)
}
