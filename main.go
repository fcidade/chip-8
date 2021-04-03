package main

import (
	"fmt"

	"github.com/franciscocid/chip-8/chip8"
)

func main() {
	c8 := chip8.New()
	fmt.Println(c8)

	// g := graphics.New()
	// err := g.Run(func(s *graphics.SDLGraphics) {})
	// if err != nil {
	// 	panic(err)
	// }
}
