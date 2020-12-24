package chip8

import (
	"fmt"
)

type AsciiMonitor struct {
	width        uint
	height       uint
	screenData   []bool
	enabledChar  string
	disabledChar string
}

func NewAsciiMonitor(width, height uint, enabledChar, disabledChar string) AsciiMonitor {
	return AsciiMonitor{
		width:        width,
		height:       height,
		screenData:   make([]bool, width*height),
		enabledChar:  enabledChar,
		disabledChar: disabledChar,
	}
}

func (g *AsciiMonitor) Clear() {
	fmt.Println("---")
}

func (g *AsciiMonitor) Draw() {
	for y := uint(0); y < g.height; y++ {
		for x := uint(0); x < g.width; x++ {
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

func (g *AsciiMonitor) PutPixel(x, y uint) {
	g.screenData[(g.width*y)+x] = true
}
