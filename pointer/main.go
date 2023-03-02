package main

import "fmt"

type student struct {
	Name           string
	Class          int
	Rollnum        string
	studentaddress address
}
type address struct {
	Lane string
	Post string
	Dist string
	Vill string
}

func main() {
	var vishal student
	vishal.Name= "vishal singh"
	vishal.Class =11
	vishal.studentaddress.Dist= "varanasi"
	vishal.studentaddress.Post="aurai"
	vishal.studentaddress.Vill="ghambhirpur"

	shivam := student{
		Name:    "shivam singh",
		Class:   13,
		Rollnum: "74847",
		studentaddress: address{
			Lane: "NH2",
			Post: "Aurai",
			Dist: "Bhadohi",
			Vill: "sawara",
		},
	}

	val := 12
	val2 := "3373"
	var vishaladdress*student
	vishaladdress =&vishal

	shivamaddress :=&shivam
	fmt.Println("vishal pointer address=",*vishaladdress,"shivam pointer address=",*shivamaddress)


	var interfaceExample interface{}

	interfaceExample = val
	fmt.Println("interface value=", interfaceExample)

	interfaceExample = val2
	fmt.Println("interface value =", interfaceExample)


	fmt.Println("hello team.....collage student's sturect", vishal)
	fmt.Println("hello team.....collage student's sturect", shivam)
}
