package main

import (
	"fmt"

	"github.com/shivam904455/Go-Practice/importexecutable/calculater"
	"github.com/shivam904455/Go-Practice/importexecutable/store"
)

func main() {

	usr := calculater.User{}
	usr.Name = "fhfsdfhs"
	// fmt.Println("sum =", calculater.Add(54646, 5464))
	fmt.Println("Sub =", calculater.Sub(54646, 5464))
	fmt.Println("Mul =", calculater.Mul(54646, 5464))
	fmt.Println("Div =", calculater.Div(54646, 5464))

	// import using interface
	var dboprs store.DBOperations

	dboprs = store.Store{}

	dboprs.SaveRecord("just print this as a record.....")

}
