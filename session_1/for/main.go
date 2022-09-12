package main

import "fmt"

func main() {
	// There is only for loop in Go
	// Like this regular for loop
	for i := 0; i <= 5; i++ {
		fmt.Println(i)
	}

	fmt.Println("---------------")

	// or this alternative way to code the above
	i := 0
	for i <= 5 {
		fmt.Println(i)
		i = i + 1
	}

	fmt.Println("---------------")

	// what if you want to do "do while"?
	// the code inside the loop will be executed once
	// regardless of the initial value of "i"
	i = 0
	for {
		fmt.Println(i)
		i++

		if i >= 5 {
			break
		}
	}

	// What about infinite loop?
	// The commented out code below will print "infinite" indefinitely
	// for {
	// 	fmt.Println("infinite")
	// }
}
