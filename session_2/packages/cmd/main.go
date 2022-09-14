package main

import (
	"github.com/JesusIslam/golang-session/session_2/packages"
)

func main() {
	p := &packages.PrintingMachine{}
	p.PublicData = "Hello world!"
	printData(p)
}

func printData(p packages.Printer) {
	p.Print()
}
