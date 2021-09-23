package main

import (
	"fmt"
	"strconv"
)

func main() {
	var orig string = "123" // ABC
	var newS string

	fmt.Printf("The size of ints is:%d\n", strconv.IntSize)
	an, error := strconv.Atoi(orig)
	if error != nil {
		fmt.Printf("orig %s is not an integer - exiting with error %s\n", orig, error)
		return
	}
	fmt.Printf("The integer is %d\n", an)
	an = an + 5
	newS = strconv.Itoa(an)
	fmt.Printf("The new string is: %s\n", newS)
}
