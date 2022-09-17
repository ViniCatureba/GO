package main

import "fmt"

type person struct {
	first string
	last  string
}

type secrecAgent struct {
	person
	ltk bool
}

func (s secrecAgent) speak() {
	fmt.Println("Hey, i am", s.first, s.last)
}

func main() {
	sa1 := secrecAgent{
		person: person{
			"james",
			"Bond",
		},
		ltk: true,
	}

	sa2 := secrecAgent{
		person: person{
			"Miss",
			"Moneypenny",
		},
		ltk: true,
	}

	fmt.Println(sa1)
	fmt.Println(sa2)
	sa1.speak()
	sa2.speak()
}
