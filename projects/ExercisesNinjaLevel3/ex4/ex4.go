package main

import "fmt"

var x int = 2002

func main() {
	for {

		if x >= 2002 && x < 2022 {
			fmt.Printf("I was alive in %v\n", x)
			x++
		} else if x == 2022 {
			fmt.Printf("I am alive in %v\n", x)
			x++
		} else {
			break
		}

	}
}
