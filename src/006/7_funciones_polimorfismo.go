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
	fmt.Println("Hola soy", a.nombre, a.apellido, "el agente secreto se presenta")
}

func (b persona) presentar() {
	fmt.Println("Hola soy", b.nombre, b.apellido, "la persona se presenta")
}

type humano interface {
	presentar()
}

func bar(h humano) {
	switch h.(type) {
	case persona:
		fmt.Println("Fui pasado a la funcion bar (persona)", h.(persona).nombre)
	case agenteSecreto:
		fmt.Println("Fui pasado a la funcion bar (agente secreto)", h.(agenteSecreto).nombre)
	}
	fmt.Println("Fui pasado a la funcion bar", h)
}

func main() {
	as1 := agenteSecreto{
		persona: persona{
			nombre:   "James",
			apellido: "Bond",
		},
		lpm: true,
	}

	as2 := agenteSecreto{
		persona: persona{
			nombre:   "Juan",
			apellido: "Perez",
		},
		lpm: true,
	}

	p := persona{
		nombre:   "Americo",
		apellido: "Cavallieri",
	}

	bar(as1)
	bar(as2)
	bar(p)
}
