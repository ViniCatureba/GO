package main

import "fmt"

func main() {
	x := []int{11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	y := x[0:5]
	z := x[5:]
	m := x[2:7]
	t := x[1:6]
	fmt.Println(y)
	fmt.Println(z)
	fmt.Println(m)
	fmt.Println(t)

}