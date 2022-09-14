package main

import "fmt"

func two(xPtr *int) {
	*xPtr = 2
}

func main() {
	x := 5
	two(&x)
	fmt.Println(x)

	n := new(int) // this is already a pointer
	two(n)
	fmt.Println(n) // this will print the address
	fmt.Println(*n)
}
