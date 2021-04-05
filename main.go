package main

import (
	"io"
	"os"

	"github.com/franciscocid/chip-8/chip8"
	"github.com/franciscocid/chip-8/graphics"
)

func main() {
	c8 := chip8.New()
	g := graphics.New(c8.Tick)
	c8.UI = g

	file, err := os.Open("roms/BLINKY.ch8")
	// file, err := os.Open("roms/random_number_test.ch8")
	// file, err := os.Open("roms/test_opcode.ch8")
	// file, err := os.Open("roms/IBM Logo.ch8")
	if err != nil {
		panic(err)
	}
	romData, err := io.ReadAll(file)
	if err != nil {
		panic(err)
	}

	c8.LoadFonts()
	c8.LoadGame(romData)

	err = g.Run()
	if err != nil {
		panic(err)
	}
}
