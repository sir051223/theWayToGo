package main

// #include <stdlib.h>
import (
	"C"
)


func main() {
	rd := Random()
	println(rd)
}

func Random() int {
	return int(C.random())
}

func Seed(i int) {
	C.srandom(C.uint(i));
}

