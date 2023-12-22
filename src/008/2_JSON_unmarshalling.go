package main

import (
	"encoding/json"
	"fmt"
)

type persona struct {
	Nombre   string `json:"Nombre"`
	Apellido string `json:"Apellido"`
	Edad     int    `json:"Edad"`
}

func main() {
	s := `[{"Nombre":"James","Apellido":"Bond","Edad":32},{"Nombre":"Americo","Apellido":"Cavallieri","Edad":35}]`
	bs := []byte(s)
	fmt.Println(bs)
	fmt.Printf("%T\n", s)
	fmt.Printf("%T\n", bs)

	var personas []persona

	err := json.Unmarshal(bs, &personas)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(personas)
}
