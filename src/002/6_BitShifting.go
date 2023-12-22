package main

import "fmt"

func main() {
	a := 4
	fmt.Printf("%d\n", a) //SISTEMA DECIMAL
	fmt.Printf("%b\n", a) //SISTEMA BINARIO

	b := a << 1
	fmt.Printf("%d\n", b) //SISTEMA DECIMAL
	fmt.Printf("%b\n", b) //SISTEMA BINARIO

	kb := 1024
	gb := kb * 1024
	tb := gb * 1024

	fmt.Printf("%d\t\t\t%b\n", kb, kb)
	fmt.Printf("%d\t\t\t%b\n", gb, gb)
	fmt.Printf("%d\t\t%b\n", tb, tb)
}
