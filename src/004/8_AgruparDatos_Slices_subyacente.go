package main

import "fmt"

func main() {
	x := []int{41, 42, 43, 44, 45, 46, 47, 48}
	fmt.Println(len(x))
	fmt.Println(cap(x))

	x = append(x, 49, 50, 51, 52) //NUEVO ARREGLO SUBYANCENTE POR FALTA DE CAPACIDAD
	fmt.Println(len(x))
	fmt.Println(cap(x))

	y := append(x[:3], x[5:]...)
	fmt.Println(len(y))
	fmt.Println(cap(y))
}
