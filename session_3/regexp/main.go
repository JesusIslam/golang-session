package main

import (
	"fmt"
	"regexp"
)

func main() {
	str := "A quick brown fox jumps over the lazy dog."
	// MatchString
	fmt.Println(regexp.MatchString("fox (jumps)", str))
	// Regexp
	// Compile, we usually use this if the pattern is provided programatically
	rgx, err := regexp.Compile("fox (jumps)")
	if err != nil {
		panic(err)
	}
	// MustCompile, we usually use this if the pattern is provided from static source like config file
	rgx = regexp.MustCompile("fox (jumps)")
	// FindString
	fmt.Println(rgx.FindString(str))
	// FindAllString
	fmt.Printf("%#v\n", rgx.FindAllString(str, -1))
	// FindStringSubmatch
	fmt.Printf("%#v\n", rgx.FindStringSubmatch(str))
	// FindAllStringSubmatch
	fmt.Printf("%#v\n", rgx.FindAllStringSubmatch(str, -1))
	// MatchString
	fmt.Println(rgx.MatchString(str))
	// ReplaceAllString
	fmt.Println(rgx.ReplaceAllString(str, "bear skips"))
}
