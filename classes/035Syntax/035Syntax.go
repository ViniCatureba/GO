package main

import "fmt"

//func (r receiver) identifier(parameters) (return(s)) { ... }

func main() {
	foo()
	bar("james")
	s1 := woo("Moneypenny")
	fmt.Println(s1)
	x, y := mouse("ian", "fleming")
	fmt.Println(x)

	fmt.Println(y)
}

func foo() {
	fmt.Println("hellow from foo")
}

//everyting in Go is pass by value
func bar(s string) { //take an argument
	fmt.Println("Hello,", s)
}

func woo(st string) string {
	return fmt.Sprint("Hellow from woo,", st)
}

func mouse(fn string, ln string) (string, bool) {
	a := fmt.Sprint(fn, " ", ln, `,says"hello"`)
	b := false
	return a, b

}
