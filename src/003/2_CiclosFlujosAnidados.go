package main

import "fmt"

func main() {
	for i := 0; i <= 100; i++ {
		fmt.Printf("El ciclo externo %d\n", i)
		for j := 0; j <= 3; j++ {
			fmt.Printf("El ciclo externo %d\tel ciclo interno %d\n", i, j)
		}
	}
}
