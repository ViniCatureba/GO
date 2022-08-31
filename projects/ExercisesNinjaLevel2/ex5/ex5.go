package main

import "fmt"

var x string

func main() {
	x = `Ola,
	tudo bem?
	this is a "raw string"`
	fmt.Println(x)
}
