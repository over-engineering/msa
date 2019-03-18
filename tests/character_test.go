package tests

import (
	"fmt"
	"testing"

	"github.com/over-engineering/msa/models"
)

func TestCraracter(t *testing.T) {
	c1 := &models.Character{
		ID:        0,
		FirstName: "chan",
		LastName:  "park",
		Job:       models.JOB_NONE,
		Status:    models.StarterStatus(),
		// Conditions: ,
		// Equipments: ,
		// Goods: ,
		// Finance: ,
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
		// Goods: ,
		// Finance: ,
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
		// Goods: ,
		// Finance: ,
		// Contract: ,
		// FanInfo: ,
		Friendships: models.Friendships{},
		// Team: ,
	}

	models.AddFriendships(c1, c2, 10)
	models.AddFriendships(c1, c2, 20)
	models.AddFriendships(c1, c2, -10)
	models.AddFriendships(c2, c3, 5)
	models.AddFriendships(c2, c3, -5)
	models.AddFriendships(c2, c3, 3)
	models.AddFriendships(c3, c1, 11.5)

	PrintCharacter(c1)
	PrintCharacter(c2)
	PrintCharacter(c3)
}

func PrintCharacter(c *models.Character) {
	fmt.Println(c.FirstName, c.LastName)

	fmt.Println("Status: ")
	PrintStatus(c.Status)

	fmt.Println("Friendships of ID", c.ID)
	for key, value := range c.Friendships {
		fmt.Println("ID", key, "freindship:", *value)
	}

}
