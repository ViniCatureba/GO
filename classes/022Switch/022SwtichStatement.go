package main

import (
	"fmt"
)

func main() {
	switch {
	case false:
		fmt.Println("This should not print")
	case (2 == 4):
		fmt.Println("this should not print")
	case (3 == 3):
		fmt.Println("prints")
		fallthrough //it allows to keep going to the next true case
	case (7 == 9):
		fmt.Println("not true ")
		fallthrough
	case (11 == 14):
		fmt.Println("not ture 2")
		fallthrough
	case (15 == 15):
		fmt.Println("true 15")
	default:
		fmt.Println("this is default!")

	}
	switch {
	case (0 > 1):
		fmt.Println("False")
	default:
		fmt.Println("switch did not found any match so it keep default!")
	}

	switch "Bond" {
	case "money":
		fmt.Println("this is money")
	case "Bond":
		fmt.Println("this is bond")
	case "hey":
		fmt.Println("this is hey")
	default:
		fmt.Println("default")
	}

	n := 2
	switch n == 2 {
	case n == 1:
		fmt.Println("this is 1")
	case n == 2:
		fmt.Printf("this is %v\n", n)
	}

	m := "hi"
	switch m {
	case "Moneypenny", "bond", "hi":
		fmt.Println("this has hi")
	default:
		fmt.Println("default")

	}

}
