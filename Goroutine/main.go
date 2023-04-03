package main

import (
	"fmt"
	"time"
)

func main() {
	a()
	b()
	c()
	d()

	go a()
	go b()
	go c()
	go d()
	time.Sleep(time.Second * 2)
}
func a() {
	for i := 1; i <= 2; i++ {
		fmt.Println("func a run hua hai I= ", i)
	}
}
func b() {
	for i := 1; i <= 2; i++ {
		fmt.Println("func b run hua hai I= ", i)
	}
}
func c() {
	for i := 1; i <= 2; i++ {
		fmt.Println("func c run hua hai I= ", i)
	}
}
func d() {
	for i := 1; i <= 2; i++ {
		fmt.Println("func d run hua hai I= ", i)
	}
}
