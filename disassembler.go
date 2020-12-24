package main

import (
	"fmt"
	"strings"
)

func Disassemble(program []uint16) string {
	var res strings.Builder

	for _, code := range program {
		decoded := decode(code)
		res.WriteString(fmt.Sprintf("0x%04x\t", decoded))
		res.WriteString(generate(decoded))
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

	case 0x00EE:
		return "RET"
	}

	x := xRegister(code)
	y := yRegister(code)
	addr := address(code)
	value := value(code)

	switch code & 0xF000 {
	case 0x0000:
		return fmt.Sprintf("SYS\t0x%03x", addr)

	case 0x1000:
		return fmt.Sprintf("JMP\t0x%03x", addr)

	case 0x2000:
		return fmt.Sprintf("CALL\t0x%03x", addr)

	case 0x3000:
		return fmt.Sprintf("SE\tV%d, 0x%03x", x, addr)

	case 0x4000:
		return fmt.Sprintf("SNE\tV%d, 0x%02x", x, value)

	case 0x5000:
		return fmt.Sprintf("SNE\tV%d, V%d", x, y)

	case 0x6000:
		return fmt.Sprintf("LD\tV%d, 0x%02x", x, value)

	case 0x7000:
		return fmt.Sprintf("ADD\tV%d, 0x%02x", x, value)

	case 0x8000:
		suffix := code & 0x000F

		switch suffix {
		case 0x0:
			return fmt.Sprintf("LD\tV%d, V%d", x, y)

		case 0x1:
			return fmt.Sprintf("OR\tV%d, V%d", x, y)

		case 0x2:
			return fmt.Sprintf("AND\tV%d, V%d", x, y)

		case 0x3:
			return fmt.Sprintf("XOR\tV%d, V%d", x, y)

		case 0x4:
			return fmt.Sprintf("ADD\tV%d, V%d", x, y)

		case 0x5:
			return fmt.Sprintf("SUB\tV%d, V%d", x, y)

		case 0x6:
			return fmt.Sprintf("SHR\tV%d, V%d", x, y)

		case 0x7:
			return fmt.Sprintf("SUBN\tV%d, V%d", x, y)

		case 0xE:
			return fmt.Sprintf("SHL\tV%d, V%d", x, y)
		}

	case 0x9000:
		return fmt.Sprintf("SNE\tV%d, V%d", x, y)

	case 0xA000:
		return fmt.Sprintf("LD\tI, 0x%03x", addr)

	case 0xB000:
		return fmt.Sprintf("JMP\tV0, 0x%03x", addr)

	case 0xC000:
		return fmt.Sprintf("RND\tV%d, 0x%02x", x, value)

	case 0xD000:
		nib := nibble(code)
		return fmt.Sprintf("DRW\tV%d, V%d, 0x%x", x, y, nib)

	case 0xE000:
		suffix := code & 0x00FF
		switch suffix {
		case 0x9E:
			return fmt.Sprintf("SKP\tV%d", x)
		case 0xA1:
			return fmt.Sprintf("SKNP\tV%d", x)
		}

	case 0xF000:
		suffix := code & 0x00FF
		switch suffix {
		case 0x07:
			return fmt.Sprintf("LD\tV%d, DT", x)

		case 0x0A:
			return fmt.Sprintf("LD\tV%d, KEY", x)

		case 0x15:
			return fmt.Sprintf("LD\tDT, V%d", x)

		case 0x18:
			return fmt.Sprintf("LD\tST, V%d", x)

		case 0x1E:
			return fmt.Sprintf("ADD\tI, V%d", x)

		case 0x29:
			return fmt.Sprintf("LD\tF, V%d", x)

		case 0x33:
			return fmt.Sprintf("LD\tB, V%d", x)

		case 0x55:
			return fmt.Sprintf("LD\t[I], V%d", x)

		case 0x65:
			return fmt.Sprintf("LD\tV%d, [I]", x)
		}

	}

	return fmt.Sprintf("DAT\t0x%04x", code)
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
		0x0065, 0xe000, 0xffc3, 0x22a2, 0x33f3, 0x65f2, 0x0064, 0x29f0,
		0x55d4, 0x0574, 0x29f1, 0x55d4, 0x0574, 0x29f2, 0x55d4, 0x0af3,
		0x0212,
	}

	fmt.Println("OPCODE\tSYNTAX\t")
	fmt.Println("------------------------------")
	fmt.Println(Disassemble(program))
}
