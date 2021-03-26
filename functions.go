package main

import "fmt"

func kelTocel(k float64) float64 {
	return k - 273.15
}
func celTofah(c float64) float64 {
	return (c * 9.0 / 5.0) + 32.0
}
func main() {
	fmt.Println(kelTocel(233))
	fmt.Println(celTofah(kelTocel(0)))
}
