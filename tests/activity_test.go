package tests

import (
	"testing"

	"github.com/over-engineering/msa/models"
)

func TestActivity(t *testing.T) {
	cList := CreateCharacters()
	c1 := cList[0]
	// c2 := cList[1]
	// c3 := cList[2]

	act := models.ActivityManager{
		Character: c1,
		// Facility:
		// Job:
	}

	PrintCharacter(c1)
	apple := models.NewFood(models.AppleID, "Apple", 100, 300)
	act.BuyGoods(apple)
	act.Eat(apple, 1)

	PrintCharacter(c1)

}
