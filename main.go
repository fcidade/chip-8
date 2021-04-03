package main

import (
	"github.com/franciscocid/chip-8/chip8"
	"github.com/franciscocid/chip-8/graphics"
)

func main() {
	c8 := chip8.New()
	g := graphics.New(c8.Tick)
	c8.UI = g

	err := g.Run()
	if err != nil {
		panic(err)
	}
}
