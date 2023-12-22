package main

import "fmt"

func main() {
	x := map[string][]string{
		"eduar_tua":    []string{`computadoras`, `montaña`, `playa`},
		"carlos_ramon": []string{`leer`, `comprar`, `música`},
		"juan_bimba":   []string{`helado`, `pintar`, `bailar`},
	}
	fmt.Println(x)

	for k, v := range x {
		fmt.Println("Registro: ", k)
		for i, val := range v {
			fmt.Println("\t", i, val)
		}
	}
}
