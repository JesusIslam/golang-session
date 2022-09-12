package main

import "fmt"

func main() {
	var x [5]float64
	// could be declared like this too
	// x := [5]float64{10, 20, 30, 40, 50}

	x[0] = 10
	x[1] = 20
	x[2] = 30
	x[3] = 40
	x[4] = 50

	var total float64 = 0
	for i := 0; i < len(x); i++ {
		total += x[i]
	}

	fmt.Println(total / float64(len(x)))
}
