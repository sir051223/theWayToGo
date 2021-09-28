package main

import "fmt"

var a int = 5
var b int = 3

var vsum, vdiff, vproduct int

func main() {
	vsum, vdiff, vproduct = f1(a, b);
	fmt.Printf("a,b sum = %d, diff = %d product = %d \n", vsum, vdiff, vproduct)

	vsum, vdiff, vproduct = f2(a, b);
	fmt.Printf("a,b sum = %d, diff = %d product = %d ", vsum, vdiff, vproduct)
}

func f1( x int, y int) (int, int, int)  {
	sum := x + y
	diff := x -y
	product := x * y
	return sum, diff, product
}

func f2(x int, y int) (sum int, diff int, product int) {
	sum = x + y
	diff = x -y
	product = x * y
	return
}