package main

import "fmt"

type student struct {
	Name           string
	Class          int
	Rollnum        string
	studentaddress address
}
type address struct {
	Lane1 string
	Post  string
	Dist  string
	Vill  string
}

func main() {
	Shivam := student{
		Name:    "Shivam singh",
		Class:   12,
		Rollnum: "62769",
		studentaddress: address{
			Lane1: "NH2",
			Post:  "Aurai",
			Dist:  "bhadohi",
			Vill:  "Sawara ",
		},
	}
	val:=122
	val2:="22782"
	var interfaceExample interface{}

		interfaceExample = val
		fmt.Println("interface value =",interfaceExample)
		interfaceExample = val2
		fmt.Println("interface value =",interfaceExample)
		interfaceExample = false 
		fmt.Println("interface value =", interfaceExample)
	
	fmt.Println("hello team.....boys's struct", Shivam)

}


