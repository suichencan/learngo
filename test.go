package main

import "fmt"

type i float64

func test(ii i) i {
	return ii + 1
}
func main() {
	var a i = 10
	b := test(a)
	fmt.Printf("%T%T\n", a, b)

}
