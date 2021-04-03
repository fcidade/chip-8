package chip8

import "fmt"

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

func (c *Chip8) ExecuteOpcode(opcode uint16) (newState Chip8State) {
	switch opcode & 0xF000 {
	case 0x0:
		fmt.Println("SYS")
	case 0x1:
		fmt.Println("JP")
	case 0x2:
		fmt.Println("JP")
	case 0x3:
		fmt.Println("JP")
	case 0x4:
		fmt.Println("JP")
	case 0x5:
		fmt.Println("JP")
	case 0x6:
		fmt.Println("JP")
	case 0x7:
		fmt.Println("JP")
	case 0x8:
		fmt.Println("JP")
	case 0x9:
		fmt.Println("JP")
	case 0xA:
		fmt.Println("JP")
	case 0xB:
		fmt.Println("JP")
	case 0xC:
		fmt.Println("JP")
	case 0xD:
		fmt.Println("JP")
	case 0xE:
		fmt.Println("JP")
	case 0xF:
		fmt.Println("JP")
	default:
		fmt.Printf("Invalid opcode: 0x%04x\n", opcode)
	}
	return
}

func New() *Chip8 {
	return &Chip8{
		CurrState:    Chip8State{},
		StateHistory: []Chip8State{},
	}
}
