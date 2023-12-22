package main

import "fmt"

type persona struct {
	nombre   string
	apellido string
	edad     int
}

func main() {
	p1 := persona{
		nombre:   "Americo",
		apellido: "Cavallieri",
		edad:     34,
	}

	p2 := persona{
		nombre:   "Condor",
		apellido: "Perez",
		edad:     80,
	}

	fmt.Println(p1)
	fmt.Println(p2)

	fmt.Println(p1.nombre, p1.apellido, p1.edad)
	fmt.Println(p2.nombre, p1.apellido, p2.edad)
}
