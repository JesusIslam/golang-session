package main

import "fmt"

type Barker interface {
	Bark()
}

type Dog struct {
}

func (d *Dog) Bark() {
	fmt.Println("woof!")
}

type Poodle struct {
	Dog
}

func (p *Poodle) Bark() {
	fmt.Println("bow!")
}

func main() {
	scoob := &Poodle{}

	doBark(scoob)
}

func doBark(b Barker) {
	b.Bark()
}
