package main

import (
	"fmt"
	"strings"
)

func main() {
	var str string = "This is an example of a string"
	fmt.Printf("T/F? Does the string \"s%\" have prefix %s ", str, "This")
	fmt.Printf("%t\n", strings.HasPrefix(str, "This"))
	fmt.Printf("T/F? Does the string \"s%\" contains %s ", str, "example")
	fmt.Printf("%t\n", strings.Contains(str, "example"))
}