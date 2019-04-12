package utils

import "fmt"

// Describer is the interface for Describe method.
type Describer interface {
	Describe() string
}

// PrintDescription prints desciption to log.
func PrintDescription(d Describer) {
	fmt.Println(d.Describe())
}
