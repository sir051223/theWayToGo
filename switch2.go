package main

import "fmt"

func main() {
	var num1 int = 7

	switch {
		case num1 < 0:
			fmt.Println("Number is negative")
		case num1 > 0 && num1 < 10:
			fmt.Println("Number is between 0 and 10")
		default:
			fmt.Println("Number is 10 or greater")
	}

	k := 6
	switch k {
		case 4: fmt.Println("was <= 4"); fallthrough;
		case 5: fmt.Println("was <= 5"); fallthrough;
		case 6: fmt.Println("was <= 6"); fallthrough;
		case 7: fmt.Println("was <= 7"); fallthrough;
		case 8: fmt.Println("was <= 8"); fallthrough;
		default: fmt.Println("default case")
	}

	month := 12
	switch {
	case month > 0 && month < 4:
		fmt.Println("spring")
	case 3 < month && month < 7:
		fmt.Println("summer")
	case 6 < month && month < 10:
		fmt.Println("autumn")
	case 9 < month && month <13:
		fmt.Println("winter")
	default:
		fmt.Println("error month")
	}
}
