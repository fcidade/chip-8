package main

import (
	"github.com/franciscocid/chip-8/chip8"
	"github.com/franciscocid/chip-8/graphics"
)

func main() {
	chip8 := chip8.New()
	g := graphics.New(chip8.Tick)
	err := g.Run()
	if err != nil {
		panic(err)
	}
}
