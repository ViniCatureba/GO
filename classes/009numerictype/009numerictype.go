package main

import (
	"fmt"
	"runtime"
)

var x int

var y float32

var z int8 = 127 //int8 can oly handle number between -128 to 127

func main() {
	x = 42

	y = 42.34543

	fmt.Println(x)
	fmt.Println(y)
	fmt.Printf("%T\n", x)
	fmt.Printf("%T\n", y)

	fmt.Printf("%T\n", z)
	fmt.Println(z)
	fmt.Println(runtime.GOOS)
	fmt.Println(runtime.GOARCH)

}
