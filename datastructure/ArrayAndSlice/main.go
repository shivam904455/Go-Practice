package main

import "fmt"

func main() {
	// integer value of array
	var list [11]int
	list[0] = 10
	list[1] = 11
	list[2] = 12
	list[3] = 13
	list[4] = 14
	list[5] = 15
	list[6] = 16
	list[7] = 17
	list[8] = 18
	list[9] = 19
	list[10] = 20
	fmt.Println("name=", list)

	// string value of array

	var name[11]string
	name[0] = "10"
	fmt.Println("name++++=",name[1])
	name[1] = "11"
	name[2] = "12"
	name[3] = "13"
	name[4] = "14"
	name[5] = "15"
	name[6] = "16"
	name[7] = "17"
	name[8] = "18"
	name[9] = "19"
	name[10] = "20"
	fmt.Println("name=", name)
	name[5]="4443"
	fmt.Println("name=",name)

// slice value of array
var students []string
students =append(students, "vishal")
students =append(students, "shivam")
fmt.Println("students = ",students [1])
}
// school admission
// array and slice