package main

import "fmt"

func main() {
	fmt.Println("Hello!")
	fmt.Printf("Hello, %s", "world!")

	var input string
	n, err := fmt.Scanln(&input)
	fmt.Println(n, err, input)

	onetwothree := fmt.Sprintf("%d", 123)
	fmt.Println(onetwothree)
}
