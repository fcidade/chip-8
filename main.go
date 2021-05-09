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

	// file, err := os.Open("roms/ibm.ch8")
	// file, err := os.Open("roms/tank.ch8")
	// file, err := os.Open("roms/pong.ch8")
	// file, err := os.Open("roms/maze.ch8")
	file, err := os.Open("roms/tetris.ch8")
	// file, err := os.Open("roms/tictactoe.ch8")
	// file, err := os.Open("roms/space.ch8")
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
