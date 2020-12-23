package chip8

import (
	"fmt"
)

type AsciiMonitor struct {
	width        int
	height       int
	screenData   []bool
	enabledChar  string
	disabledChar string
}

func NewAsciiMonitor(width int, height int, enabledChar string, disabledChar string) AsciiMonitor {
	return AsciiMonitor{
		width:        width,
		height:       height,
		screenData:   make([]bool, width*height),
		enabledChar:  enabledChar,
		disabledChar: disabledChar,
	}
}

func (g *AsciiMonitor) Draw() {
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

func (g *AsciiMonitor) PutPixel(x, y int, active bool) {
	g.screenData[(g.width*y)+x] = active
}
