package main

import "fmt"

func main() {
	p := struct {
		nombre   string
		apellido string
		edad     int
	}{
		nombre:   "Americo",
		apellido: "Cavallieri",
		edad:     34,
	}
	fmt.Println(p)
}
