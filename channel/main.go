package main

import (
	"fmt"
	"sync"
)

func main() {
	// channel buffring
	shiva2 := make(chan int, 5)
	go func() {
		shiva2 <- 1
		shiva2 <- 2
		shiva2 <- 3
		shiva2 <- 4
		shiva2 <- 5
		// shiva2<-6

	}()
	i := 0

	for val := range shiva2 {
		fmt.Println("buffring wala example output hai", val)
		if i == 4 {
			break
		} else {
			i++
		}
	}
// partia sync
	var shiva1 chan int = make(chan int)
	wg := sync.WaitGroup{}
	wg.Add(1)
	go sendNumber(shiva1)
	ReciveNumber(shiva1, &wg)
	wg.Wait()
	close(shiva2)
}
func sendNumber(shiva1 chan int) {
	for i := 0; i <= 10; i++ {
		shiva1 <- i
	}
	close(shiva1)
}

func ReciveNumber(shiva2 <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for val := range shiva2 {
		fmt.Println(val)
		
	}

}
