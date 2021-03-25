package main

import (
	"fmt"
	"math/rand"
)

func main() {
	n := 33
	for {
		var i = rand.Intn(100) + 1
		if n > i {
			fmt.Println(i, "too small")
		} else if n < i {
			fmt.Println(i, "too big")
		} else {
			fmt.Println(i, "you win")
			break
		}
	}
}
