package main

import "fmt"

func main() {
	data := []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(average(data))
	fmt.Println(add(1, 2))
	fmt.Println(sub(1, 2))
	fmt.Println(add_gte(5, 6, 10))
	fmt.Println(variadic_average(1, 2, 3, 4, 5, 6, 7, 8, 9, 10))
	fmt.Println(variadic_average(data...))
}

func average(nums []float64) float64 {
	total := 0.0
	for _, v := range nums {
		total += v
	}

	return total / float64(len(nums))
}

func add(n1, n2 float64) float64 {
	return n1 + n2
}

func sub(n1, n2 float64) (result float64) {
	result = n1 - n2

	return
}

func add_gte(n1, n2, gte float64) (float64, bool) {
	n := n1 + n2
	isgte := n >= gte

	return n, isgte
}

func variadic_average(nums ...float64) float64 {
	total := 0.0
	for _, v := range nums {
		total += v
	}

	return total / float64(len(nums))
}
