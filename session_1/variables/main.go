package main

import "fmt"

const (
	Four int8 = 4
	Five      = 5
)

var (
	Six         = 6
	Seven int16 = 7
)

func main() {
	var one int
	one = 1

	var two int = 2

	three := 3

	fmt.Println(one, two, three)
	fmt.Println(Four, Five, Six, Seven)
}
