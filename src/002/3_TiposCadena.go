package main

import (
	"fmt"
)

func main() {
	s1 := "Hola Mundo"
	fmt.Println(s1)

	bs := []byte(s1)
	fmt.Println(bs)
	fmt.Printf("%T\n", bs)

	for i := 0; i < len(s1); i++ {
		fmt.Printf("%#U\n", s1[i])
	}

	for i, v := range s1 {
		fmt.Printf("En el Indice %d el valor es %#x\n", i, v)
	}
}
