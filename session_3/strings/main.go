package main

import (
	"fmt"
	"strings"
)

func main() {
	// contains
	s := "A quick brown fox jumps over the lazy dog."
	fmt.Println(strings.Contains(s, "brown"))
	// fields
	splitted := strings.Fields(s)
	fmt.Println(len(splitted), splitted)
	// has prefix
	fmt.Println(strings.HasPrefix(s, "A quick"))
	// has suffix
	fmt.Println(strings.HasSuffix(s, "lazy dog."))
	// join
	fmt.Println(strings.Join(splitted, ","))
	// replace all
	fmt.Println(strings.ReplaceAll(s, "fox", "racoon"))
	// split
	fmt.Println(strings.Split(s, "jumps"))
	// to lower
	fmt.Println(strings.ToLower(s))
	// to upper
	fmt.Println(strings.ToUpper(s))
	// trim space``
	s = " " + s + " "
	fmt.Println("-"+s+"-", "-"+strings.TrimSpace(s)+"-")
}
