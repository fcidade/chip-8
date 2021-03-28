package chip8

type Graphics interface {
	Clear()
	TogglePixel(x, y int) (isAlreadyToggled bool)
}
