package main

import "fmt"

const (
	y = 1 << iota
	x
)

func main() {
	fmt.Printf("%d\t%b\t%#x\n", y, y, y)

	fmt.Printf("%d\t%b\t%#x\n", x, x, x)
	fmt.Println("")
	x := 42
	y := x << 1

	fmt.Printf("%v\t%b\t%#x\n", y, y, y)

	fmt.Printf("%d\t%b\t%#x\n", x, x, x)

}
