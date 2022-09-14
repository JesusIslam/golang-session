package main

import "fmt"

type Square struct {
	width  int
	length int
	area   int
}

func main() {
	sq := Square{}
	// or you could declare it like this
	// this will make a Square struct with
	// width 10, length 20, and area 0
	//  sq := Square{10, 20, 0}

	sq.width = 10
	sq.length = 20
	calculateArea(&sq)

	fmt.Println(sq.area)

	/// you can also make a pointer to a struct
	sqp := &Square{}
	sqp.width = 10
	sqp.length = 20
	calculateArea(sqp)

	fmt.Println(sqp.area)

	sq.length = 10
	sq.width = 10
	sq.calculateArea()
	fmt.Println(sq.area)

	sqp.length = 10
	sqp.width = 10
	sqp.calculateArea()
	fmt.Println(sqp.area)
}

func calculateArea(sq *Square) {
	sq.area = sq.width * sq.length
}

func (s *Square) calculateArea() {
	s.area = s.width * s.length
}

// below will not work because it passed Square as a copy
// so even if you changed the width and length, the area
// will not change
// func (s Square) calculateArea() {
// 	s.area = s.width * s.length
// }
