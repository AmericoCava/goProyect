package main

import f "fmt"

func main() {
	et := []string{"Eduardo", "Tua", "Crossfit", "Beisbol", "Montañista"}
	f.Println(et)
	jl := []string{"Jacinto", "Lopez", "Correr", "Nadar", "Bailar"}
	f.Println(jl)

	cp := [][]string{et, jl}
	f.Println(cp)
}
