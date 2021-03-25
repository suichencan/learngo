package main

import (
	"fmt"
	"math/rand"
)

func main() {
	total := 0.0
	for total <= 20.0 {
		n := rand.Intn(3)
		switch n {
		case 0:
			total += 0.05
		case 1:
			total += 0.1
		case 2:
			total += 0.25
		}
		fmt.Printf("Now total is %05.2f\n", total)
	}
}
