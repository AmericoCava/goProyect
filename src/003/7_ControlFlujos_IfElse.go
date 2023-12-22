package main

import f "fmt"

func main() {
	x := 42

	if x == 40 {
		f.Println("El valor de x es: 40")
	} else if x == 41 {
		f.Println(" El valor de x es: 41")
	} else {
		f.Println("El valor de x es: ", x)
	}
}
