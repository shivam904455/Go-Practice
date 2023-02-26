package main

import "fmt"

func main() {
	const number2 = 45645

	var number int
	number = -12
	number = 5645464

	var addrsOfNumber *int
	addrsOfNumber = &number

	number3:=5654
	fmt.Println(number3)

	var decnum float32
	decnum = 13211.3554151
	var addOfFloat *float32
	addOfFloat = &decnum

	var flags bool
	flags = true
	var addOfFlag *bool
	addOfFlag = &flags

	var name string
	name = "coding concept"
	var addOfName *string
	addOfName = &name

	fmt.Printf("number value = %v, decimal value =%v, flags value=%v,name value =%s\n", number, decnum, flags, name)
	fmt.Printf("address OfNumber value = %v, add of decimal value =%v, address of flag value=%v,name value =%v\n", addrsOfNumber,
		*addOfFloat, addOfFlag, addOfName)
}
