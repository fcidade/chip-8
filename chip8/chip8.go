package chip8

import "fmt"

type Chip8 struct {
	memory  []uint
	pc      uint
	v       []uint
	program []uint
	i       *uint
}

func New() Chip8 {
	c := Chip8{
		memory: make([]uint, 0xFFF),
		v:      make([]uint, 16),
		pc:     0,
	}
	c.i = &c.v[0xF]
	return c
}

func (c *Chip8) LoadProgram(data []uint) {
	c.program = data
}

func (c Chip8) fetch() uint {
	return c.program[c.pc]
}

func (c Chip8) decode(hex uint) uint {
	return (hex & 0xFF00 >> 8) | (hex & 0x00FF << 8)
}

func (c *Chip8) execute(code uint) {

	switch code & 0xF000 {
	case 0x0000:
		if code == 0x00E0 {
			fmt.Println("Screen cleared!")
		}

	case 0x6000:
		register := c.getRegister(code)
		value := c.getValue(code)
		fmt.Printf("Set V%d to 0x%x!\n", register, value)

	case 0xC000:
		register := c.getRegister(code)
		ran := 0
		value := c.getValue(code)
		fmt.Printf("Put in V%d generated random number '%d' and bitwise & w/ '%d' \n", register, ran, value)

	case 0xA000:
		nnn := code & 0x0FFF
		fmt.Printf("Set I to: 0x%x\n", nnn)

	case 0xF000:
		// There are other cases, do in the future
		fmt.Printf("I'm being honest: i didn't understand 0x%x\n", code)

	default:
		fmt.Printf("No case for: 0x%x\n", code)
	}

	c.pc++
}

func (c Chip8) getRegister(code uint) uint {
	return code & 0x0F00 >> 8
}

func (c Chip8) getValue(code uint) uint {
	return code & 0x00FF
}

func (c Chip8) Run() {
	for range c.program {
		c.execute(c.decode(c.fetch()))
	}
}
