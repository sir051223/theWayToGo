package main

import (
	"fmt"
	"strings"
)

func main() {
	var origs string = "Hi there!"
	var newS string

	newS = strings.Repeat(origs, 5)
	fmt.Printf("The new repeated string is: %s\n", newS)
}