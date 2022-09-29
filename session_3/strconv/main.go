package main

import (
	"fmt"
	"strconv"
)

func main() {
	var i int
	i, _ = strconv.Atoi("-42")
	fmt.Println(i)

	var s string
	s = strconv.Itoa(-42)
	fmt.Println(s)

	var b bool
	b, _ = strconv.ParseBool("true")
	fmt.Println(b)

	var f float64
	f, _ = strconv.ParseFloat("3.1415", 64)
	fmt.Println(f)

	var i64 int64
	i64, _ = strconv.ParseInt("-42", 10, 64)
	fmt.Println(i64)

	var u uint64
	u, _ = strconv.ParseUint("42", 10, 64)
	fmt.Println(u)

	s = strconv.FormatBool(true)
	fmt.Println(s)
	s = strconv.FormatFloat(3.1415, 'E', -1, 64)
	fmt.Println(s)
	s = strconv.FormatInt(-42, 16)
	fmt.Println(s)
	s = strconv.FormatUint(42, 16)
	fmt.Println(s)
}
