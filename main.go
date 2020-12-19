package main

import (
	"github.com/franciscocid/vm/chip8"
)

func main() {

	// g := chip8.NewAsciiGraphics(64, 32, "X", " ")
	program := []uint{
		0x0065,
		0xe000,
		0xffc3,
		0x22a2,
		0x33f3,
		// 0x65f2,
		// 0x0064,
		// 0x29f0,
	}

	c8 := chip8.New()
	c8.LoadProgram(program)
	c8.Run()

	// log.Println("Starting...")

	// vm := NewMachine()

	// vm.loadProgram(program)
	// vm.run()

	// log.Println("Ram:")
	// log.Println(vm.ram)
	// log.Println("End.")
}
