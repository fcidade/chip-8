package chip8

import "fmt"

// syscall: SYS instructions were originally called on chip-8 computers
// but we don't need them on our emulation, so they're just gonna be ignored.
func (c *Chip8) syscall(addr uint16) Chip8State {
	fmt.Printf("Syscall w/ address: 0x%04x\n", addr)
	return c.CurrState
}

func (c *Chip8) clearScreen() Chip8State {
	nextState := c.CurrState
	c.UI.Clear()
	fmt.Println("Screen cleared!")
	return nextState
}

// returnFromSubroutine: RET instruction gets the address on  the top of
// the stack and sets it as the current program counter, returning from the subroutine
func (c *Chip8) returnFromSubroutine() Chip8State {
	nextState := c.CurrState

	addressToReturn := c.CurrState.Stack[c.CurrState.SP-0x1]
	nextState.PC = addressToReturn
	nextState.SP--

	fmt.Printf("Return from subroutine to address: 0x%04x\n", addressToReturn)
	return nextState
}

// jumpToAddress: SYS instruction sets the current program counter to the
// address received
func (c *Chip8) jumpToAddress(addr uint16) Chip8State {
	fmt.Printf("Jump to address: 0x%03x\n", addr)
	nextState := c.CurrState
	nextState.PC = addr
	return nextState
}

// callSubroutine: CALL instruction adds current program counter to the stack and
// sets it to the received address
func (c *Chip8) callSubroutine(addr uint16) Chip8State {
	fmt.Printf("Call subroutine on address: 0x%04x\n", addr)
	nextState := c.CurrState

	nextState.Stack[c.CurrState.SP] = c.CurrState.PC
	nextState.SP++
	nextState.PC = addr
	return nextState
}

func (c *Chip8) skipIfVxEqualValue(x, value uint8) Chip8State {
	fmt.Printf("Comparing if V%x value (0x%x) is equal to 0x%x\n", x, c.CurrState.V[x], value)
	nextState := c.CurrState
	// TODO!
	return nextState
}

func (c *Chip8) skipIfVxNotEqualValue(x, value uint8) Chip8State {
	fmt.Printf("Comparing if V%x value (0x%x) it *NOT* equal to 0x%x\n", x, c.CurrState.V[x], value)
	nextState := c.CurrState
	// TODO!
	return nextState
}

func (c *Chip8) skipIfVxEqualVy(x, y uint8) Chip8State {
	vx := c.CurrState.V[x]
	nextState := c.CurrState
	vy := c.CurrState.V[y]
	fmt.Printf("Comparing if V%x value (0x%x) it equal to V%x value (0x%x)\n", x, vx, y, vy)
	// TODO!
	return nextState
}

func (c *Chip8) loadIntoVx(x, value uint8) Chip8State {
	nextState := c.CurrState
	// TODO!
	return nextState
}

func (c *Chip8) addToVx(x, value uint8) Chip8State {
	nextState := c.CurrState
	// TODO!
	return nextState
}

func (c *Chip8) loadVxIntoVy(x, y uint8) Chip8State {
	vx := c.CurrState.V[x]
	nextState := c.CurrState
	vy := c.CurrState.V[y]
	fmt.Printf("Comparing if V%x value (0x%x) it equal to V%x value (0x%x)\n", x, vx, y, vy)
	// TODO!
	return nextState
}

func (c *Chip8) loadBitwiseVxOrVyIntoVx(x, y uint8) Chip8State {
	vx := c.CurrState.V[x]
	nextState := c.CurrState
	vy := c.CurrState.V[y]
	fmt.Printf("Comparing if V%x value (0x%x) it equal to V%x value (0x%x)\n", x, vx, y, vy)
	// TODO!
	return nextState
}

func (c *Chip8) loadBitwiseVxAndVyIntoVx(x, y uint8) Chip8State {
	vx := c.CurrState.V[x]
	nextState := c.CurrState
	vy := c.CurrState.V[y]
	fmt.Printf("Comparing if V%x value (0x%x) it equal to V%x value (0x%x)\n", x, vx, y, vy)
	// TODO!
	return nextState
}

func (c *Chip8) loadBitwiseVxExclusiveOrVyIntoVx(x, y uint8) Chip8State {
	vx := c.CurrState.V[x]
	nextState := c.CurrState
	vy := c.CurrState.V[y]
	fmt.Printf("Comparing if V%x value (0x%x) it equal to V%x value (0x%x)\n", x, vx, y, vy)
	// TODO!
	return nextState
}

func (c *Chip8) addVyToVx(x, y uint8) Chip8State {
	vx := c.CurrState.V[x]
	nextState := c.CurrState
	vy := c.CurrState.V[y]
	fmt.Printf("Comparing if V%x value (0x%x) it equal to V%x value (0x%x)\n", x, vx, y, vy)
	// TODO!
	return nextState
}

func (c *Chip8) subtractVxByVy(x, y uint8) Chip8State {
	vx := c.CurrState.V[x]
	nextState := c.CurrState
	vy := c.CurrState.V[y]
	fmt.Printf("Comparing if V%x value (0x%x) it equal to V%x value (0x%x)\n", x, vx, y, vy)
	// TODO!
	return nextState
}

func (c *Chip8) shiftVxRight(x uint8) Chip8State {
	nextState := c.CurrState
	vx := c.CurrState.V[x]
	fmt.Printf("Comparing if V%x value (0x%x) it equal to \n", x, vx)
	// TODO!
	return nextState
}

func (c *Chip8) loadVySubtractedByVxIntoVx(x, y uint8) Chip8State {
	vx := c.CurrState.V[x]
	nextState := c.CurrState
	vy := c.CurrState.V[y]
	fmt.Printf("Comparing if V%x value (0x%x) it equal to V%x value (0x%x)\n", x, vx, y, vy)
	// TODO!
	return nextState
}

func (c *Chip8) shiftVxLeft(x uint8) Chip8State {
	nextState := c.CurrState
	vx := c.CurrState.V[x]
	fmt.Printf("Comparing if V%x value (0x%x) it equal to \n", x, vx)
	// TODO!
	return nextState
}

func (c *Chip8) skipIfVxNotEqualVy(x, y uint8) Chip8State {
	vx := c.CurrState.V[x]
	nextState := c.CurrState
	vy := c.CurrState.V[y]
	fmt.Printf("Comparing if V%x value (0x%x) it equal to V%x value (0x%x)\n", x, vx, y, vy)
	// TODO!
	return nextState
}

func (c *Chip8) loadAddressIntoI(addr uint16) Chip8State {
	fmt.Printf("Syscall w/ address: 0x%04x\n", addr)
	nextState := c.CurrState
	return nextState
}

func (c *Chip8) jumpToAddressPlusV0(addr uint16) Chip8State {
	fmt.Printf("Syscall w/ address: 0x%04x\n", addr)
	nextState := c.CurrState
	return nextState
}

func (c *Chip8) loadRandomValueBitwiseAndValueIntoVx(addr uint16) Chip8State {
	fmt.Printf("Syscall w/ address: 0x%04x\n", addr)
	nextState := c.CurrState
	return nextState
}

func (c *Chip8) drawSprite(x, y, nibble uint8) Chip8State {
	vx := c.CurrState.V[x]
	nextState := c.CurrState
	vy := c.CurrState.V[y]
	fmt.Printf("Comparing if V%x value (0x%x) it equal to V%x value (0x%x)\n", x, vx, y, vy)
	// TODO!
	return nextState
}

func (c *Chip8) skipIfVxKeyIsPressed(x, y uint8) Chip8State {
	vx := c.CurrState.V[x]
	nextState := c.CurrState
	vy := c.CurrState.V[y]
	fmt.Printf("Comparing if V%x value (0x%x) it equal to V%x value (0x%x)\n", x, vx, y, vy)
	// TODO!
	return nextState
}

func (c *Chip8) skipIfVxKeyIsNotPressed(x, y uint8) Chip8State {
	vx := c.CurrState.V[x]
	nextState := c.CurrState
	vy := c.CurrState.V[y]
	fmt.Printf("Comparing if V%x value (0x%x) it equal to V%x value (0x%x)\n", x, vx, y, vy)
	// TODO!
	return nextState
}

func (c *Chip8) loadDelayTimerIntoVx(x uint8) Chip8State {
	nextState := c.CurrState
	vx := c.CurrState.V[x]
	fmt.Printf("Comparing if V%x value (0x%x) it equal to \n", x, vx)
	// TODO!
	return nextState
}

func (c *Chip8) waitButtonPressAndLoadIntoVx(x uint8) Chip8State {
	nextState := c.CurrState
	vx := c.CurrState.V[x]
	fmt.Printf("Comparing if V%x value (0x%x) it equal to \n", x, vx)
	// TODO!
	return nextState
}

func (c *Chip8) loadVxIntoDelayTimer(x uint8) Chip8State {
	nextState := c.CurrState
	vx := c.CurrState.V[x]
	fmt.Printf("Comparing if V%x value (0x%x) it equal to \n", x, vx)
	// TODO!
	return nextState
}

func (c *Chip8) loadVxIntoSoundTimer(x uint8) Chip8State {
	nextState := c.CurrState
	vx := c.CurrState.V[x]
	fmt.Printf("Comparing if V%x value (0x%x) it equal to \n", x, vx)
	// TODO!
	return nextState
}

func (c *Chip8) addVxToI(x uint8) Chip8State {
	nextState := c.CurrState
	vx := c.CurrState.V[x]
	fmt.Printf("Comparing if V%x value (0x%x) it equal to \n", x, vx)
	// TODO!
	return nextState
}

func (c *Chip8) loadVxDigitSpriteAddressIntoI(x uint8) Chip8State {
	nextState := c.CurrState
	vx := c.CurrState.V[x]
	fmt.Printf("Comparing if V%x value (0x%x) it equal to \n", x, vx)
	// TODO!
	return nextState
}

func (c *Chip8) loadVxDigitsIntoI(x uint8) Chip8State {
	nextState := c.CurrState
	vx := c.CurrState.V[x]
	fmt.Printf("Comparing if V%x value (0x%x) it equal to \n", x, vx)
	// TODO!
	return nextState
}

func (c *Chip8) loadRangeV0ToVxIntoMemoryStartingFromI(x uint8) Chip8State {
	nextState := c.CurrState
	vx := c.CurrState.V[x]
	fmt.Printf("Comparing if V%x value (0x%x) it equal to \n", x, vx)
	// TODO!
	return nextState
}

func (c *Chip8) loadMemoryStartingFromIIntoRangeV0ToVx(x uint8) Chip8State {
	nextState := c.CurrState
	vx := c.CurrState.V[x]
	fmt.Printf("Comparing if V%x value (0x%x) it equal to \n", x, vx)
	// TODO!
	return nextState
}

func (c *Chip8) invalidOpcode() Chip8State {
	fmt.Println("Invalid opcode! Ignoring...")
	return c.CurrState
}
