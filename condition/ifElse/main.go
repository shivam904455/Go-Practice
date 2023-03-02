package main

import "fmt"

func main() {
	num := 12

	if num == 10 {
		fmt.Println("hello team...")
	} else if num == 11 {
		fmt.Println("hello guys...+++")
	} else {
		fmt.Println("no a same its num", num)
	}
	event := "pta nhi kya karana hai"
	if event == "open door" {
		fmt.Println("gate kholo")
	} else if event == "close door" {
		fmt.Println("band karo gate")
	} else {
		fmt.Println("pta  nhi kya karana hai", event)
	}

	num1 := 60
	if num1 == 0 {
		fmt.Println("val 0 h")
	} else if num1 < 0 {
		fmt.Println("negative value hai....", num1)
	} else {
		fmt.Println("positive value hai...", num1)
	}
	num1 = 2
	if (num1 == 0 && num1 > 0) && (num1%2 == 0) {
		fmt.Println("postive hai or even bhi hai")
	} else if (num1 == 0 && num > 0) && (num1%2 != 0) {
		fmt.Println("positive value hai aur odd hai...", num1)
	} else if num1 < 0 {
		fmt.Println("nigative value  hai...", num1)
	} else {
		fmt.Println("positive value hai...", num)
	}
}
