package main

import "fmt"

type student struct {
	Name           string
	Class          int
	Rollnumber     string
	studentAddress Address
}
type Address struct {
	Vill string
	Dist string
	Post string
}

func main() {
	Collage := student{
		Name:       "SHivam Singh",
		Class:      12,
		Rollnumber: "8309238209",
		studentAddress: Address{
			Vill: "sawara ",
			Dist: "Bhadohi",
			Post: "Aurai",
		},
	}
	fmt.Println("hello team...collage student'struct", Collage)

}
