package main

import "github.com/franciscocid/chip-8/graphics"

func main() {
	g := graphics.New()
	err := g.Run(func(s *graphics.SDLGraphics) {})
	if err != nil {
		panic(err)
	}
}
