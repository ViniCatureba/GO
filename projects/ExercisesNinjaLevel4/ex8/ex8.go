package main

import "fmt"

func main() {
	x := map[string]string{
		"Vinicius_Catureba": "Sorvete",
		"Joao_Marques":      "Pastel",
		"Igor_Reis":         "Gym",
	}
	for t := len(x); t == 0; t++ {
		fmt.Printf("%v\t%v\n", t, x)
	}

}
