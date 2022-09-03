package main

import "fmt"

func main() {
	x := 1
	for {
		if x == 1 {
			fmt.Println(x)
			x++
		} else if x == 2 {
			fmt.Println(x)
			x++
		} else {
			fmt.Println(x)
			break
		}
	}
}
