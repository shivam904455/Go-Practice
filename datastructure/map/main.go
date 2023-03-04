package main

import "fmt"

func main() {
	o:= make(map[int]int)
	o[23] = 10
	o[12] = 11
	m := make(map[string]int)
	m["vishal"] = 10
	m["pqr"] = 12

	val, ok := m["pqr"]
	fmt.Println("val =", val, "ok =", ok)
	val, ok = m["vishal"]
	if ok {
		fmt.Println("val =", val, "ok =", ok)
	} else {
		fmt.Println("value not found")
	}
}
