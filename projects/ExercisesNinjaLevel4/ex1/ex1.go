package main

import (
	"fmt"
)

func main() {
	x := [5]int{1, 2, 3, 4, 9}

	for i, v := range x {
		fmt.Println(i, v)

	}
	fmt.Printf("%T", x)
}
