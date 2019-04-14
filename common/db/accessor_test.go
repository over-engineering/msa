package db

import (
	"fmt"

	"github.com/over-engineering/msa/common/starter"
)

const (
	path = "mongodb://localhost:27017"
	db   = "oven"
	c    = "character"
)

func ExampleGetCharacter() {
	m := New(path, db, c)
	char, _ := starter.CreateNewCharacter(
		"ID0000000001",
		"1995-02-02T15:04:05+09:00",
		"Chan", "Park",
		1,
		70, 40, 70, 40, 50, 30, 50, 30,
		180, 70, 15,
	)
	// fmt.Println(*char)
	if err := m.PostCharacter(*char); err != nil {
		fmt.Println(err)
	}
	// res, _ := m.GetCharacter("ID0000000001")
	// fmt.Print(res)
	// Output:
}
