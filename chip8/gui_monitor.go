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
	width      			int
	height     			int
	screenData 			[]bool

	window		 			*sdl.Window
	renderer	 			*sdl.Renderer

	lastKeyPressed 	uint8
}

func NewGuiMonitor(width, height int) *GuiMonitor {
	g := GuiMonitor{
		width:      		width,
		height:     		height,
		screenData: 		make([]bool, width*height),
		lastKeyPressed: 0xFF,
	}

	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		panic(err)
	}

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

	return &g
}

func (g *GuiMonitor) Clear() {
	g.renderer.SetDrawColor(0, 0, 0, 0)
	g.renderer.Clear()
}

func (g *GuiMonitor) PutPixel(x, y int) {
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
			switch t := event.(type) {
			case *sdl.KeyboardEvent:
				keyCode := t.Keysym.Sym

				if t.State == sdl.PRESSED {
					g.handleKeys(keyCode)
				}

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

func (g *GuiMonitor) setKey(key uint8) {
	g.lastKeyPressed = key
}

	
func (g *GuiMonitor) handleKeys(key sdl.Keycode) {
	keyboard := map[sdl.Keycode]uint8{
		sdl.K_1: 0x01, sdl.K_2: 0x02, sdl.K_3: 0x03, sdl.K_4: 0x0C,
		sdl.K_q: 0x04, sdl.K_w: 0x05, sdl.K_e: 0x05, sdl.K_r: 0x0D,
		sdl.K_a: 0x07, sdl.K_s: 0x08, sdl.K_d: 0x08, sdl.K_f: 0x0E,
		sdl.K_z: 0x0A, sdl.K_x: 0x00, sdl.K_c: 0x0B, sdl.K_v: 0x0F,
	}

	if v, ok := keyboard[key]; ok {
		g.setKey(v)
	}
}

func (g *GuiMonitor) KeyPressed() uint8 {
	aux := g.lastKeyPressed
	g.setKey(0xFF)
	return aux
}