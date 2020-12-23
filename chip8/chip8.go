package chip8

import "fmt"

type Chip8 struct {
	memory []uint8
	pc     uint16
	v      []uint8
	i      uint16
	dt		 uint8
	st 		 uint8
}

func New() Chip8 {
	return Chip8{
		memory: make([]uint8, 0xFFF),
		v:      make([]uint8, 16),
		pc:     0x200,
		i: 0x0000,
	}
}

func (c8 *Chip8) LoadProgram(programData []uint8) {
	for i, d := range programData {
		c8.memory[0x200 + i] = d
	}
}

func (c8 Chip8) fetch() uint {
	// return c8.memory[c8.pc]
	return 0
}

func (c8 Chip8) decode(hex uint) uint {
	return (hex & 0xFF00 >> 8) | (hex & 0x00FF << 8)
}

func (c8 *Chip8) execute(code uint) {
	fmt.Println("Placeholder! ;)")
}

func (c8 Chip8) getRegister(code uint) uint {
	return code & 0x0F00 >> 8
}

func (c8 Chip8) getValue(code uint) uint {
	return code & 0x00FF
}

func (c8 Chip8) Run() {
	for i := 0x200; i < 0xFFF; i++ {
		c8.execute(c8.decode(c8.fetch()))
	}
}
