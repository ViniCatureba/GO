package main

import "fmt"

func main() {
	ans := struct {
		company_name string
		niche        map[string]int
		num_clients  int
	}{
		company_name: "Ford",
		niche: map[string]int{
			"cars":   1,
			"trucks": 2,
			"bikes":  3,
		},
		num_clients: 1023,
	}
	fmt.Println(ans.company_name)
	for k, v := range ans.niche {
		fmt.Println(k, v)
	}
	fmt.Println(ans.num_clients)

}
