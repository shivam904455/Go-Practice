package main

import (
	"fmt"
	"os"
)

func main() {
	// a := []int{
	// 	1, 2, 3, 4,2,3,4,4,32,14,4,43,44,2,22,
	// }
	// fmt.Println(a)
	// fmt.Println((len(a)))
	// pass := "12345"
	// fmt.Println(len(pass))
	// if len(pass) < 8 {
	// 	fmt.Println(fmt.Errorf("password not match criteria , password %v", pass))
	// } else {

	// 	savetodb(pass)
	// }
	dbconnection := false
	if !dbconnection {
		fatalerror("can not connect db so closing program")
	}
	pass := "12345"
	fmt.Println(len(pass))
	if len(pass) < 8 {
	err:=fmt.Errorf("password not match criteria , password %v", pass)
		// fmt.Println(err)
		panic(err)
	} else {

		savetodb(pass)
	}
}
func fatalerror(message string) {
	fmt.Printf("some error which is causing fatal errpr ,message %v\n  ", message)
	os.Exit(1)
}

func savetodb(data string) {
	fmt.Println("record save success ,record", data)
}
