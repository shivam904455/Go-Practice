package main

import "fmt"

func main() {
	f := 5
	g := 6
	sum := add(f, g)
	fmt.Println("add=", sum)

	fmt.Println("sub=", sub(f, g))

	fmt.Println("mul=", mul(f, g))
	// faling case for devide
	divide, err := divid(f, 0)

	if err == nil {

		fmt.Println("divide =", divide)

	} else {
		fmt.Println("divide hi nhi hua = ", err)
	}
	// passing case for devide
	divide, err = divid(f, 2)

	if err == nil {

		fmt.Println("divid=", divide)

	} else {
		fmt.Println("divid hi nhi hua = ", divid)
	}
	fmt.Println("f =", f, "g =", g)
}

// for adding two number
func add(a, b int) int {
	c := a + b
	return c
}

// for subtraction two number
func sub(a, b int) int {
	return a - b
}

// for multiply two number
func mul(a, b int) int {
	return a * b
}

// for devid two number
func divid(d, e int) (int, error) {
	if e != 0 {
		return d / e, nil
	} else {
		return 0, fmt.Errorf("deviding with 0 not possible")
	}
}
