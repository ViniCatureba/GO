package main

import "fmt"

// if you ever need to declarate a variable outside a function you gotta you var, cannot use the short declaration.
// DECLARE the variable "y"
// ASSIGN the value 43
// declare & assign = initilization
var y = 43

//DECLARE there is a variable with the IDENTIFIER "z"
//and that the VARIABLE  with the IDENTIFIER "z" is of TYPE int
// ASSIGNS the ZERO value of TYPE in to "z"
// false for booleans, 0 for integers, 0.0 for floats, "" for strings,
// and nil for pointers, functions, interfaces, slices, channels, and maps.
var z int

func main() {
	//short declaration operator
	// DECLARE a variable and ASSIGN a VALUE
	x := 42
	println(x)

	fmt.Println(y)

	foo()
	println(z)
	
	
	//it is also possible to use this var declarations:
	
	var (
		t = 43
		g = 33
	)
	
	b := 13
	fmt.Println(t, g, b)
}

func foo() {
	fmt.Println("again:", y)
}
