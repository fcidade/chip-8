package main

import (
	"log"
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

func (vm *Machine) execute(opcode Opcode) (keepRunning bool) {
	switch opcode {
	case hlt:
		log.Println("Exiting...")
		return false

	case lda:
		vm.pc++
		vm.a = uint16(vm.fetch())
		log.Printf("Load to register A: %d (%#x)\n", vm.a, vm.a)

	case ldx:
		vm.pc++
		vm.x = uint16(vm.fetch())
		log.Printf("Load to register X: %d (%#x)\n", vm.x, vm.x)

	case ldy:
		vm.pc++
		vm.y = uint16(vm.fetch())
		log.Printf("Load to register Y: %d (%#x)\n", vm.y, vm.y)

	case nop:
		log.Println("Doing nothing.")
	}

	vm.pc++
	return true
}

func (vm *Machine) run() {
	vm.running = true
	for {
		if !vm.execute(vm.fetch()) {
			return
		}
	}
}

func main() {
	program := []Opcode{
		nop,
		lda, 0x10,
		ldx, 0x20,
		ldy, 0x11,
	}

	log.Println("Starting...")

	vm := NewMachine()

	vm.loadProgram(program)
	vm.run()

	log.Println("End.")
}