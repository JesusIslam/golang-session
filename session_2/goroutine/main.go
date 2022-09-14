package main

import (
	"fmt"
	"time"
)

func work() {
	i := 0
	for {
		fmt.Println(i)
		i++

		<-time.After(time.Second)
	}
}

func main() {
	go work()

	fmt.Println("Press any key to end")
	var input string
	fmt.Scanln(&input)
}
