package chip8

import (
	"fmt"
	"os"

	"github.com/therecipe/qt/widgets"
)

type QtMonitor struct {
	width      uint
	height     uint
	screenData []bool
}

func NewQtMonitor(width, height uint) QtMonitor {
	m := QtMonitor{
		width:      width,
		height:     height,
		screenData: make([]bool, width*height),
	}

	app := widgets.NewQApplication(len(os.Args), os.Args)

	window := widgets.NewQMainWindow(nil, 0)
	window.SetWindowTitle("Hello world example")
	window.SetMinimumSize2(200, 200)

	layout := widgets.NewQVBoxLayout()

	mainWidget := widgets.NewQWidget(nil, 0)
	mainWidget.SetLayout(layout)

	input := widgets.NewQLineEdit(nil)
	input.SetPlaceholderText("1. write something")
	layout.AddWidget(input, 0, 0)

	button := widgets.NewQPushButton2("2. click me", nil)
	layout.AddWidget(button, 0, 0)

	window.SetCentralWidget(mainWidget)
	window.Show()

	go app.Exec()

	return m
}

func (g *QtMonitor) Clear() {
	fmt.Println("Clear")
}

func (g *QtMonitor) Draw() {
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

func (g *QtMonitor) PutPixel(x, y uint) {
	fmt.Printf("PUT %d, %d\n", x, y)
	// g.screenData[(g.width*y)+x] = true
}
