package chip8

import (
	"fmt"
	"github.com/veandco/go-sdl2/sdl"
)

const (
	chipWidth  int = 64
	chipHeight     = 32
)

type GuiMonitor struct {
	width      int
	height     int
	screenData []bool

	window		 *sdl.Window
	renderer	 *sdl.Renderer
}

func NewGuiMonitor(width, height int) GuiMonitor {
	g := GuiMonitor{
		width:      width,
		height:     height,
		screenData: make([]bool, width*height),
	}

	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		panic(err)
	}
	// defer sdl.Quit()

	window, err := sdl.CreateWindow(
		"Test",
		sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		int32(width), int32(height),
		sdl.WINDOW_SHOWN,
	)
	if err != nil { panic(err) }

	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil { panic(err) }
	
	g.Clear()

	sdl.Delay(200)

	g.window = window
	g.renderer = renderer

	return g
}

func (g *GuiMonitor) Clear() {
	g.renderer.SetDrawColor(0, 0, 0, 0)
	g.renderer.Clear()
}

func (g *GuiMonitor) PutPixel(x, y int) {
	fmt.Printf("PUT %d, %d\n", x, y)
	g.renderer.SetDrawColor(255, 255, 255, 255)

	xScale := g.width / chipWidth
	yScale := g.height / chipHeight
	rect := sdl.Rect{
		int32(x * xScale),
		int32(y * yScale), 
		int32(xScale),
		int32(yScale),
	}
	g.renderer.FillRect(&rect)
}

func (g *GuiMonitor) Update(updateFn func()) {
	running := true
	for running {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				fmt.Println("Quit")
				running = false
				break
			}
		}

		rect := sdl.Rect{500, 500, 100, 100}
		g.renderer.FillRect(&rect)
		updateFn()

		g.renderer.Present()
	}
}