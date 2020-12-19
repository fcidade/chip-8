package chip8

import (
	"fmt"
)

type AsciiGraphics struct {
	width        int
	height       int
	screenData   []bool
	enabledChar  string
	disabledChar string
}

func NewAsciiGraphics(width int, height int, enabledChar string, disabledChar string) AsciiGraphics {
	return AsciiGraphics{
		width:        width,
		height:       height,
		screenData:   make([]bool, width*height),
		enabledChar:  enabledChar,
		disabledChar: disabledChar,
	}
}

func (g *AsciiGraphics) Draw() {
	for y := int(0); y < g.height; y++ {
		for x := int(0); x < g.width; x++ {
			isEnabled := g.screenData[(y*g.width)+x]
			if isEnabled {
				fmt.Print(g.enabledChar)
			} else {
				fmt.Print(g.disabledChar)
			}
		}
		fmt.Println()
	}
}

func (g *AsciiGraphics) PutPixel(x, y int, active bool) {
	g.screenData[(g.width*y)+x] = active
}
