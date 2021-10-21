package main

import "fmt"

func main() {
	i := []string{"i", "am", "test", "some"}
	res := childfun(i...)
	fmt.Println(res)
}

func childfun(a ...string) int {
	if len(a) == 0 {
		return 0
	}
	for _, v := range a {
		fmt.Println(v)
	}
	return 1
}