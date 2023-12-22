package main

import "fmt"

var x int

type persona struct {
	nombre   string
	apellido string
}

func main() {
	p := persona{
		nombre:   "Americo",
		apellido: "Cavallieri",
	}
	fmt.Println(p)
}
