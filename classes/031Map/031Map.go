package main

import "fmt"

func main() {
	m := map[string]int{
		"James":           32,
		"Miss Moneypenny": 27,
	}
	fmt.Println(m)

	fmt.Println(m["James"])
	fmt.Println(m["Barnabas"])

	v, ok := m["Barnabas"]
	fmt.Println(v)
	fmt.Println(ok)

	if v, ok := m["Miss Moneypenny"]; ok { // this if statement will verify if this map is true and if is will print.
		fmt.Println("THIS IS THE IF PRINT ", v)
	}

	m["todd"] = 33 // add a key in the map

	for k, v := range m {
		fmt.Println(k, v)
	}

	delete(m, "James")
	fmt.Println(m)

	if t, d := m["Miss Moneypenny"]; d {
		fmt.Println("value: ", t)
		delete(m, "Miss Moneypenny")
	}
	fmt.Println(m)
}
