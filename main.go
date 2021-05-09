package main

import (
	"io"
	"math/rand"
	"os"

	"github.com/franciscocid/chip-8/chip8"
)

func main() {
	c8 := chip8.New()
	g := chip8.NewGraphicsSDL(c8)
	rand.Seed(1)

	// file, err := os.Open("roms/BLINKY.ch8")
	file, err := os.Open("roms/random_number_test.ch8")
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
