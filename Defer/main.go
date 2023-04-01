package main

import "fmt"

func main() {
	for i := 1; i < 11; i++ {
		defer fmt.Println(i)
	}
	defer fmt.Println("hello")
	defer fmt.Println("world")
	fmt.Println("shivam singh")
	defer fmt.Println("********************")
	fmt.Println("-----------------------")
	
}
