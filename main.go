package main

import (
	"fmt"
	"os"
)

type Opcode uint16

const (
	adc Opcode = iota
	hlt
	lda
	ldx
	ldy
	nop 
)

type Machine struct {
	a uint16
	x uint16
	y uint16
	pc uint16
	sp int16
	running bool
	program []Opcode
}

/*NewMachine -> Temp */
func NewMachine() Machine {
	return Machine{
		a: 0, x: 0, y: 0, pc :0, sp: -1,
		running: false, 
		program: []Opcode{},
	}
}

func (vm *Machine) loadProgram(program []Opcode) {
	vm.program = program
}

func (vm *Machine) fetch() Opcode {
	if int(vm.pc) < len(vm.program) {
		return vm.program[vm.pc]
	}
	return hlt
}

func (vm *Machine) execute(opcode Opcode) {
	switch opcode {
	case hlt:
		os.Exit(0)

	case lda:
		vm.pc++
		vm.a = uint16(vm.fetch())
		fmt.Printf("Load to register A: %d (%#x)\n", vm.a, vm.a)

	case ldx:
		vm.pc++
		vm.x = uint16(vm.fetch())
		fmt.Printf("Load to register X: %d (%#x)\n", vm.x, vm.x)

	case ldy:
		vm.pc++
		vm.y = uint16(vm.fetch())
		fmt.Printf("Load to register Y: %d (%#x)\n", vm.y, vm.y)

	case nop:
		fmt.Println("Doing nothing.")
	}

	vm.pc++
}

func (vm *Machine) run() {
	vm.running = true
	for vm.running {
		vm.execute(vm.fetch())
	}
}

func main() {
	program := []Opcode{
		nop,
		lda, 0x10,
		ldx, 0x20,
		ldy, 0x11,
	}

	fmt.Println("Starting...")

	vm := NewMachine()

	vm.loadProgram(program)
	vm.run()

	fmt.Println("End.")
}