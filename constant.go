package main

import "fmt"

const (
	a = iota
	b = iota
	c = iota
)

const (
	Sunday = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

func main() {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(Saturday)
}



