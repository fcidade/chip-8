package chip8

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
)

type Chip8Tick func() *Chip8

type SDLGraphics struct {
	Title  string
	Width  int
	Height int

	running  bool
	renderer *sdl.Renderer

	tickFn Chip8Tick
}

func (g *SDLGraphics) Run() error {
	err := g.setup()
	if err != nil {
		return err
	}

	for g.running {

		g.handleEvents()

		c8 := g.tickFn()

		g.renderer.SetDrawColor(255, 255, 255, 255)

		for y := uint8(0); y < ScreenHeight; y++ {
			for x := uint8(0); x < ScreenWidth; x++ {
				if c8.CurrState.GetPixel(x, y) {
					g.renderer.DrawPoint(int32(x), int32(y))
				}
			}
		}

		g.renderer.Present()

		sdl.Delay(100)

	}

	return nil
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

	renderer.SetLogicalSize(int32(ScreenWidth), int32(ScreenHeight))
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
			// g.tickFn()
		}
	}
}

func NewGraphicsSDL(tickFn Chip8Tick) *SDLGraphics {
	return &SDLGraphics{
		Title:   "Chip-8",
		Width:   640,
		Height:  320,
		running: true,
		tickFn:  tickFn,
	}
}
