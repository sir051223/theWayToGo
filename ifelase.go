package main

import "fmt"

func main() {
	var first int = 1
	var cond int

	if first <= 0 {
		fmt.Println("first is less than or equal to 0")
	} else if first > 0 && first < 5 {
		fmt.Println("first is between 0 and 5")
	} else {
		fmt.Println("first is 5 or greater")
	}

	if cond = 10; cond > 10 {
		fmt.Println("cond is greater than 10")
	} else {
		fmt.Println("cond is not greater than 10")
	}
}