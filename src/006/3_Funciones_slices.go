package main

import "fmt"

func main() {
	xi := []int{2, 3, 4, 5, 6, 7, 8, 9}
	x := sum(xi...)
	fmt.Println("El valor total alamacenado en la variable es", x)
}

func sum(x ...int) int {
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	fmt.Println(len(x))
	fmt.Println(cap(x))

	suma := 0
	for i, v := range x {
		suma += v
		fmt.Println("El valor en el indice", i, "suma", v, "al total quedando", suma)
	}
	fmt.Println("el total es", suma)

	return suma
}
