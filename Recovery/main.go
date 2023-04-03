package main

import (
	"fmt"
	"os"
)

func main() {
	// defer recoverpanic()
	dbconnection := true

	if !dbconnection {
		fatalError("database not connected")
	}
	a := []string{
		"a wala record save hua hai",
	}
	savetodb(a)

	b := []string{}

	savetodb(b)
	c := []string{
		"c wala record save hua hai",
	}
	savetodb(c)
}
func recoverpanic(){
	r:=recover()
	if r != nil{
		fmt.Println("kucch to jhol hua hai",r)
	}
}
func savetodb(record []string) {
	// defer recoverpanic()
	if len(record) < 1 {
		panicError("no record found to save ")
	} else {
		fmt.Println("record =", record)
	}
}
func fatalError(anything interface{}) {
	fmt.Printf("fatal error occourd message =%v", anything)
	os.Exit(1)
}
func panicError(anything interface{}) {
	defer recoverpanic()
	fmt.Println("panic error occourd")
	panic(anything)
}
