package chip8

import (
	"fmt"
)

type GuiMonitor struct {
	width      uint
	height     uint
	screenData []bool
}

func NewGuiMonitor(width, height uint) GuiMonitor {
	m := GuiMonitor{
		width:      width,
		height:     height,
		screenData: make([]bool, width*height),
	}

	return m
}

func (g *GuiMonitor) Clear() {
	fmt.Println("Clear")
}

func (g *GuiMonitor) Draw() {
	// for y := uint(0); y < g.height; y++ {
	// 	for x := uint(0); x < g.width; x++ {
	// 		isEnabled := g.screenData[(y*g.width)+x]
	// 		if isEnabled {
	// 			fmt.Print(g.enabledChar)
	// 		} else {
	// 			fmt.Print(g.disabledChar)
	// 		}
	// 	}
	// 	fmt.Println()
	// }
}

func (g *GuiMonitor) PutPixel(x, y uint) {
	fmt.Printf("PUT %d, %d\n", x, y)
	// g.screenData[(g.width*y)+x] = true
}
