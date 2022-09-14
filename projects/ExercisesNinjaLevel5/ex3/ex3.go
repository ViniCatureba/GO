package main

import "fmt"

type car struct {
	door  int
	color string
}

type truck struct {
	car
	fourWheel bool
}

type sedan struct {
	car
	luxury bool
}

func main() {
	ford := truck{
		car: car{
			door:  4,
			color: "red",
		},
		fourWheel: true,
	}

	corolla := sedan{
		car: car{
			door:  4,
			color: "grey",
		},
		luxury: false,
	}

	fmt.Println(ford)
	fmt.Println(corolla)

}
