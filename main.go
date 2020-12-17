package main

import (
	"log"
)

const (
	adc uint16 = iota
	hlt
	lda
	ldx
	ldy
	nop
	sta
	stx
	sty
)

type Machine struct {
	a uint16
	x uint16
	y uint16
	pc uint16
	sp int16
	running bool
	program []uint16
	ram []uint16
}

/*NewMachine -> Temp */
func NewMachine() Machine {
	return Machine{
		a: 0, x: 0, y: 0, pc :0, sp: -1,
		running: false, 
		program: []uint16{},
		ram: make([]uint16, 0x000F),
	}
}

func (vm *Machine) read(addr uint16) uint16 {
	return vm.ram[addr]
}

func (vm *Machine) write(addr uint16, value uint16) {
	vm.ram[addr] = value
}

func (vm *Machine) loadProgram(program []uint16) {
	vm.program = program
}

func (vm *Machine) fetch() uint16 {
	if int(vm.pc) < len(vm.program) {
		return vm.program[vm.pc]
	}
	return hlt
}

func (vm *Machine) execute(opcode uint16) {
	switch opcode {
	case hlt:
		log.Println("Exiting...")
		vm.running = false

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

	case sta:
		vm.pc++
		var addr uint16 = uint16(vm.fetch())
		vm.pc++
		addr = (vm.fetch() << 8) | addr

		vm.write(addr, vm.a)

		log.Printf("Writing register A to ram : %d (%#x) on %d (%#x)\n", vm.a, vm.a, addr, addr)

	case stx:
		vm.pc++
		var addr uint16 = uint16(vm.fetch())
		vm.pc++
		addr = (vm.fetch() << 8) | addr

		vm.write(addr, vm.x)

		log.Printf("Writing register X to ram : %d (%#x) on %d (%#x)\n", vm.x, vm.x, addr, addr)

	case sty:
		vm.pc++
		var addr uint16 = uint16(vm.fetch())
		vm.pc++
		addr = (vm.fetch() << 8) | addr

		vm.write(addr, vm.y)

		log.Printf("Writing register Y to ram : %d (%#x) on %d (%#x)\n", vm.y, vm.y, addr, addr)

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
	program := []uint16{
		nop,
		lda, 0x10,
		ldx, 0x20,
		ldy, 0x11,
		sta, 0x01, 0x00,
		stx, 0x02, 0x00,
		sty, 0x03, 0x00,
	}

	log.Println("Starting...")

	vm := NewMachine()

	vm.loadProgram(program)
	vm.run()

	log.Println("Ram:")
	log.Println(vm.ram)
	log.Println("End.")
}