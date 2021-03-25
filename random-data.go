package main

import (
	"fmt"
	"math/rand"
)

func main() {
	for i := 10; i > 0; i-- {
		year := rand.Intn(1000) + 2000
		leap := year/400 == 0 || (year/4 == 0 && year/100 != 0)
		month := rand.Intn(12) + 1
		daysInMonth := 31
		switch month {
		case 2:
			if leap {
				daysInMonth = 29
			} else {
				daysInMonth = 28
			}
		case 4, 6, 9, 11:
			daysInMonth = 30
		}
		day := rand.Intn(daysInMonth) + 1
		fmt.Println(year, "年", month, "月", day, "日")
	}
}
