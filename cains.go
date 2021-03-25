package main

import "fmt"

func main() {
	const distances = 236000000000000000
	const lightspeed = 299792
	const secondsPerDay = 86400
	const daysPerYear = 365
	gn := distances / secondsPerDay / daysPerYear / lightspeed
	fmt.Println("光年是", gn)
}
