package main

import "fmt"

func main() {
	x := min(1, 3, 2, 0)
	fmt.Printf("The mininum is: %d\n", x)
	arr := []int{7,9,3,5,1,0,-1}
	x = min(arr...)
	fmt.Printf("The mininum in the array arr is: %d", x)
}

func min(a ...int) int {
	if len(a)==0 {
		return 0
	}
	min := a[0]
	for _, v := range a {
		if v < min {
			min = v
		}
	}
	return min
}
