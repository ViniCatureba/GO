package main

import "fmt"

func main() {
	fmt.Println("Hello, class!")
	foo()
	fmt.Println(" smth more")

	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}

	}

	bar()
}
func foo() {
	fmt.Println("I am in foo")
}

func bar() {
	fmt.Println("and then we exited")
}
