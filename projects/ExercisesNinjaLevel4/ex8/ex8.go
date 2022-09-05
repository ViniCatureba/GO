package main

import (
	"fmt"
)

func main() {
	x := map[string][]string{
		"Vinicius_Catureba": []string{"Sorvete", "cafe", "snowboard"},
		"Joao_Marques":      []string{"Pastel", "praia", "viajar"},
		"Igor_Reis":         []string{"Gym", "carro", "pagode"},
	}
	fmt.Println(x)
	fmt.Println(len(x))

	for k, v := range x {
		fmt.Println("This is the record for ", k)
		for i, v2 := range v {
			fmt.Println("\t", i, v2)
		}
	}
}
