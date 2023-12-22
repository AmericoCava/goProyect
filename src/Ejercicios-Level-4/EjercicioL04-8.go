package main

import "fmt"

func main() {
	m := map[string][]string{
		"eduar_tua":    []string{`computadoras`, `montaña`, `playa`},
		"carlos_ramon": []string{`leer`, `comprar`, `música`},
		"juan_bimba":   []string{`helado`, `pintar`, `bailar`},
	}

	m["luis_perez"] = []string{"jugar", "pintar", "futbol"}

	for k, v := range m {
		fmt.Println("Registro: ", k)
		for i, val := range v {
			fmt.Println("\t", i, val)
		}
	}
}
