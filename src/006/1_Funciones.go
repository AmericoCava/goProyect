package main

import (
	"fmt"
)

func main() {
	foo()
	bar("Americo")
	s1 := woo("Americo")
	fmt.Println(s1)
	x, y := saludar("Americo", "Cavallieri")
	fmt.Println(x)
	fmt.Println(y)
}

// func (r receptor) identificador(parametros) retorno(s...) {codigo}
func foo() {
	fmt.Println("Hola")
}

func bar(s string) {
	fmt.Println("Hola,", s)
}

func woo(st string) string {
	return fmt.Sprint("Holaassssss,", st)
}

func saludar(n string, a string) (string, bool) {
	p := fmt.Sprint(n, " ", a, " dice hola.")
	q := true
	return p, q
}
