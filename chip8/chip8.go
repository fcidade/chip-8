package chip8

import (
	//"fmt"
	"log"
	"math/rand"
)

type Chip8 struct {
	g      *GuiMonitor
	memory []uint8
	pc     uint16
	sp     uint8
	stack  []uint16
	v      []uint8
	i      uint16
	dt     uint8
	st     uint8
}

func (c8 *Chip8) getI() uint16 {
	return c8.i
}

func (c8 *Chip8) setI(value uint16) {
	c8.i = value
}

func (c8 *Chip8) getV(number uint8) uint8 {
	return c8.v[number]
}

func (c8 *Chip8) setV(number, value uint8) {
	c8.v[number] = value
}

func NewChip8(g *GuiMonitor) Chip8 {
	return Chip8{
		g:      g,
		memory: make([]uint8, 0xFFF),
		v:      make([]uint8, 16, 16),
		pc:     0x200,
		sp:     0x0,
		stack:  make([]uint16, 16),
		i:      0x000,
		dt:     0x00,
		st:     0x00,
	}
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
	}
}

func (c8 *Chip8) LoadProgram(programData []uint8) {
	for i, d := range programData {
		c8.Write(uint16(0x200+i), d)
	}
}

func (c8 *Chip8) fetch() uint16 {
	most_significant := c8.Read(c8.pc)
	c8.pc++
	less_significant := c8.Read(c8.pc)
	c8.pc++
	return (uint16(most_significant) << 8) | uint16(less_significant)
}

func (c8 *Chip8) execute() {
	code := c8.fetch()
	log.Printf("Fetch: 0x%04x\n", code)

	switch code {
	case 0x00E0:
		c8.g.Clear()
		log.Printf("CLS")
		return

	case 0x00EE:
		// TODO!!! Check if it's working

		if c8.sp > 0 {
			c8.sp--
			c8.pc = c8.stack[c8.sp]
			c8.stack = c8.stack[:c8.sp]
		}
		log.Printf("RET")
		return
	}

	x := xRegister(code)
	y := yRegister(code)
	addr := address(code)
	value := value(code)
	key := c8.g.KeyPressed()

	switch code & 0xF000 {
	case 0x0000:
		log.Printf("SYS\t0x%03x", addr)

	case 0x1000:
		c8.pc = addr
		log.Printf("JMP\t0x%03x", addr)

	case 0x2000:
		c8.stack = append(c8.stack, c8.pc)
		c8.sp++
		c8.pc = addr
		log.Printf("CALL\t0x%03x", addr)

	case 0x3000:
		if c8.getV(x) == value {
			c8.pc += 2
		}
		log.Printf("SE\tV%d, 0x%03x", x, addr)

	case 0x4000:
		if c8.getV(x) != value {
			c8.pc += 2
		}
		log.Printf("SNE\tV%d, 0x%02x", x, value)

	case 0x5000:
		if c8.getV(x) == c8.getV(y) {
			c8.pc += 2
		}
		log.Printf("SNE\tV%d, V%d", x, y)

	case 0x6000:
		c8.setV(x, value)

		log.Printf("LD\tV%d, 0x%02x", x, value)
		log.Println(c8.v)

	case 0x7000:
		c8.setV(x, c8.getV(x) + value)

		log.Printf("ADD\tV%d, 0x%02x", x, value)

	case 0x8000:
		suffix := code & 0x000F

		switch suffix {
		case 0x0:
			c8.setV(x, c8.getV(y))
			log.Printf("LD\tV%d, V%d", x, y)

		case 0x1:
			c8.setV(x, c8.getV(x) | c8.getV(y))
			log.Printf("OR\tV%d, V%d", x, y)

		case 0x2:
			c8.setV(x, c8.getV(x) & c8.getV(y))
			log.Printf("AND\tV%d, V%d", x, y)

		case 0x3:
			c8.setV(x, c8.getV(x) ^ c8.getV(y))
			log.Printf("XOR\tV%d, V%d", x, y)

		case 0x4:
			sum := c8.getV(x) + c8.getV(y)
			if sum > 0xFF {
				c8.setV(0xF, 1)
			} else {
				c8.setV(0xF, 0)
			}
			c8.setV(x, sum & 0xFF)
			log.Printf("ADD\tV%d, V%d", x, y)

		case 0x5:
			sub := c8.getV(x) - c8.getV(y)
			if c8.getV(x) > c8.getV(y) {
				c8.setV(0xF, 1)
			} else {
				c8.setV(0xF, 0)
			}
			c8.setV(x, sub & 0xFF)
			log.Printf("SUB\tV%d, V%d", x, y)

		case 0x6:
			c8.setV(x, c8.getV(x) >> 1)
			if c8.getV(x) & 0b1 == 1 {
				c8.setV(0xF, 1)
			} else {
				c8.setV(0xF, 0)
			}
			log.Printf("SHR\tV%d, V%d", x, y)

		case 0x7:
			sub := c8.getV(y) - c8.getV(x)
			if c8.getV(y) > c8.getV(x) {
				c8.setV(0xF, 1)
			} else {
				c8.setV(0xF, 0)
			}
			c8.setV(x, sub & 0xFF)
			log.Printf("SUBN\tV%d, V%d", x, y)

		case 0xE:
			c8.setV(x, c8.getV(x) << 1)
			if c8.getV(x) & 0x80 == 1 {
				c8.setV(0xF, 1)
			} else {
				c8.setV(0xF, 0)
			}
			log.Printf("SHR\tV%d, V%d", x, y)
			log.Printf("SHL\tV%d, V%d", x, y)
		}

	case 0x9000:
		if c8.getV(x) != c8.getV(y) {
			c8.pc += 2
		}
		log.Printf("SNE\tV%d, V%d", x, y)

	case 0xA000:
		c8.setI(addr)
		log.Printf("LD\tI, 0x%03x", addr)

	case 0xB000:
		c8.pc = addr + uint16(c8.getV(0))
		log.Printf("JMP\tV0, 0x%03x", addr)

	case 0xC000:
		c8.setV(x, random() & value)
		log.Printf("RND\tV%d, 0x%02x", x, value)

	case 0xD000:
		nib := nibble(code)
		var (
			width  uint8 = 4
			height       = nib
		)

		for i := uint8(0); i < height; i++ {
			currRow := c8.Read(c8.getI() + uint16(i))

			for j := uint8(0); j < width; j++ {
				if currRow&(0x80>>j) != 0 {
					c8.g.PutPixel(int(c8.getV(x)+j), int(c8.getV(y)+i))
				}
			}
		}

		// TODO: !!!!!!!!!!!!!!! CHECK COLLISION !!!!!!!!!!!!!!

		log.Printf("DRW\tV%d, V%d, 0x%x", x, y, nib)

	case 0xE000:
		suffix := code & 0x00FF

		switch suffix {
		case 0x9E:
			if key == c8.getV(x) {
				c8.pc += 2
			}
			log.Printf("SKP\tV%d", x)
		case 0xA1:
			if key != c8.getV(x) {
				c8.pc += 2
			}
			log.Printf("SKNP\tV%d", x)
		}

	case 0xF000:
		suffix := code & 0x00FF

		switch suffix {
		case 0x07:
			c8.setV(x, c8.dt)
			log.Printf("LD\tV%d, DT", x)

		case 0x0A:
			if key > 0x0f {
				c8.pc -= 2
			}
			c8.setV(x, key)

			log.Printf("LD\tV%d, KEY", x)

		case 0x15:
			c8.dt = c8.getV(x)
			log.Printf("LD\tDT, V%d", x)

		case 0x18:
			c8.st = c8.getV(x)
			log.Printf("LD\tST, V%d", x)

		case 0x1E:
			c8.i = c8.i + uint16(c8.getV(x))
			log.Printf("ADD\tI, V%d", x)

		case 0x29:
			spriteWidth := uint8(5)
			c8.setI(uint16(c8.getV(x) * spriteWidth))
			log.Printf("LD\tF, V%d", x)

		case 0x33:
			c8.Write(c8.getI(), c8.getV(x)/100)
			c8.Write(c8.getI()+1, c8.getV(x)%100/10)
			c8.Write(c8.getI()+2, c8.getV(x)%10)
			log.Printf("LD\tB, V%d", x)

		case 0x55:
			log.Printf("LD\t[I], V%d", x)

		case 0x65:
			for i := uint8(0); i <= x; i++ {
				c8.setV(i, c8.Read(c8.getI() + uint16(i)))
			}
			log.Printf("LD\tV%d, [I]\n", x)
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

func nibble(code uint16) uint8 {
	return uint8(code & 0x000F)
}

func random() uint8 {
	return uint8(rand.Uint32() % 0xFF)
}

func (c8 *Chip8) Tick() {
	c8.execute()
}
