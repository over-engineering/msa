package starter

import "fmt"

func ExampleCreateNewCharacter() {
	nc, _ := CreateNewCharacter(
		"1995-02-02T15:04:05+09:00",
		"Chan", "Park",
		1,
		70, 40, 70, 40, 50, 30, 50, 30,
		180, 70, 15,
	)
	fmt.Println(nc)
	// Output:
	// Hello
}
