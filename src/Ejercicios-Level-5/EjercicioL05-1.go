package main

import "fmt"

type persona struct {
	nombre     string
	apellido   string
	saboresFav []string
}

func main() {
	p := persona{
		nombre:   "Americo",
		apellido: "Cavallieri",
		saboresFav: []string{
			"Chocolate",
			"Limon",
			"Maracuya",
		},
	}
	fmt.Println(p)
	fmt.Println(p.nombre)
	fmt.Println(p.apellido)

	for i, v := range p.saboresFav {
		fmt.Println("\t", i, v)
	}
}
