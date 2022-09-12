package main

import "fmt"

func main() {
	// This will declare an empty map of string as key and int as value
	// if you try to access this, it would return a runtime error
	var kv map[string]int

	// You can instantiate it using
	kv = map[string]int{}
	kv["one"] = 1
	fmt.Println(kv)
	// Or
	kv = make(map[string]int)
	kv["one"] = 1
	fmt.Println(kv)

	// Or you can make a map with the memory already reserved
	// this would reserve memory for 32 keys
	kv = make(map[string]int, 32)
	kv["one"] = 1
	fmt.Println(kv)

	// You can use "delete()" function to delete a key from a map
	kv["two"] = 2
	delete(kv, "one")
	fmt.Println(kv)

	// You can also check for key's existence
	one, ok := kv["one"]
	fmt.Println(one, ok)
	// one would be the zero value of an int 0 and ok would return false
}
