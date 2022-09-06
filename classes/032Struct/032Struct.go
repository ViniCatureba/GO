package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

func main() {
	p1 := person{
		first: "Vinicius",
		last:  "Bigao",
		age:   19,
	}
	p2 := person{
		first: "Joao",
		last:  "Marques",
		age:   20,
	}
	fmt.Println(p1)
	fmt.Println(p2)
	fmt.Println(p1.first, p1.last, p1.age)
	fmt.Println(p2.first, p2.last, p2.age)
}
