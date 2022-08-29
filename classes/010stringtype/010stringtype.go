package main

import "fmt"

func main() {
	s := `"Hello, 
world"`
	fmt.Println(s)
	fmt.Printf("%T\n", s)

	d := "Hello, world"

	fmt.Println(d)
	fmt.Printf("%T\n", d)

	for i := 0; i < len(s); i++ {
		fmt.Printf("%#U", s[i])
	}

	fmt.Println("")

	for i, v := range s {
		fmt.Printf("at index possition %d we have hex %#x\n", i, v)
	}
}
