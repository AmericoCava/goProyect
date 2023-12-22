package main

import "fmt"

func main() {
	x := []string{"Batman", "Jefe", "Vestido de negro"}
	y := []string{"Robin", "Ayudante", "Vestido de colores"}

	fmt.Println(x)
	fmt.Println(y)

	z := [][]string{x, y}
	fmt.Println(z)

	for i, reg := range z {
		fmt.Println("Registro: ", i)
		for j, val := range reg {
			fmt.Printf("\tIndice: %v\tvalor: %v\n", j, val)
		}
	}
}
