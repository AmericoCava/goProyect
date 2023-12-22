package main

import f "fmt"

func main() {
	if true {
		f.Println(001)
	}
	if false {
		f.Println(002)
	}
	if !true {
		f.Println(003)
	}
	if !false {
		f.Println(004)
	}
	if 42 == 42 {
		f.Println(005)
	}
}
