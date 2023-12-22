package main

import (
    "fmt"
)

func numerosPares(lista []int) []int {
    var pares []int
    for _, numero := range lista {
        if numero%3 == 0 {
            pares = append(pares, numero)
        }
    }
    return pares
}

func main() {
    name := "Americo"
    fmt.Println("Hello", name)

    lista := []int{1,2,3,4}
    fmt.Println(lista)

    listaPares := numerosPares(lista)
    fmt.Println(listaPares)

    // Crear un mapa vacío
    m := make(map[string]int)

    // Asignar valores a las claves
    m["clave1"] = 42
    m["clave2"] = 13
    
    // Imprimir el mapa
    fmt.Println(m)
}

