package main

import (
	"fmt"
	"time"
)

// you can add direction to the channel
// <-chan will make dataCh only able to receive data
// you cannot input data to dataCh
// while chan<- will be the opposite
// use just chan if you want bidirectional channel
func work(dataCh <-chan string) {
	for {
		data := <-dataCh
		// this code will block
		// until there is a data
		// coming in from the dataCh channel

		fmt.Println(" ----- ", data, " ----- ")
	}
}

func main() {
	dataCh := make(chan string)
	// the code above will make unbuffered channel
	// it is synchronous which means it would stop
	// the code that try to write into it if there is
	// a data inside the channel.
	// dataCh := make(chan string, 10)
	// this meanwhile will make buffered channel,
	// as long as there is less than 10 data in it, it would
	// not block the code that try to write into it.
	go work(dataCh)

	fmt.Println("What do you want to print?")
	var input string
	fmt.Scanln(&input)

	dataCh <- input

	// This function returns a channel that
	// would block for a second
	<-time.After(time.Second)
}
