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

	m["Americo"] = 25
	f.Println(m)

	for k, v := range m {
		f.Println(k, v)
	}
}
