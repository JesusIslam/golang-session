package packages

import "fmt"

type Printer interface {
	Print()
}

type privateInterface interface {
}

type PrintingMachine struct {
	privateData string
	PublicData  string
	Printed     bool
}

func (p *PrintingMachine) Print() {
	fmt.Println(p.PublicData)
	p.Printed = true
}

type privateStruct struct {
}

const (
	privateConstant = 1
	PublicConstant  = 2
)

var (
	privateVar = 0
	PublicVar  = 2
)

func Print(data string) {
	privatePrint(data)
}

func privatePrint(data string) {
	fmt.Println(data)
}
