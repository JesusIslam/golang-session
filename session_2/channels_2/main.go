package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		for {
			c1 <- "from 1"
			time.Sleep(time.Second * 2)
		}
	}()

	go func() {
		for {
			c2 <- "from 2"
			time.Sleep(time.Second * 3)
		}
	}()

	go func() {
		for {
			// the statement will block until any
			// of the channels outputted a data
			select {
			case msg1 := <-c1:
				fmt.Println(msg1)
			case msg2 := <-c2:
				fmt.Println(msg2)
				// there is a default as well
				// so the select here would not block
				// if that's what you need
				// default:
				// 	fmt.Println("DEFAULT")
			}
		}
	}()

	fmt.Println("Press any key to exit")
	var input string
	fmt.Scanln(&input)
}
