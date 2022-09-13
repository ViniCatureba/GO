package main

import (
	"fmt"
)

type person struct {
	first_name    string
	last_name     string
	fav_ice_cream []string
}

func main() {
	p1 := person{
		first_name: "vinicius",
		last_name:  "bigao",
		fav_ice_cream: []string{
			"vannila",
			"chocolate",
			"oreo",
		},
	}

	p2 := person{
		first_name: "Isa",
		last_name:  "Bigao",
		fav_ice_cream: []string{
			"pistache",
			"coffe",
			"coke",
		},
	}

	fmt.Println(p1.first_name)
	fmt.Println(p1.last_name)
	for i, v := range p1.fav_ice_cream {
		fmt.Println(i, v)

	}

	fmt.Println(p2.first_name)
	fmt.Println(p2.last_name)
	for i, v := range p2.fav_ice_cream {
		fmt.Println(i, v)
	}

}
