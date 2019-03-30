package tests

import (
	"fmt"
	"time"

	"github.com/over-engineering/msa/models/character"
	"github.com/over-engineering/msa/models/condition"
	"github.com/over-engineering/msa/models/types"
)

func ExampleNewCharacter() {
	newChar := character.NewCharacter(
		types.UID("Character#001"),
		"Chan",
		"Park",
		character.NewStatus(),
	)

	drugAddiction := condition.NewDrugAddiction(
		float32(10),
		time.Now(),
		time.Duration(5*time.Hour),
	)

	fmt.Println(*newChar)
	fmt.Println(drugAddiction)
	newChar.AddCondition(drugAddiction)
	newChar.AddCondition(drugAddiction)
	fmt.Println(*newChar)
	err := newChar.RemoveCondition(0)
	fmt.Println(err)
	fmt.Println(*newChar)
	newChar.UpdateCondition(time.Now().Add(time.Duration(5 * time.Hour)))
	fmt.Println(*newChar)

	// Output:
	// {Character#001 Chan Park 0 {{false 0001-01-01 00:00:00 +0000 UTC 0001-01-01 00:00:00 +0000 UTC 0} 0 0 0 map[] map[] map[] 0 0 0 0 50 0 0 0 0 map[] map[]} {} []}
	// {DrugAddiction 10 2h46m40s}
	// {Character#001 Chan Park 0 {{false 0001-01-01 00:00:00 +0000 UTC 0001-01-01 00:00:00 +0000 UTC 0} 0 0 0 map[] map[] map[] 0 0 0 0 40 0 0 0 0 map[] map[]} {} [{DrugAddiction 10 2h46m40s}]}
	// <nil>
	// {Character#001 Chan Park 0 {{false 0001-01-01 00:00:00 +0000 UTC 0001-01-01 00:00:00 +0000 UTC 0} 0 0 0 map[] map[] map[] 0 0 0 0 50 0 0 0 0 map[] map[]} {} []}
}
