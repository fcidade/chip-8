package chip8

import (
	"fmt"
	"log"
	"math/rand"
)

type Chip8 struct {
	g 		 AsciiMonitor
	memory []uint8
	pc     uint16
	v      []uint8
	i      uint16
	dt		 uint8
	st 		 uint8
}

func New(g AsciiMonitor) Chip8 {
	c8 := Chip8{
		g: g,
		memory: make([]uint8, 0xFFF),
		v:      make([]uint8, 16),
		pc:     0x200,
		i: 			0x0000,
	}
	c8.LoadFonts()
	return c8
}

func (c8 *Chip8) Read(addr uint16) uint8 {
	return c8.memory[addr]
}

func (c8 *Chip8) Write(addr uint16, value uint8) {
	c8.memory[addr] = value
}

func (c8 *Chip8) LoadFonts() {
	fonts := []uint8{
		0xF0, 0x90, 0x90, 0x90, 0xF0, // 0
		0x20, 0x60, 0x20, 0x20, 0x70, // 1
		0xF0, 0x10, 0xF0, 0x80, 0xF0, // 2
		0xF0, 0x10, 0xF0, 0x10, 0xF0, // 3
		0x90, 0x90, 0xF0, 0x10, 0x10, // 4
		0xF0, 0x80, 0xF0, 0x10, 0xF0, // 5
		0xF0, 0x80, 0xF0, 0x90, 0xF0, // 6
		0xF0, 0x10, 0x20, 0x40, 0x40, // 7
		0xF0, 0x90, 0xF0, 0x90, 0xF0, // 8
		0xF0, 0x90, 0xF0, 0x10, 0xF0, // 9
		0xF0, 0x90, 0xF0, 0x90, 0x90, // A
		0xE0, 0x90, 0xE0, 0x90, 0xE0, // B
		0xF0, 0x80, 0x80, 0x80, 0xF0, // C
		0xE0, 0x90, 0x90, 0x90, 0xE0, // D
		0xF0, 0x80, 0xF0, 0x80, 0xF0, // E
		0xF0, 0x80, 0xF0, 0x80, 0x80, // F
	}
	for addr, value := range fonts {
		c8.Write(uint16(addr), value)
		fmt.Println(uint16(addr), value)
	}
}

func (c8 *Chip8) LoadProgram(programData []uint8) {
	for i, d := range programData {
		c8.Write(uint16(0x200 + i), d)
	}
}

func (c8 *Chip8) fetch() uint16 {
	most_significant:= c8.Read(c8.pc)
	c8.pc++
	less_significant := c8.Read(c8.pc)
	c8.pc++
	return (uint16(most_significant) << 8) | uint16(less_significant)
}

func (c8 *Chip8) execute() {
	code := c8.fetch()
	switch code {
	case 0x00E0:
		c8.g.Clear()
		log.Printf("CLS")

	case 0x00EE:
		log.Printf("RET")
	}

	x := xRegister(code)
	y := yRegister(code)
	addr := address(code)
	value := value(code)

	switch code & 0xF000 {
	case 0x0000:
		return
		log.Printf("SYS\t0x%03x", addr)

	case 0x1000:
		log.Printf("JMP\t0x%03x", addr)

	case 0x2000:
		log.Printf("CALL\t0x%03x", addr)

	case 0x3000:
		log.Printf("SE\tV%d, 0x%03x", x, addr)

	case 0x4000:
		log.Printf("SNE\tV%d, 0x%02x", x, value)

	case 0x5000:
		log.Printf("SNE\tV%d, V%d", x, y)

	case 0x6000:
		c8.v[x] = value

		log.Printf("LD\tV%d, 0x%02x", x, value)
		log.Println(c8.v)

	case 0x7000:
		log.Printf("ADD\tV%d, 0x%02x", x, value)

	case 0x8000:
		suffix := code & 0x000F

		switch suffix {
		case 0x0:
			log.Printf("LD\tV%d, V%d", x, y)

		case 0x1:
			log.Printf("OR\tV%d, V%d", x, y)

		case 0x2:
			log.Printf("AND\tV%d, V%d", x, y)

		case 0x3:
			log.Printf("XOR\tV%d, V%d", x, y)

		case 0x4:
			log.Printf("ADD\tV%d, V%d", x, y)

		case 0x5:
			log.Printf("SUB\tV%d, V%d", x, y)

		case 0x6:
			log.Printf("SHR\tV%d, V%d", x, y)

		case 0x7:
			log.Printf("SUBN\tV%d, V%d", x, y)

		case 0xE:
			log.Printf("SHL\tV%d, V%d", x, y)
		}

	case 0x9000:
		log.Printf("SNE\tV%d, V%d", x, y)

	case 0xA000:
		c8.i = addr
		log.Printf("LD\tI, 0x%03x", addr)

	case 0xB000:
		log.Printf("JMP\tV0, 0x%03x", addr)

	case 0xC000:
		c8.v[x] = random() & value
		log.Printf("RND\tV%d, 0x%02x", x, value)

	case 0xD000:
		nib := nibble(code)
		sprite := c8.i
		for i := uint16(0); i < nib; i++ {
			if c8.Read(sprite) != 0 {
				c8.g.PutPixel(uint(c8.v[x]), uint(c8.v[x]))
			}
			sprite++
		}
		c8.g.Draw()
		log.Printf("DRW\tV%d, V%d, 0x%x", x, y, nib)

	case 0xE000:
		suffix := code & 0x00FF

		switch suffix {
		case 0x9E:
			log.Printf("SKP\tV%d", x)
		case 0xA1:
			log.Printf("SKNP\tV%d", x)
		}

	case 0xF000:
		suffix := code & 0x00FF

		switch suffix {
		case 0x07:
			log.Printf("LD\tV%d, DT", x)

		case 0x0A:
			log.Printf("LD\tV%d, KEY", x)

		case 0x15:
			log.Printf("LD\tDT, V%d", x)

		case 0x18:
			log.Printf("LD\tST, V%d", x)

		case 0x1E:
			log.Printf("ADD\tI, V%d", x)

		case 0x29:
			log.Printf("LD\tF, V%d", x)

		case 0x33:
			c8.Write(c8.i, x)
			c8.Write(c8.i + 1, x)
			c8.Write(c8.i + 2, x)
			log.Printf("LD\tB, V%d", x)

		case 0x55:
			log.Printf("LD\t[I], V%d", x)

		case 0x65:
			log.Printf("LD\tV%d, [I]", x)
		}
	}
}

func xRegister(code uint16) uint8 {
	return uint8((code & 0x0F00) >> 8)
}

func yRegister(code uint16) uint8 {
	return uint8((code & 0x00F0) >> 8)
}

func value(code uint16) uint8 {
	return uint8(code & 0x00FF)
}

func address(code uint16) uint16 {
	return code & 0x0FFF
}

func nibble(code uint16) uint16 {
	return code & 0x000F
}

func random() uint8 {
	return uint8(rand.Uint32() % 0xFF)
}

func (c8 Chip8) Run() {
	for c8.pc < 0xFFE {
		c8.execute()
	}
}
