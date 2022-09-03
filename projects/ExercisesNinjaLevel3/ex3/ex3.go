package main

import "fmt"

func main() {
	for i := 2002; i <= 2022; i++ {
		if i == 2022 {
			fmt.Printf("I am alive in %v", i)
		} else {
			fmt.Printf("I was alive in %v\n", i)
		}

	}
}
