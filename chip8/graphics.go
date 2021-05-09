package chip8

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
)

var Keyboard2Chip8 = map[sdl.Keycode]uint8{
	sdl.K_1: 0x1, sdl.K_2: 0x1, sdl.K_3: 0x3, sdl.K_4: 0xC,
	sdl.K_q: 0x4, sdl.K_w: 0x5, sdl.K_e: 0x6, sdl.K_r: 0xD,
	sdl.K_a: 0x7, sdl.K_s: 0x8, sdl.K_d: 0x9, sdl.K_f: 0xE,
	sdl.K_z: 0xB, sdl.K_x: 0x0, sdl.K_c: 0xB, sdl.K_v: 0xF,
}

type SDLGraphics struct {
	Title  string
	Width  int
	Height int

	running  bool
	renderer *sdl.Renderer

	c8 *Chip8
}

func (g *SDLGraphics) Run() error {
	err := g.setup()
	if err != nil {
		return err
	}

	for g.running {

		g.handleEvents()

		g.c8.Tick()

		g.renderer.SetDrawColor(0, 0, 0, 0)
		g.renderer.Clear()

		g.renderer.SetDrawColor(255, 255, 255, 255)

		for y := uint8(0); y < ScreenHeight; y++ {
			for x := uint8(0); x < ScreenWidth; x++ {
				if g.c8.CurrState.GetPixel(x, y) {
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
		switch t := event.(type) {
		case *sdl.QuitEvent:
			fmt.Println("Quit")
			g.running = false
		case *sdl.KeyboardEvent:
			key := Keyboard2Chip8[t.Keysym.Sym]
			if t.Type == sdl.KEYDOWN {
				g.c8.PressKey(key)
			} else {
				g.c8.ReleaseKey(key)
			}
		}
	}
}

func NewGraphicsSDL(c8 *Chip8) *SDLGraphics {
	return &SDLGraphics{
		Title:   "Chip-8",
		Width:   ScreenWidth * 8,
		Height:  ScreenHeight * 8,
		running: true,
		c8:      c8,
	}
}
