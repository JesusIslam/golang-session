package main

import "fmt"

type Dog struct {
}

func (d *Dog) Bark() {
	fmt.Println("woof!")
}

type Poodle struct {
	Dog
}

func main() {
	scoob := &Poodle{}
	scoob.Bark()
}
