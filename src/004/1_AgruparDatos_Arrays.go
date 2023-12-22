package main

import f "fmt"

func main() {
	// ARRAY O ARREGLO CON CANTIDAD DEFINIDA DE DATOS
	var x [5]int
	f.Println(x)
	f.Println(len(x))

	//TEST
	var z []int
	z = append(z, 1, 2, 3)
	f.Println(z)
	f.Println(cap(z))
	z = append(z, 5)
	f.Println(cap(z))

	// SLICE PARA CUALQUIER TIPO DE DATO
	var r []interface{}
	r = append(r, 1)
	r = append(r, "Hola")
	r = append(r, 3.14)
	f.Println(r)
	f.Println(cap(r))

	// SLICE CON MAPAS ADENTRO
	var e []interface{}

	map1 := map[string]int{"clave1": 10, "clave2": 20}
	e = append(e, map1)
	map2 := map[string]string{"nombre": "Juan", "apellido": "Pérez"}
	e = append(e, map2)
	f.Println(e)
	f.Println(cap(e))

	//RECORRER SLICE CON MAPA DENTRO
	for _, v := range e {
		f.Println(v)
	}
}
