package main

import "fmt"

func main() {
	x := make([]int, 10, 12)
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))

	x[0] = 42
	x[9] = 999

	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))

	x = append(x, 3423)

	fmt.Println(x)

	x = append(x, 3424)
	x = append(x, 3425)
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x)) // now the compiler doubles the capasity of this array,bc we added more indexes than its capacities
}
