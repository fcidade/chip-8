package chip8

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

var Keyboard2Chip8 = map[sdl.Keycode]uint8{
	sdl.K_1: 0x1, sdl.K_2: 0x2, sdl.K_3: 0x3, sdl.K_4: 0xC,
	sdl.K_q: 0x4, sdl.K_w: 0x5, sdl.K_e: 0x6, sdl.K_r: 0xD,
	sdl.K_a: 0x7, sdl.K_s: 0x8, sdl.K_d: 0x9, sdl.K_f: 0xE,
	sdl.K_z: 0xA, sdl.K_x: 0x0, sdl.K_c: 0xB, sdl.K_v: 0xF,
}

type SDLGraphics struct {
	Title  string
	Width  int
	Height int

	running  bool
	renderer *sdl.Renderer

	c8 *Chip8

	font *ttf.Font
}

func (g *SDLGraphics) Run() error {

	if err := g.setup(); err != nil {
		return err
	}

	const FPS = 500.0
	const secsPerUpdate = 1 / FPS
	var current, elapsed, lag float64
	previous := float64(sdl.GetTicks()) * 0.001

	for g.running {
		current = float64(sdl.GetTicks()) * 0.001
		elapsed = current - previous
		previous = current

		if elapsed > 1.0 {
			continue
		}

		lag += elapsed

		// Input/Events
		g.handleEvents()

		// Update
		for lag >= secsPerUpdate {
			g.c8.Tick(secsPerUpdate)
			lag -= secsPerUpdate
		}

		// Draw
		g.renderer.SetDrawColor(0, 0, 0, 0)
		g.renderer.Clear()

		pivotX, pivotY, pivotW, pivotH := 50, 50, ScreenWidth*4, ScreenHeight*4
		borderSize := 10

		g.drawBackground(pivotX, pivotY, pivotW, pivotH, borderSize)
		g.drawChip8(pivotX, pivotY, pivotW, pivotH)

		err := g.text("teste", 100, 100)
		if err != nil {
			return err
		}

		g.renderer.Present()
	}

	return nil
}

func (g *SDLGraphics) text(msg string, x, y int) error {
	t, err := g.font.RenderUTF8Blended(msg, sdl.Color{
		R: 255,
		G: 255,
		B: 255,
		A: 255,
	})
	if err != nil {
		return err
	}

	texture, err := g.renderer.CreateTextureFromSurface(t)
	if err != nil {
		return err
	}

	err = g.renderer.Copy(texture, nil,

		&sdl.Rect{
			X: int32(x),
			Y: int32(y),
			W: 50,
			H: 50,
		})
	if err != nil {
		return err
	}

	return nil
}

func (g *SDLGraphics) selectMainPalette() {
	g.renderer.SetDrawColor(194, 62, 128, 255)
}

func (g *SDLGraphics) selectBackgroundPalette() {
	g.renderer.SetDrawColor(0, 0, 0, 0)
}

func (g *SDLGraphics) drawBackground(pivotX, pivotY, pivotW, pivotH, borderSize int) {
	backgroundRect := &sdl.Rect{
		X: int32(pivotX - borderSize),
		Y: int32(pivotY - borderSize),
		W: int32(pivotW + borderSize*2),
		H: int32(pivotH + borderSize*2),
	}
	g.selectMainPalette()
	g.renderer.FillRect(backgroundRect)
	foregroundRect := &sdl.Rect{
		X: int32(pivotX),
		Y: int32(pivotY),
		W: int32(pivotW),
		H: int32(pivotH),
	}
	g.selectBackgroundPalette()
	g.renderer.FillRect(foregroundRect)
}

func (g *SDLGraphics) drawChip8(pivotX, pivotY, pivotW, pivotH int) error {
	g.selectMainPalette()
	pixelWidth := int(pivotW / ScreenWidth)
	pixelHeight := int(pivotH / ScreenHeight)

	for y := 0; y < ScreenHeight; y++ {
		for x := 0; x < ScreenWidth; x++ {
			if g.c8.CurrState.GetPixel(uint8(x), uint8(y)) {
				rect := &sdl.Rect{
					X: int32(pivotX + x*pixelWidth),
					Y: int32(pivotY + y*pixelHeight),
					W: int32(pixelWidth),
					H: int32(pixelHeight),
				}
				g.renderer.FillRect(rect)
			}
		}
	}

	return nil
}

func (g *SDLGraphics) setup() error {
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		return err
	}

	if err := ttf.Init(); err != nil {
		return err
	}

	font, err := ttf.OpenFont("./assets/monogram.ttf", 12) // TODO: Make it dynamic
	if err != nil {
		return err
	}
	g.font = font

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

	// renderer.SetLogicalSize(int32(ScreenWidth), int32(ScreenHeight))
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
		Width:   600,
		Height:  400,
		running: true,
		c8:      c8,
	}
}
