package main

import "fmt"

type persona struct {
	nombre   string
	apellido string
}

type agenteSecreto struct {
	persona
	lpm bool
}

func (a agenteSecreto) presentar() {
	fmt.Println("Hola soy", a.nombre, a.apellido)
}

func main() {
	as1 := agenteSecreto{
		persona: persona{
			nombre:   "James",
			apellido: "Bond",
		},
		lpm: true,
	}
	fmt.Println(as1)
	as1.presentar()
}
