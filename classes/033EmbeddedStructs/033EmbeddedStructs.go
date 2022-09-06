package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}
type secretAgent struct {
	person
	ltk bool
}

func main() {
	sa1 := secretAgent{
		person: person{
			first: "James",
			last:  "Bond",
			age:   32,
		},
		ltk: true,
	}
	p2 := person{
		first: "Joao",
		last:  "Marques",
		age:   20,
	}

	fmt.Println(sa1)
	fmt.Println(p2)
	fmt.Println(sa1.person.first, sa1.person.last, sa1.person.age, sa1.ltk)
	fmt.Println(p2.first, p2.last, p2.age)

}
