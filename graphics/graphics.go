package graphics

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
)

type SDLGraphics struct {
	Title         string
	Width         int
	Height        int
	LogicalWidth  int
	LogicalHeight int

	running      bool
	screenPixels []bool
	renderer     *sdl.Renderer

	tickFn func()
}

func (g *SDLGraphics) Run() error {
	err := g.setup()
	if err != nil {
		return err
	}

	for g.running {

		g.handleEvents()

		g.Clear()

		g.renderer.SetDrawColor(255, 255, 255, 255)

		for i, active := range g.screenPixels {
			if active {
				g.renderer.DrawPoint(int32(i%g.LogicalWidth), int32(i%g.LogicalHeight))
			}
		}

		g.renderer.Present()

	}

	return nil
}

func (g *SDLGraphics) TogglePixel(x, y int) (isAlreadyToggled bool) {
	index := x + (y * g.LogicalWidth)
	isAlreadyToggled = g.screenPixels[index]
	g.screenPixels[index] = true
	return
}

func (g *SDLGraphics) Clear() {
	g.renderer.SetDrawColor(0, 0, 0, 255)
	g.renderer.Clear()
}

func (g *SDLGraphics) setup() error {
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		return err
	}

	window, err := sdl.CreateWindow(
		g.Title,
		sdl.WINDOWPOS_UNDEFINED,
		sdl.WINDOWPOS_UNDEFINED,
		int32(g.Width),
		int32(g.Height),
		sdl.WINDOW_SHOWN,
	)
	if err != nil {
		return err
	}

	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		return err
	}

	renderer.SetLogicalSize(int32(g.LogicalWidth), int32(g.LogicalHeight))
	g.renderer = renderer
	return nil
}

func (g *SDLGraphics) handleEvents() {
	for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
		switch event.(type) {
		case *sdl.QuitEvent:
			fmt.Println("Quit")
			g.running = false
		case *sdl.KeyboardEvent:
			g.tickFn()
		}
	}
}

func New(tickFn func()) *SDLGraphics {
	return &SDLGraphics{
		Title:         "Chip-8",
		Width:         640,
		Height:        320,
		LogicalWidth:  64,
		LogicalHeight: 32,
		running:       true,
		screenPixels:  make([]bool, 64*32),
		tickFn:        tickFn,
	}
}
