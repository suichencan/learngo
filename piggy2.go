package main

import (
	"fmt"
	"math/rand"
)

func main() {
	total := 0
	for total <= 2000 {
		n := rand.Intn(3)
		switch n {
		case 0:
			total += 5
		case 1:
			total += 10
		case 2:
			total += 25
		}
		dollar := total / 100
		cent := total % 100
		fmt.Printf("Now totil is %d.%02d\n", dollar, cent)
	}
}
