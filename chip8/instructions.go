package chip8

import (
	"fmt"
	"math/rand"
)

// syscall: SYS instructions were originally called on chip-8 computers
// but we don't need them on our emulation, so they're just gonna be ignored.
func (c *Chip8) syscall(addr uint16) State {
	fmt.Printf("Syscall w/ address: 0x%04x", addr)
	return c.CurrState
}

// clearScreen: CLS instruction sends a signal to clear the user interface
func (c *Chip8) clearScreen() State {
	nextState := c.CurrState
	nextState.Graphics = [ScreenHeight]uint64{}
	fmt.Println("Screen cleared!")
	return nextState
}

// returnFromSubroutine: RET instruction gets the address on  the top of
// the stack and sets it as the current program counter, returning from the subroutine
func (c *Chip8) returnFromSubroutine() State {
	nextState := c.CurrState

	addressToReturn := c.CurrState.Stack[c.CurrState.SP-0x1]
	nextState.PC = addressToReturn
	nextState.SP--

	fmt.Printf("Return from subroutine to address: 0x%04x", addressToReturn)
	return nextState
}

// jumpToAddress: SYS instruction sets the current program counter to the
// address received
func (c *Chip8) jumpToAddress(addr uint16) State {
	fmt.Printf("Jump to address: 0x%03x", addr)
	nextState := c.CurrState
	nextState.PC = addr
	return nextState
}

// callSubroutine: CALL instruction adds current program counter to the stack and
// sets it to the received address
func (c *Chip8) callSubroutine(addr uint16) State {
	fmt.Printf("Call subroutine on address: 0x%04x", addr)
	nextState := c.CurrState

	nextState.Stack[c.CurrState.SP] = c.CurrState.PC
	nextState.SP++
	nextState.PC = addr
	return nextState
}

// skipIfVxEqualValue: SE Vx, byte instruction should skip the next opcode if Vx value
// equals the value in kk
func (c *Chip8) skipIfVxEqualValue(x, value uint8) State {
	fmt.Printf("Skip next instruction if V%x value (0x%x) is equal to 0x%x", x, c.CurrState.V[x], value)
	nextState := c.CurrState

	if c.CurrState.V[x] == value {
		nextState.PC += 2
		fmt.Printf(" -> Skipped next instruction: OP %04x", nextState.Opcode())
	} else {
		fmt.Printf(" -> Continued without skip")
	}

	return nextState
}

// skipIfVxNotEqualValue: SNE Vx, byte instruction should skip the next opcode if Vx value
// is NOT equals the value in kk
func (c *Chip8) skipIfVxNotEqualValue(x, value uint8) State {
	fmt.Printf("Skip next instruction if V%x value (0x%x) is NOT equal to 0x%x", x, c.CurrState.V[x], value)
	nextState := c.CurrState

	if c.CurrState.V[x] != value {
		nextState.PC += 2
		fmt.Printf(" -> Skipped next instruction: OP %04x", nextState.Opcode())
	} else {
		fmt.Printf(" -> Continued without skip")
	}

	return nextState
}

// skipIfVxEqualVy: SE Vx, Vy instruction should skip the next opcode if Vx value
// equals the value in Vy
func (c *Chip8) skipIfVxEqualVy(x, y uint8) State {
	vx, vy := c.CurrState.V[x], c.CurrState.V[y]
	fmt.Printf("Skip next instruction if V%x value (0x%x) is equal to V%x value (0x%x)", x, vx, y, vy)

	nextState := c.CurrState

	if vx == vy {
		nextState.PC += 2
		fmt.Printf(" -> Skipped next instruction: OP %04x", nextState.Opcode())
	} else {
		fmt.Printf(" -> Continued without skip")
	}

	return nextState
}

// loadIntoVx: LD Vx, byte Instruction 6xkk should load the received value into Vx
func (c *Chip8) loadIntoVx(x, value uint8) State {
	fmt.Printf("Loading value 0x%02x into V%x", value, x)
	nextState := c.CurrState
	nextState.V[x] = value
	return nextState
}

// addToVx: ADD Vx, byte Instruction 7xkk should add the received value into Vx
func (c *Chip8) addToVx(x, value uint8) State {
	fmt.Printf("Adding value 0x%02x to V%x", value, x)
	nextState := c.CurrState
	nextState.V[x] += value
	return nextState
}

// loadIntoVx: LD Vx, Vy Instruction 8xy0 should load the Vy value into Vx
func (c *Chip8) loadVxIntoVy(x, y uint8) State {
	fmt.Printf("Loading value of V%x (0x%02x) into V%x", y, c.CurrState.V[y], x)
	nextState := c.CurrState
	nextState.V[x] = c.CurrState.V[y]
	return nextState
}

// loadBitwiseVxOrVyIntoVx: OR Vx, Vy Instruction 8xy1 should load the Vy BITWISE OR Vx value into Vx
func (c *Chip8) loadBitwiseVxOrVyIntoVx(x, y uint8) State {
	vx, vy := c.CurrState.V[x], c.CurrState.V[y]
	fmt.Printf("Loading value of V%x (0x%02x) BITWISE OR V%x (0x%02x) into V%x", x, vx, y, vy, x)
	nextState := c.CurrState

	nextState.V[x] = vx | vy
	return nextState
}

// loadBitwiseVxAndVyIntoVx: AND Vx, Vy Instruction 8xy2 should load the Vy BITWISE AND Vx value into Vx
func (c *Chip8) loadBitwiseVxAndVyIntoVx(x, y uint8) State {
	vx, vy := c.CurrState.V[x], c.CurrState.V[y]
	fmt.Printf("Loading value of V%x (0x%02x) BITWISE AND V%x (0x%02x) into V%x", x, vx, y, vy, x)
	nextState := c.CurrState

	nextState.V[x] = vx & vy
	return nextState
}

// loadBitwiseVxExclusiveOrVyIntoVx: XOR Vx, Vy Instruction 8xy3 should load the Vy BITWISE XOR Vx value into Vx
func (c *Chip8) loadBitwiseVxExclusiveOrVyIntoVx(x, y uint8) State {
	vx, vy := c.CurrState.V[x], c.CurrState.V[y]
	fmt.Printf("Loading value of V%x (0x%02x) BITWISE XOR V%x (0x%02x) into V%x", x, vx, y, vy, x)
	nextState := c.CurrState

	nextState.V[x] = vx ^ vy
	return nextState
}

// addVyToVx: Instruction 8xy4 should add the Vy value into the current Vx value
// If the sum overflows (so, it's bigger than 0xFF), set VF to 1
func (c *Chip8) addVyToVx(x, y uint8) State {
	vx, vy := c.CurrState.V[x], c.CurrState.V[y]
	fmt.Printf("Loading value of V%x (0x%02x) + V%x (0x%02x) into V%x", x, vx, y, vy, x)
	nextState := c.CurrState

	var sum uint16 = uint16(vx) + uint16(vy)
	nextState.V[x] = uint8(sum & 0x00FF)

	if sum > 0xFF {
		nextState.V[0xF] = 0x01
		fmt.Printf("\tCarry flag set to 1")
	} else {
		nextState.V[0xF] = 0x00
		fmt.Printf("\tCarry flag set to 0")
	}

	return nextState
}

// subtractVxByVy: Instruction 8xy5 should subtract the Vy value into the current Vx value
// If the sub overflows (so, it's less than 0), set VF to 1
func (c *Chip8) subtractVxByVy(x, y uint8) State {
	vx, vy := c.CurrState.V[x], c.CurrState.V[y]
	fmt.Printf("Loading value of V%x (0x%02x) - V%x (0x%02x) into V%x", x, vx, y, vy, x)
	nextState := c.CurrState

	if vx > vy {
		nextState.V[0xF] = 0x01
		fmt.Printf("\tNo Borrow flag set to 1")
	} else {
		nextState.V[0xF] = 0x00
		fmt.Printf("\tNo Borrow flag set to 0")
	}

	nextState.V[x] = vx - vy

	return nextState
}

// shiftVxRight: SHR Vx {, Vy} Instruction 8xy6 should shift right the bits on Vx and VF should be set to 1 if least significant bit is 1
func (c *Chip8) shiftVxRight(x uint8) State {
	fmt.Printf("Shifting right the value of V%x (0x%02x)", x, c.CurrState.V[x])
	nextState := c.CurrState

	nextState.V[x] = c.CurrState.V[x] >> 1
	nextState.V[0xF] = nextState.V[x] & 0x01

	return nextState
}

// loadVySubtractedByVxIntoVx: SUB Vx, Vy Instruction 8xy7 should load the Vy subtracted by Vx value into Vx
func (c *Chip8) loadVySubtractedByVxIntoVx(x, y uint8) State {
	vx, vy := c.CurrState.V[x], c.CurrState.V[y]
	fmt.Printf("Loading value of V%x (0x%02x) - V%x (0x%02x) into V%x", y, vy, x, vx, x)
	nextState := c.CurrState

	if vy > vx {
		nextState.V[0xF] = 0x01
		fmt.Printf("\tNo Borrow flag set to 1")
	} else {
		nextState.V[0xF] = 0x00
		fmt.Printf("\tNo Borrow flag set to 0")
	}

	nextState.V[x] = vy - vx

	return nextState
}

// shiftVxRight: SHL Vx {, Vy} Instruction 8xyE should shift left the bits on Vx and VF should be set to 1 if most significant bit is 1
func (c *Chip8) shiftVxLeft(x uint8) State {
	fmt.Printf("Shifting left the value of V%x (0x%02x)", x, c.CurrState.V[x])
	nextState := c.CurrState

	nextState.V[x] = c.CurrState.V[x] << 1
	nextState.V[0xF] = nextState.V[x] >> (ByteSize - 1)

	return nextState
}

// skipIfVxNotEqualVy: SNE Vx, Vy instruction should skip the next opcode if Vx value
// is NOT equals the value in Vy
func (c *Chip8) skipIfVxNotEqualVy(x, y uint8) State {
	vx, vy := c.CurrState.V[x], c.CurrState.V[y]
	fmt.Printf("Skip next instruction if V%x value (0x%x) is NOT equal to V%x value (0x%x)", x, vx, y, vy)

	nextState := c.CurrState

	if vx != vy {
		nextState.PC += 2
		fmt.Printf(" -> Skipped next instruction: OP %04x", nextState.Opcode())
	} else {
		fmt.Printf(" -> Continued without skip")
	}

	return nextState
}

// loadAddressIntoI: LD I, addr instruction Annn should load the received address into I
func (c *Chip8) loadAddressIntoI(addr uint16) State {
	fmt.Printf("Loading value 0x%03x into I", addr)
	nextState := c.CurrState
	nextState.I = addr
	return nextState
}

// jumpToAddressPlusV0: JMP V0, addr instruction Bnnn should jump the program counter to the received address + V0
func (c *Chip8) jumpToAddressPlusV0(addr uint16) State {
	sum := uint16(c.CurrState.V[0x0]) + addr
	fmt.Printf("Jump to address of V0 + %03x: 0x%03x", addr, sum)
	nextState := c.CurrState
	nextState.PC = sum
	return nextState
}

// loadRandomValueBitwiseAndValueIntoVx: RND Vx, byte instruction Cxkk should load a random value into Vx BITWISE AND received value
func (c *Chip8) loadRandomValueBitwiseAndValueIntoVx(x, value uint8) State {
	nextState := c.CurrState
	randomValue := uint8(rand.Intn(0x100)) & value
	fmt.Printf("Loading value 0x%02x into V%x", randomValue, x)
	nextState.V[x] = randomValue
	return nextState
}

// drawSprite: (DRW Vx, Vy, nibble) Instruction Dxyn draws a sprite
func (c *Chip8) drawSprite(x, y, value uint8) State {
	vx := c.CurrState.V[x]
	vy := c.CurrState.V[y]
	fmt.Printf("Drawing a sprite (0x%03x) on coords: %d, %d", c.CurrState.I, vx, vy)
	nextState := c.CurrState
	var width uint8 = 8
	var height uint8 = value

	nextState.V[0xF] = 0x00
	for row := uint8(0); row < height; row++ {
		spriteRow := c.CurrState.I + uint16(row)
		sprite := c.CurrState.Memory[spriteRow]

		for col := uint8(0); col < width; col++ {
			if sprite&(FirstFontBitMask>>col) != 0 {
				if isAlreadyPainted := nextState.SetPixel(vx+col, vy+row); isAlreadyPainted {
					nextState.V[0xF] = 0x01
				}
			}
		}
	}

	return nextState
}

// skipIfVxKeyIsPressed: (SKP Vx Key) Instruction Ex9E should skip next instruction if Vx key is pressed
func (c *Chip8) skipIfVxKeyIsPressed(x uint8) State {
	nextState := c.CurrState
	vx := c.CurrState.V[x]
	fmt.Printf("Skip next instruction if V%x (0x%02x) key is pressed", x, vx)

	if c.CurrState.Keyboard[vx] {
		nextState.PC += 2
		fmt.Printf(" -> Skipped next instruction: OP %04x", nextState.Opcode())
	} else {
		fmt.Printf(" -> Continued without skip")
	}
	return nextState
}

// skipIfVxKeyIsNotPressed: (SKNP Vx Key) Instruction ExA1 should skip next instruction if Vx key is NOT pressed
func (c *Chip8) skipIfVxKeyIsNotPressed(x uint8) State {
	nextState := c.CurrState
	vx := c.CurrState.V[x]
	fmt.Printf("Skip next instruction if V%x (0x%02x) key is released", x, vx)

	if !c.CurrState.Keyboard[vx] {
		nextState.PC += 2
		fmt.Printf(" -> Skipped next instruction: OP %04x", nextState.Opcode())
	} else {
		fmt.Printf(" -> Continued without skip")
	}
	return nextState
}

// loadDelayTimerIntoVx: LD Vx, DT Instruction Fx15 should load the Vx value into Delay Timer
func (c *Chip8) loadDelayTimerIntoVx(x uint8) State {
	fmt.Printf("Loading value 0x%02x into V%x", c.CurrState.DelayTimer, x)
	nextState := c.CurrState
	nextState.V[x] = c.CurrState.DelayTimer
	return nextState
}

// waitButtonPressAndLoadIntoVx: (LD Vx, Key) Instruction Fx0A should wait until a key is pressed and then load the key into Vx
func (c *Chip8) waitButtonPressAndLoadIntoVx(x uint8) State {
	nextState := c.CurrState
	fmt.Printf("Waiting for key press...")
	for i, key := range c.CurrState.Keyboard {
		if key {
			nextState.V[x] = uint8(i)
			return nextState
		}
	}
	nextState.PC -= 2
	return nextState
}

// loadVxIntoDelayTimer: LD DT, Vx Instruction Fx15 should load the Vx value into Delay Timer
func (c *Chip8) loadVxIntoDelayTimer(x uint8) State {
	fmt.Printf("Loading value V%x value (0x%02x) into Delay Timer", x, c.CurrState.V[x])
	nextState := c.CurrState
	nextState.DelayTimer = nextState.V[x]
	return nextState
}

// loadVxIntoSoundTimer: LD Vx, ST Instruction Fx18 should load the Vx value into Sound Timer
func (c *Chip8) loadVxIntoSoundTimer(x uint8) State {
	fmt.Printf("Loading value V%x value (0x%02x) into Sound Timer", x, c.CurrState.V[x])
	nextState := c.CurrState
	nextState.SoundTimer = nextState.V[x]
	return nextState
}

// addVxToI: ADD I, Vx Instruction Fx1E adds the value of Vx into the existing value in I
func (c *Chip8) addVxToI(x uint8) State {
	nextState := c.CurrState
	fmt.Printf("Loading value of I (0x%03x) + V%x (0x%02x) into I", c.CurrState.I, x, c.CurrState.V[x])
	nextState.I += uint16(c.CurrState.V[x])
	return nextState
}

// loadVxDigitSpriteAddressIntoI: LD F, Vx Instruction Fx29 loads the address of the Vx character sprite into I
func (c *Chip8) loadVxDigitSpriteAddressIntoI(x uint8) State {
	nextState := c.CurrState
	fmt.Printf("Loading address of the V%x character sprite (%x) into I", x, c.CurrState.V[x])
	nibble := uint16(0x0F & c.CurrState.V[x])
	nextState.I = FontsStartAddress + (5 * nibble)
	return nextState
}

// loadVxDigitsIntoI: (LD B, Vx) Instruction Fx33 should load the Vx digits into Memory at I, I+1 and I+3
func (c *Chip8) loadVxDigitsIntoI(x uint8) State {
	nextState := c.CurrState
	vx := c.CurrState.V[x]
	fmt.Printf("Loading digits of the V%x (%x) into I (0x%03x), I+1 and I+2", x, vx, c.CurrState.I)
	firstDigit := vx / 100
	secondDigit := vx / 10 % 10
	thirdDigit := vx % 10
	nextState.Memory[c.CurrState.I] = firstDigit
	nextState.Memory[c.CurrState.I+1] = secondDigit
	nextState.Memory[c.CurrState.I+2] = thirdDigit
	return nextState
}

// loadRangeV0ToVxIntoMemoryStartingFromI: (LD [I], Vx) Instruction Fx55 should loads the V[0:x] into memory starting by I
func (c *Chip8) loadRangeV0ToVxIntoMemoryStartingFromI(x uint8) State {
	nextState := c.CurrState
	fmt.Printf("Loading values from V0 to V%x starting from I (0x%03x)", x, c.CurrState.I)
	for i := uint8(0); i <= x; i++ {
		nextState.Memory[c.CurrState.I+uint16(i)] = c.CurrState.V[i]
	}
	return nextState
}

// loadMemoryStartingFromIIntoRangeV0ToVx: (LD Vx, [I]) Instruction Fx65 should loads the into V[0:x] the memory values starting by I
func (c *Chip8) loadMemoryStartingFromIIntoRangeV0ToVx(x uint8) State {
	nextState := c.CurrState
	fmt.Printf("Loading values into V0 to V%x starting from I (0x%03x)", x, c.CurrState.I)
	for i := uint8(0); i <= x; i++ {
		nextState.V[i] = c.CurrState.Memory[c.CurrState.I+uint16(i)]
	}
	return nextState
}
