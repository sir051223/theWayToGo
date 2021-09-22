package main

import "fmt"

type Rope string

func main() {
	var a, b Rope = "string A", "string B"
	fmt.Printf("a has the value: %s \n", a)
	fmt.Printf("b has the value: %s", b)
}
