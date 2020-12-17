package main

import "fmt"

type Opcode uint8

const (
	nop Opcode = iota
)

type VM struct {
	a int
	x int
	y int
	pc int
	sp int
}

func (vm *VM) fetch() Opcode {
	return nop
}

func (vm *VM) execute(opcode Opcode) {
	switch opcode {
	case nop:
		fmt.Println("Doing nothing.")
	}
}

func main() {
	fmt.Println("Starting...")
	vm := VM{a:0,x:0,y:0,pc:0,sp:-1}
	vm.execute(vm.fetch())
	fmt.Println("End.")
}