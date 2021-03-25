package main

import (
	"fmt"
	"math/rand"
	"time"
)

var name = ""

func main() {
	rand.Seed(time.Now().UnixNano())
	for i := 10; i > 0; i-- {
		n := rand.Intn(3)
		switch n {
		case 0:
			name = "Virgin Galactic"
		case 1:
			name = "SpaceX"
		case 2:
			name = "Space Adventures"
		}
		speed := rand.Intn(15) + 16
		days := 62100000 / speed / 24 / 60 / 60
		re := "单程"
		cost := speed + 20
		n = rand.Intn(2)
		if n == 1 {
			re = "往返"
			cost = cost * 2
		}

		fmt.Printf("%-20v %v   %v    %3v\n", name, days, re, cost)

	}

}
