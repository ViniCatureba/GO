package main

import "fmt"

func main() {
	x := []int{4, 5, 7, 8, 42}
	fmt.Println(x[0])
	fmt.Println(x)
	fmt.Println(x[1:])  // will give me from the 1st position until the end of the slice
	fmt.Println(x[1:3]) // will give me from the 1st position until the 3rd number of the slice ([0,1,2])

	for i := 0; i < len(x); i++ {
		fmt.Println(i, x[i])
	}

}
