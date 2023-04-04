package main

import (
	"fmt"
	"sync"
)

func a(shivamsingh *sync.WaitGroup) {
defer shivamsingh.Done()	
fmt.Println("a is running.....")
	
}
func main() {
vishalSingh:=sync.WaitGroup{}

// var wg sync.WaitGroup

	vishalSingh.Add(2)

go a(&vishalSingh)
go a(&vishalSingh)

fmt.Println("**********")

fmt.Println("-----------")
vishalSingh.Wait()
}