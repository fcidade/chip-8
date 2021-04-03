package chip8

import "fmt"

func (c *Chip8) sys(addr uint16) (newState Chip8State) {
	fmt.Printf("Syscall w/ address: 0x%04x\n", addr)
	return
}

func (c *Chip8) cls() (newState Chip8State) {
	fmt.Println("Clear screen")
	// TODO!
	return
}

func (c *Chip8) ret() (newState Chip8State) {
	fmt.Printf("Return from subroutine to 0x%04x\n", 0x00 /*FIXME*/)
	// TODO!
	return
}

func (c *Chip8) jmp(addr uint16) (newState Chip8State) {
	fmt.Printf("Jump to address: 0x%04x\n", addr)
	newState.PC = addr
	return
}
