package main

import "fmt"

func SchoolID() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}
func main() {
	result := SchoolID()
	fmt.Println("result", result())
	fmt.Println("result", result())
	fmt.Println("result", result())
	fmt.Println("result", result())
	fmt.Println("result", result())
	fmt.Println("result", result())
	fmt.Println("result", result())
	fmt.Println("result", result())
}
