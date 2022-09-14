package main

import "fmt"

func main() {
	num := 3
	makeEven := makeEvenGenerator()
	fmt.Println(makeEven(num))

	n := factorial(uint(num))
	fmt.Println(n)
}

func makeEvenGenerator() func(n int) int {
	return func(i int) int {
		m := i % 2

		if m < 0 {
			i = i - 1
		} else if m > 0 {
			i = i + 1
		}

		return i
	}
}

func factorial(x uint) uint {
	if x == 0 {
		return 1
	}

	return x * factorial(x-1)
}
