package main

import "fmt"

func main() {
	x := sum(2, 3, 4, 5, 6, 7, 8, 9)
	fmt.Println("The total is,", x)
}

func sum(x ...int) int {
	fmt.Println("Hello, playgournd")

	fmt.Println(x)
	fmt.Printf("%T\n", x)

	sum := 0

	for i, v := range x {
		sum += v
		fmt.Println("for item in index possitions", i, " we are now adding", v, " to the tootal wich is now ", sum)
	}

	return sum
}
