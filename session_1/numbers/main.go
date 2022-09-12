package main

import (
	"fmt"
)

func main() {
	// Any undeclared number without decimal would be created as "int"
	fmt.Println("2 + 2 = ", 2+2)
	// Except if at least one of the operated variable is declared integer
	// then the rest of the operated numbers will follow
	var i int = 3
	fmt.Println("3 + 2 = ", i+2)
	// Any undeclared number with decimal would be created as float64
	// except if the system only support 32-bit floating point, it would be float32
	fmt.Println("1.5 + 1.5 = ", 1.5+1.5)
	// If at least one of the operated variable is declared float64
	// then the rest of the operated numbers would be also float64
	var j float64 = 3.5
	fmt.Println("3.5 + 1.5 = ", j+1.5)
}
