package main

import "fmt"

func main() {
	fmt.Println(len("The length of the string"))

	fmt.Println("The length of the string"[3])

	// You can do substring in two different directions and one cut
	// You should remember that the index is started at 0 though
	// so [4:] would mean getting all characters starting with the index 4
	// or the fifth characters
	fmt.Println("The length of the string"[4:])

	// While [:2] means getting all characters until index 3
	// but NOT get the index 3 characters and after
	fmt.Println("The length of the string"[:3])

	// To cut string, you can do this
	// This will print the word "length"
	fmt.Println("The length of the string"[4:10])

	// You can use append operator to join two strings
	fmt.Println("The length" + " of the string")
}
