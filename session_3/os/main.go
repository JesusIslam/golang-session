package main

import (
	"fmt"
	"os"
)

func main() {
	// Getenv
	fmt.Println(os.Getenv("PWD"))
	// Getwd
	fmt.Println(os.Getwd())
	// Hostname
	fmt.Println(os.Hostname())
	// Mkdir
	fmt.Println(os.Mkdir("./test", os.ModePerm))
	// MkdirAll
	fmt.Println(os.MkdirAll("./test/test", os.ModePerm))
	// MkdirTemp
	// Remove
	fmt.Println(os.RemoveAll("./test/test"))
	// RemoveAll
	fmt.Println(os.Remove("./test"))
	// Rename
	fmt.Println(os.Mkdir("./test", os.ModePerm))
	fmt.Println(os.Rename("./test", "./testing"))
	defer os.Remove("./testing")
	// Setenv
	os.Setenv("TESTENV", "TESTING")
	fmt.Println(os.Getenv("TESTENV"))
	// TempDir
	fmt.Println(os.TempDir())
	// File
	// Create
	f, err := os.Create("./test.txt")
	if err != nil {
		panic(err)
	}
	defer os.Remove(f.Name())
	// CreateTemp
	ftmp, err := os.CreateTemp("./", "testfile_")
	if err != nil {
		panic(err)
	}
	ftmp.Close()
	os.Remove(ftmp.Name())
	// Write
	fmt.Println(f.WriteString("test"))
	f.Close()
	// ReadFile
	data, err := os.ReadFile("./test.txt")
	fmt.Println(string(data))
	// Chmod
	fmt.Println(os.Chmod("./test.txt", 0755))
	// Chown
	fmt.Println(os.Chown("didasy", os.Getuid(), os.Getgid()))
	// Open
	of, err := os.Open("./test.txt")
	if err != nil {
		panic(err)
	}
	defer of.Close()
	// IsExist
	_, err = os.Open("./notexist")
	fmt.Println(os.IsExist(err))
	// IsNotExist
	fmt.Println(os.IsNotExist(err))
	// ReadDir
	info, err := os.ReadDir(".")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%#v", info)
	// Stat
	fileStat, err := os.Stat("./test.txt")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%#v", fileStat)

	os.Exit(0)
}
