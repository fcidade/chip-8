package chip8

type Graphics interface {
	Clear()
	TogglePixel(x, y int) (isAlreadyToggled bool)
}

type Chip8 struct {
	CurrState    Chip8State
	StateHistory []Chip8State
	UI           Graphics
}

func (c *Chip8) LoadGame(gameData []uint8) {
	c.StateHistory = make([]Chip8State, 0)

	c.CurrState = Chip8State{}
	for i, data := range gameData {
		c.CurrState.Memory[i] = data
	}
}

func (c *Chip8) Tick() {
	opcode := c.CurrState.Opcode()

	newState := c.ExecuteOpcode(opcode)

	c.StateHistory = append(c.StateHistory, c.CurrState)
	c.CurrState = newState
}

func (c *Chip8) ExecuteOpcode(opcode uint16) Chip8State {
	addr := opcode & 0x0FFF

	switch opcode {
	case 0x00E0:
		return c.cls()
	case 0x00EE:
		return c.ret()
	}

	switch opcode & 0xF000 {
	case 0x0:
		return c.sys(addr)
	case 0x1:
		return c.jmp(addr)
	case 0x2:
	case 0x3:
	case 0x4:
	case 0x5:
	case 0x6:
	case 0x7:
	case 0x8:
	case 0x9:
	case 0xA:
	case 0xB:
	case 0xC:
	case 0xD:
	case 0xE:
	case 0xF:
	}
	return c.CurrState
}

func New() *Chip8 {
	return &Chip8{
		CurrState:    Chip8State{},
		StateHistory: []Chip8State{},
	}
}
