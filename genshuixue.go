package main

import (
	"fmt"
)

type Code string

type Programer interface {
	writeHelloWorld() Code
}

func (p *GoProgramer)writeHelloWorld() Code {
	return "golang"
}
func (p *PhpProgramer)writeHelloWorld() Code {
	return "phper"
}

func writeFirstProgram(p Programer) {
	fmt.Printf("%T %v\n", p, p.writeHelloWorld())
}

type GoProgramer struct {}
type PhpProgramer struct {}

func main() {
	gop := new(GoProgramer)
	phper := &PhpProgramer{}
	writeFirstProgram(gop)
	writeFirstProgram(phper)
}