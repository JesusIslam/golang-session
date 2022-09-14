package main

import (
	"fmt"
	"os"
)

func main() {
	f, err := os.Open("./test.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	defer catchPanic()
}

func catchPanic() {
	panicString := recover()
	fmt.Println(panicString)
}
