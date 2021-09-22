package main

import (
	"fmt"
)

const c = "C"

const pi float32 = 3.14159

var v int = 5

type (
	IZ int
)

type T struct {}

func init() { // initialization of package
	fmt.Println("initialization test")
}

func main() {
	var a int
	b := 5.0
	c := int(b)
	fmt.Println(c)
	Func1()
	fmt.Println(a)
}

func Func1() {
	var a IZ = 6
	c := int(a)
	d := IZ(c)
	fmt.Println(a)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(pi)
}

func (t T) Method() {
	
}