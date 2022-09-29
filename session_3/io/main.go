package main

import (
	"bytes"
	"fmt"
	"io"
)

func main() {
	// bytes.Buffer satisfies:
	// io.Writer
	// io.Reader
	// io.ReadWriter
	buf := &bytes.Buffer{}
	target := &bytes.Buffer{}
	// writestring
	_, err := io.WriteString(buf, "Hello!")
	if err != nil {
		panic(err)
	}
	// readall
	data, err := io.ReadAll(buf)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(data))
	// copy
	_, err = io.Copy(target, buf)
	if err != nil {
		panic(err)
	}
	fmt.Println(target.String())
}
