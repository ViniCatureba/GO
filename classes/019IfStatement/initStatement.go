package main

import "fmt"

func exemple2() {
	if x := 42; x == 2 { // x will be only in the scope of this if statement
		fmt.Println("002")
	}
	fmt.Println("here is a statement")
	fmt.Println("something else")
}
