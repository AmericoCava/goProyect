package main

import "fmt"

func main() {
	s := struct {
		nombre     string
		amigos     map[string]int
		bebidasFav []string
	}{
		nombre: "Americo",
		amigos: map[string]int{
			"Pato":   222,
			"juan":   555,
			"condor": 777,
		},
		bebidasFav: []string{
			"Fanta",
			"Bils",
			"Pap",
		},
	}
	fmt.Println(s)
}
