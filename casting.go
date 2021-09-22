package main

import (
	"fmt"
	"math"
)

func main() {
	var c1 complex64 = 5 + 10i
	fmt.Printf("The value is: %v", c1)
	var n int16 = 34
	var m int32
	var a int = 1
	var u uint8
	var i int
	var f float64 = 1222.8
	u,error := Uint8FromInt(a)
	// fmt.Println(u)
	fmt.Printf("uint8 is:%s\n", u)
	fmt.Println(error)
	i = IntFromFloat64(f)
	fmt.Println(i)

	//if error {
	//	fmt.Fprint(error)
	//}
	// m = n // cannot use n (type int16) as type int32 in assignment
	// fmt.Printf(m)
	m = int32(n)
	fmt.Printf("32 bit int is:%s\n", m)
	fmt.Printf("16 bit int is:%5.2e\n", n)
}

func Uint8FromInt(n int) (uint8, error) {
    if 0 <= n && n <= math.MaxUint8 {
		return uint8(n), nil
	}
	return 0, fmt.Errorf("d% is out of the unit8 range", n)
}

func IntFromFloat64(x float64) int {
    if math.MinInt32 <= x && x <= math.MaxInt32 {
		whole, fraction := math.Modf(x)
		if fraction >= 0.5 {
			whole++
		}
		return int(whole)
	}
	panic(fmt.Sprintf("%g is out of the int32 range", x))
}