package chip8

import (
	"log"
)

const (
	adc int16 = iota
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
	a       int16
	x       int16
	y       int16
	pc      int16
	sp      int16
	running bool
	program []int16
	ram     []int16
}

/*NewMachine -> Temp */
func NewMachine() Machine {
	return Machine{
		a: 0, x: 0, y: 0, pc: 0, sp: -1,
		running: false,
		program: []int16{},
		ram:     make([]int16, 0x000F),
	}
}

func (vm *Machine) read(addr int16) int16 {
	return vm.ram[addr]
}

func (vm *Machine) write(addr int16, value int16) {
	vm.ram[addr] = value
}

func (vm *Machine) loadProgram(program []int16) {
	vm.program = program
}

func (vm *Machine) fetch() int16 {
	if int(vm.pc) < len(vm.program) {
		return vm.program[vm.pc]
	}
	return hlt
}

func (vm *Machine) execute(opcode int16) {
	switch opcode {
	case hlt:
		log.Println("Exiting...")
		vm.running = false

	case lda:
		vm.pc++
		vm.a = int16(vm.fetch())
		log.Printf("Load to register A: %d (%#x)\n", vm.a, vm.a)

	case ldx:
		vm.pc++
		vm.x = int16(vm.fetch())
		log.Printf("Load to register X: %d (%#x)\n", vm.x, vm.x)

	case ldy:
		vm.pc++
		vm.y = int16(vm.fetch())
		log.Printf("Load to register Y: %d (%#x)\n", vm.y, vm.y)

	case nop:
		log.Println("Doing nothing.")

	case sta:
		vm.pc++
		var addr int16 = int16(vm.fetch())
		vm.pc++
		addr = (vm.fetch() << 8) | addr

		vm.write(addr, vm.a)

		log.Printf("Writing register A to ram : %d (%#x) on %d (%#x)\n", vm.a, vm.a, addr, addr)

	case stx:
		vm.pc++
		var addr int16 = int16(vm.fetch())
		vm.pc++
		addr = (vm.fetch() << 8) | addr

		vm.write(addr, vm.x)

		log.Printf("Writing register X to ram : %d (%#x) on %d (%#x)\n", vm.x, vm.x, addr, addr)

	case sty:
		vm.pc++
		var addr int16 = int16(vm.fetch())
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
