package main

import "fmt"

func main() {

	xs1 := []string{"James", "Bond", "Shaken, not stirred"}
	xs2 := []string{"Miss", "Moneypenny", "Helloooooo, James."}

	xxs := [][]string{xs1, xs2}

	fmt.Println(xxs)

	for i, xs := range xxs {
		fmt.Printf("record: %v\n", i)
		for j, val := range xs {
			fmt.Printf("\t index position: %v \t value: \t %v\n", j, val)
		}
	}

}
