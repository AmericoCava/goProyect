package main

import (
	"fmt"
)

var a int
var b string = "Programa"
var c bool
var d bool = true

func main(){
	e := 42
	f := "dice Hola Mundo."
	g := `El doctor dice que comer vegeatles es
			saludable.`
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)
	fmt.Println(g)
	fmt.Println("Mi", b, f)

	//PRINT PARA OINSERTAR VARIABLES ESPECIALES DE GO %v %s %d
	fmt.Printf("el valor de la variavles es: %v\n", a)
	fmt.Printf("%v\n", b)
	fmt.Printf("%T\n", b)

	//PRINT DE GO QUE SE LLAMA SPRINT QUE TRANSFORMA TODO A STRING
	s1:= fmt.Sprint(e)
	fmt.Println(s1)
}