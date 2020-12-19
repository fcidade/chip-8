package main

import "github.com/franciscocid/vm/chip8"

// ----------------------

func main() {

	g := chip8.NewAsciiGraphics(64, 32, "X", " ")

	g.Draw()
	for i := 0; i < 10; i++ {
		g.PutPixel(i, i, true)
	}
	g.Draw()

	// program := []int16{
	// 	nop,
	// 	lda, 0x10,
	// 	ldx, 0x20,
	// 	ldy, 0x11,
	// 	sta, 0x01, 0x00,
	// 	stx, 0x02, 0x00,
	// 	sty, 0x03, 0x00,
	// }

	// log.Println("Starting...")

	// vm := NewMachine()

	// vm.loadProgram(program)
	// vm.run()

	// log.Println("Ram:")
	// log.Println(vm.ram)
	// log.Println("End.")
}
