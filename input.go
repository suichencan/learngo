package main

import "fmt"

func main() {
	str := "hello"
	var bl bool
	if str == "true" || str == "yes" || str == "1" {
		bl = true
	} else if str == "false" || str == "no" || str == "0" {
		bl = false
	} else {
		fmt.Println("input error")
	}
	fmt.Println(bl)
}
