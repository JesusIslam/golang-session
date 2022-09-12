package main

import "fmt"

func main() {
	var x []float64

	x = make([]float64, 5)
	// could be declared like this too
	// without having to declare x first
	// x := make([]float64, 5)
	// what make([]float64, 5) does is it would create
	// a slice with underlying array with 5 size, and filled with zeroes already
	fmt.Println(x)

	x[0] = 10
	x[1] = 20
	x[2] = 30
	x[3] = 40
	x[4] = 50

	// also another way
	// x := []float64{10, 20, 30, 40, 50}

	var total float64 = 0
	for i := 0; i < len(x); i++ {
		total += x[i]
	}

	fmt.Println(total / float64(len(x)))

	// There is also another way to declare a slice
	// this would make and EMPTY underlying array with 5 size
	// so you could not just access x with the index without filling it first
	x = make([]float64, 0, 5)
	fmt.Println(x)

	// this is how you fill empty slice
	x = append(x, 10, 20, 30, 40, 50)
	fmt.Println(x)

	// And finally this is how you copy a slice, but make sure the destination slice
	// has more capacity than the source slice, otherwise the elements would be truncated
	z := make([]float64, 3)
	// also don't make z an empty slice otherwise nothing would be copied over
	copy(z, x)
	fmt.Println(z)
}
