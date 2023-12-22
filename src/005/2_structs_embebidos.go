package main

import "fmt"

type persona struct {
	nombre   string
	apellido string
	edad     int
}

type agenteSecreto struct {
	persona
	lpm bool
}

func main() {
	as1 := agenteSecreto{
		persona: persona{
			nombre:   "James",
			apellido: "Bond",
			edad:     32,
		},
		lpm: true,
	}
	fmt.Println(as1)
	fmt.Println(as1.nombre, as1.apellido, as1.edad, as1.lpm)
}
