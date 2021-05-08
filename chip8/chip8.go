package chip8

import "fmt"

type Chip8 struct {
	CurrState    State
	StateHistory []State
}

const (
	ScreenWidth                = 64
	ScreenHeight               = 32
	ProgramStartAddress uint16 = 0x200
	FontsStartAddress   uint16 = 0x050
)

const (
	NibbleSize         = 4
	ByteSize           = NibbleSize * 2
	FirstFontBitMask   = 0x80
	FirstScreenBitMask = 0x1 << (ByteSize * 7)
)

func (c *Chip8) LoadGame(gameData []uint8) {
	c.StateHistory = make([]State, 0)

	c.CurrState = State{
		PC: ProgramStartAddress,
	}
	for i, data := range gameData {
		c.CurrState.Memory[int(ProgramStartAddress)+i] = data
	}
}

func (c8 *Chip8) LoadFonts() {
	fonts := []uint8{
		0xF0, 0x90, 0x90, 0x90, 0xF0, // 0
		0x20, 0x60, 0x20, 0x20, 0x70, // 1
		0xF0, 0x10, 0xF0, 0x80, 0xF0, // 2
		0xF0, 0x10, 0xF0, 0x10, 0xF0, // 3
		0x90, 0x90, 0xF0, 0x10, 0x10, // 4
		0xF0, 0x80, 0xF0, 0x10, 0xF0, // 5
		0xF0, 0x80, 0xF0, 0x90, 0xF0, // 6
		0xF0, 0x10, 0x20, 0x40, 0x40, // 7
		0xF0, 0x90, 0xF0, 0x90, 0xF0, // 8
		0xF0, 0x90, 0xF0, 0x10, 0xF0, // 9
		0xF0, 0x90, 0xF0, 0x90, 0x90, // A
		0xE0, 0x90, 0xE0, 0x90, 0xE0, // B
		0xF0, 0x80, 0x80, 0x80, 0xF0, // C
		0xE0, 0x90, 0x90, 0x90, 0xE0, // D
		0xF0, 0x80, 0xF0, 0x80, 0xF0, // E
		0xF0, 0x80, 0xF0, 0x80, 0x80, // F
	}
	for addr, value := range fonts {
		c8.CurrState.Memory[int(FontsStartAddress)+addr] = value
	}
}

func (c *Chip8) Tick() *Chip8 {
	fmt.Printf("PC %03x\t", c.CurrState.PC)
	opcode := c.CurrState.Opcode()
	c.CurrState.PC += 2

	newState := c.ExecuteOpcode(opcode)

	if newState.DelayTimer > 0 {
		newState.DelayTimer--
	}
	if newState.SoundTimer > 0 {
		newState.SoundTimer--
	}

	c.StateHistory = append(c.StateHistory, c.CurrState)
	c.CurrState = newState
	return c
}

func (c *Chip8) ExecuteOpcode(opcode uint16) State {
	fmt.Printf("OP %04x\t", opcode)

	switch opcode {
	case 0x00E0:
		return c.clearScreen()
	case 0x00EE:
		return c.returnFromSubroutine()
	}

	addr := opcode & 0x0FFF
	x := uint8(opcode & 0x0F00 >> ByteSize)
	y := uint8(opcode & 0x00F0 >> NibbleSize)
	value := uint8(opcode & 0x00FF)
	nibble := uint8(opcode & 0x000F)

	firstOpcodeByte := opcode >> (NibbleSize * 3)
	switch firstOpcodeByte {
	case 0x0:
		return c.syscall(addr)
	case 0x1:
		return c.jumpToAddress(addr)
	case 0x2:
		return c.callSubroutine(addr)
	case 0x3:
		return c.skipIfVxEqualValue(x, value)
	case 0x4:
		return c.skipIfVxNotEqualValue(x, value)
	case 0x5:
		return c.skipIfVxEqualVy(x, y)
	case 0x6:
		return c.loadIntoVx(x, value)
	case 0x7:
		return c.addToVx(x, value)
	case 0x8:
		switch opcode & 0x000F {
		case 0x0:
			return c.loadVxIntoVy(x, y)
		case 0x1:
			return c.loadBitwiseVxOrVyIntoVx(x, y)
		case 0x2:
			return c.loadBitwiseVxAndVyIntoVx(x, y)
		case 0x3:
			return c.loadBitwiseVxExclusiveOrVyIntoVx(x, y)
		case 0x4:
			return c.addVyToVx(x, y)
		case 0x5:
			return c.subtractVxByVy(x, y)
		case 0x6:
			return c.shiftVxRight(x)
		case 0x7:
			return c.loadVySubtractedByVxIntoVx(x, y)
		case 0xE:
			return c.shiftVxLeft(x)
		}
	case 0x9:
		return c.skipIfVxNotEqualVy(x, y)
	case 0xA:
		return c.loadAddressIntoI(addr)
	case 0xB:
		return c.jumpToAddressPlusV0(addr)
	case 0xC:
		return c.loadRandomValueBitwiseAndValueIntoVx(x, value)
	case 0xD:
		return c.drawSprite(x, y, nibble)
	case 0xE:
		switch opcode & 0x00FF {
		case 0x9E:
			return c.skipIfVxKeyIsPressed(x, y)
		case 0xA1:
			return c.skipIfVxKeyIsNotPressed(x, y)
		}
	case 0xF:
		switch opcode & 0x00FF {
		case 0x07:
			return c.loadDelayTimerIntoVx(x)
		case 0x0A:
			return c.waitButtonPressAndLoadIntoVx(x)
		case 0x15:
			return c.loadVxIntoDelayTimer(x)
		case 0x18:
			return c.loadVxIntoSoundTimer(x)
		case 0x1E:
			return c.addVxToI(x)
		case 0x29:
			return c.loadVxDigitSpriteAddressIntoI(x)
		case 0x33:
			return c.loadVxDigitsIntoI(x)
		case 0x55:
			return c.loadRangeV0ToVxIntoMemoryStartingFromI(x)
		case 0x65:
			return c.loadMemoryStartingFromIIntoRangeV0ToVx(x)
		}
	}
	return c.CurrState
}

func New() *Chip8 {
	return &Chip8{
		CurrState:    State{},
		StateHistory: []State{},
	}
}
