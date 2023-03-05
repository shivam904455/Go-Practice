package main

import "fmt"

func main() {

	fmt.Println("1")
	fmt.Println("2")
	fmt.Println("3")
	fmt.Println("4")
	fmt.Println("5")
	fmt.Println("6")
	fmt.Println("7")
	fmt.Println("8")
	fmt.Println("9")
	fmt.Println("10")
	// start point :ind point :change point
	var i int
	for i = -100; i <= 100; i += 1 {
		fmt.Println(i)
	}
	i = 10
	condition := true
	for condition {

		if i >= 101 {
			condition = false
		} else {
			i++
			fmt.Println("dusra wala for loop hai", i)
		}

	}
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
	for i := 0; i < len(list); i++ {
		fmt.Println(list[i])
	}
	for i, j := range list {
		fmt.Println("at", i, "value is", j)
	}
	// map declarition
	m := make(map[string]string)
	// save data in map
	m["xyz"] = "10"
	m["pqr"] = "PQR"
	m["a"] = "AAA"
	m["b"] = "BBB"
	m["c"] = "CCC"
	m["d"] = "DDD"

	for key, val := range m {
		fmt.Println("key =", key, "val =", val)
	}
}
