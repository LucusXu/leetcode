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

func dosomething(p interface{}) {

	switch v:=p.(type) {
	case int:
		fmt.Println("int", v)
	case string:
		fmt.Println("string", v)
	default:
		fmt.Println("unkonwn type")
	}
}
func main() {
	gop := new(GoProgramer)
	phper := &PhpProgramer{}
	writeFirstProgram(gop)
	writeFirstProgram(phper)
	dosomething(10)
	dosomething("lucus")
}