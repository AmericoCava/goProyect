package main

import f "fmt"

func main() {
	switch {
	case 2 == 4:
		f.Println("No Deberia Imprimir")
	case 3 == 3:
		f.Println("Deberia Imprimir")
		fallthrough
	case 4 == 5:
		f.Println(" No deberia imprimir2")
	default:
		f.Println("Imprimiendo desde default")
	}
}
