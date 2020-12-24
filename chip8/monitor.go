package chip8

type Monitor interface {
	Clear()
	Draw()
	PutPixel(x, y uint)
}
