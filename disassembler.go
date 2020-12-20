package main

import (
	"fmt"
	"strings"
)

func Disassemble(program []uint16) string {
	var res strings.Builder

	for _, code := range program {
		res.WriteString(generate(decode(code)))
		res.WriteString("\n")
	}

	return res.String()
}

func decode(code uint16) uint16 {
	return (code & 0xFF00 >> 8) | (code & 0x00FF << 8)
}

func generate(code uint16) string {
	switch code {
	case 0x00E0:
		return "CLS"
	}

	switch code & 0xF000 {
	case 0x1000:
		addr := address(code)
		return fmt.Sprintf("JMP 0x%03x", addr)

	case 0x6000:
		reg := xRegister(code)
		value := value(code)
		return fmt.Sprintf("LD V%d, 0x%02x", reg, value)

	case 0x7000:
		reg := xRegister(code)
		value := value(code)
		return fmt.Sprintf("ADD V%d, 0x%02x", reg, value)

	case 0xA000:
		addr := address(code)
		return fmt.Sprintf("LD I, 0x%03x", addr)

	case 0xC000:
		reg := xRegister(code)
		value := value(code)
		return fmt.Sprintf("RND V%d, 0x%02x", reg, value)

	case 0xD000:
		x := xRegister(code)
		y := yRegister(code)
		nib := nibble(code)
		return fmt.Sprintf("DRW V%d, V%d, 0x%x", x, y, nib)

	case 0xF000:
		suffix := code & 0x00FF
		switch suffix {
		case 0x000A:
			reg := xRegister(code)
			return fmt.Sprintf("LD V%d, KEY", reg)

		case 0x0033:
			reg := xRegister(code)
			return fmt.Sprintf("LD B, V%d", reg)
			
		case 0x0065:
			reg := xRegister(code)
			return fmt.Sprintf("LD V%d, [I]", reg)
			
		case 0x0029:
			reg := xRegister(code)
			return fmt.Sprintf("LD F, V%d", reg)
		}
		
	}

	return fmt.Sprintf("NOT YET: %04x", code)
}

func xRegister(code uint16) uint16 {
	return code & 0x0F00 >> 8
}

func yRegister(code uint16) uint16 {
	return code & 0x00F0 >> 8
}

func value(code uint16) uint16 {
	return code & 0x00FF
}

func address(code uint16) uint16 {
	return code & 0x0FFF
}

func nibble(code uint16) uint16 {
	return code & 0x000F
}

func main() {

	program := []uint16{
		0xe000, 0x2aa2, 0x0c60, 0x0861, 0x1fd0, 0x0970, 0x39a2, 0x1fd0,
		0x48a2, 0x0870, 0x1fd0, 0x0470, 0x57a2, 0x1fd0, 0x0870, 0x66a2,
		0x1fd0, 0x0870, 0x75a2, 0x1fd0, 0x2812, 0x00ff, 0x00ff, 0x003c,
		0x003c, 0x003c, 0x003c, 0x00ff, 0xffff, 0xff00, 0x3800, 0x3f00,
		0x3f00, 0x3800, 0xff00, 0xff00, 0x0080, 0x00e0, 0x00e0, 0x0080,
		0x0080, 0x00e0, 0x00e0, 0xf880, 0xfc00, 0x3e00, 0x3f00, 0x3b00,
		0x3900, 0xf800, 0xf800, 0x0003, 0x0007, 0x000f, 0x00bf, 0x00fb,
		0x00f3, 0x00e3, 0xe043, 0xe000, 0x8000, 0x8000, 0x8000, 0x8000,
		0xe000, 
	}

	fmt.Println(Disassemble(program))
}
