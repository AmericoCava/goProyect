package main

import f "fmt"

func main() {
	m := map[string]int{
		"Batman": 32,
		"Robin":  27,
	}

	f.Println(m)
	f.Println(m["Batman"])
	f.Println(m["Eduard"])

	v, ok := m["Eduard"]
	f.Println(v, ok)

	if v, ok := m["Robin"]; ok {
		f.Println("Imprimiendo desde el ok", v, ok)
	}
}
