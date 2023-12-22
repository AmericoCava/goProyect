package main

import "fmt"

type persona struct {
	nombre     string
	apellido   string
	saboresFav []string
}

func main() {
	p1 := persona{
		nombre:   "Americo",
		apellido: "Cavallieri",
		saboresFav: []string{
			"Chocolate",
			"Limon",
			"Maracuya",
		},
	}
	p2 := persona{
		nombre:   "Eduard",
		apellido: "Tua",
		saboresFav: []string{
			"Fresa",
			"piña",
			"Chirimoya",
		},
	}
	m := map[string]persona{
		p1.apellido: p1,
		p2.apellido: p2,
	}
	fmt.Println(m)

	for _, v := range m {
		fmt.Println(v.nombre)
		fmt.Println(v.apellido)
		for i, v := range v.saboresFav {
			fmt.Println(" ", i, v)
		}
		fmt.Println("----------------")
	}
}
