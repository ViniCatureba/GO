package main

import "fmt"

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

	m := map[string]person{
		p1.last_name: p1,
		p2.last_name: p2,
	}

	fmt.Println(m)

	for k, v := range m {
		fmt.Println(k, v)
	}
}
