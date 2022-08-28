package main

import "fmt"

var y string
var z int

func main() {
	//DECLARE a varible to be of a certain TYPE
	//AND THEN assing A vaulue OF that type to the variable

	fmt.Println("printing the value of y", y, "ending")
	fmt.Printf("%T\n", y)
	y = "BOND, James bond"

	fmt.Println("print the value of y", y, "ending")
	fmt.Printf("%T", y)
	fmt.Println(y)
	fmt.Println(y)
	fmt.Println(z)
	fmt.Printf("%T", z)
}
