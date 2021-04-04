package chip8

import (
	"fmt"
	"testing"

	"github.com/franciscocid/chip-8/mocks"
	"github.com/stretchr/testify/assert"
)

func TestChip8State(t *testing.T) {
	t.Run("", func(t *testing.T) {
		t.Skip()
		c := &State{}
		c.Memory[0] = 0x12
		c.Memory[1] = 0x34
		fmt.Printf("0x%04x\n", c.Opcode())
	})
}

func TestChip8(t *testing.T) {
	t.Run("(SYS addr) Instructions on range 0nnn should be ignored, as they are actually SYS calls", func(t *testing.T) {
		c := New()
		oldState := c.CurrState
		newState := c.ExecuteOpcode(0x0000)
		assert.Equal(t, oldState, newState, "State should not be altered")
	})

	t.Run("(CLS) Instruction 00E0 should clear the screen", func(t *testing.T) {
		c := New()
		uiMock := new(mocks.Graphics)
		uiMock.On("Clear").Return()
		c.UI = uiMock

		oldState := c.CurrState
		newState := c.ExecuteOpcode(0x00E0)

		assert.Equal(t, oldState, newState, "State should not be altered")
		assert.True(t, uiMock.AssertCalled(t, "Clear"), "Graphics Clear function should have been called")
		assert.True(t, uiMock.AssertNumberOfCalls(t, "Clear", 1), "Graphics Clear function should only have been called once")
	})

	t.Run("(RET) Instruction 00EE should return from a subroutine", func(t *testing.T) {
		c := New()
		c.CurrState.Stack[0x0] = 0x222
		c.CurrState.SP = 0x1
		newState := c.ExecuteOpcode(0x00EE)
		assert.Equal(t, uint16(0x222), newState.PC, "Should set Program Counter back to the value on the top of the stack")
		assert.Equal(t, uint8(0x0), newState.SP, "Stack pointer should decrement by 1")
	})

	t.Run("(JMP addr) Instruction 1nnn should set program counter to the be received address", func(t *testing.T) {
		c := New()
		newState := c.ExecuteOpcode(0x1333)
		assert.Equal(t, uint16(0x333), newState.PC, "Program Counter should be equal to the received address")
	})

	t.Run("(CALL addr) Instruction 2nnn should call subroutine on address received", func(t *testing.T) {
		c := New()
		c.CurrState.PC = 0x250
		newState := c.ExecuteOpcode(0x2325)
		assert.Equal(t, uint16(0x325), newState.PC, "Program Counter should be equal to the received address")
		assert.Equal(t, uint8(0x1), newState.SP, "Stack Pointer should be incremented by 1")
		assert.Equal(t, uint16(0x250), newState.Stack[0x0], "First element of the Stack should contain the previous Program Counter")
	})

	t.Run("(SE Vx, byte) Instruction 3xkk skips next instruction if Vx equals kk", func(t *testing.T) {
		c := New()
		c.CurrState.PC = 0x200
		c.CurrState.V[0x0] = 0xFF

		newState := c.ExecuteOpcode(0x30FF)

		assert.Equal(t, uint8(0xFF), newState.V[0x0], "Vx should have the value 0xFF")
		assert.Equal(t, uint16(0x202), newState.PC, "Program Counter should increment by 2")
	})

	t.Run("(SE Vx, byte) Instruction 3xkk should NOT skip next instruction if Vx is NOT equal kk", func(t *testing.T) {
		c := New()
		c.CurrState.PC = 0x200
		c.CurrState.V[0x0] = 0x00

		newState := c.ExecuteOpcode(0x30FF)

		assert.NotEqual(t, uint8(0xFF), newState.V[0x0], "Vx should NOT have the value 0xFF")
		assert.Equal(t, uint16(0x200), newState.PC, "Program Counter should remain the same")
	})

	t.Run("(SNE Vx, byte) Instruction 4xkk skips next instruction if Vx is NOT equals kk", func(t *testing.T) {
		c := New()
		c.CurrState.PC = 0x200
		c.CurrState.V[0x0] = 0x00

		newState := c.ExecuteOpcode(0x40FF)

		assert.NotEqual(t, uint8(0xFF), newState.V[0x0], "Vx should NOT have the value 0xFF")
		assert.Equal(t, uint16(0x202), newState.PC, "Program Counter should increment by 2")
	})

	t.Run("(SNE Vx, byte) Instruction 4xkk should NOT skip next instruction if Vx equals kk", func(t *testing.T) {
		c := New()
		c.CurrState.PC = 0x200
		c.CurrState.V[0x0] = 0xFF

		newState := c.ExecuteOpcode(0x40FF)

		assert.Equal(t, uint8(0xFF), newState.V[0x0], "Vx should have the value 0xFF")
		assert.Equal(t, uint16(0x200), newState.PC, "Program Counter should remain the same")
	})

	t.Run("(SE Vx, Vy) Instruction 5xy0 skips next instruction if Vx equals Vy", func(t *testing.T) {
		c := New()
		c.CurrState.PC = 0x200
		c.CurrState.V[0x0] = 0xFF
		c.CurrState.V[0x1] = 0xFF

		newState := c.ExecuteOpcode(0x5010)

		assert.Equal(t, uint8(0xFF), newState.V[0x0], "Vx should have the value 0xFF")
		assert.Equal(t, uint8(0xFF), newState.V[0x1], "Vy should also have the value 0xFF")
		assert.Equal(t, uint16(0x202), newState.PC, "Program Counter should increment by 2")
	})

	t.Run("(SE Vx, Vy) Instruction 5xy0 should NOT skip next instruction if Vx is NOT equal Vy", func(t *testing.T) {
		c := New()
		c.CurrState.PC = 0x200
		c.CurrState.V[0x0] = 0x00
		c.CurrState.V[0x1] = 0xFF

		newState := c.ExecuteOpcode(0x5010)

		assert.Equal(t, uint8(0x00), newState.V[0x0], "Vx should have the value 0x00")
		assert.NotEqual(t, uint8(0x00), newState.V[0x1], "Vy should NOT have the value 0x00")
		assert.Equal(t, uint16(0x200), newState.PC, "Program Counter should remain the same")
	})

	t.Run("(LD Vx, byte) Instruction 6xkk should load the kk value into Vx", func(t *testing.T) {
		c := New()
		newState := c.ExecuteOpcode(0x62EE)
		assert.Equal(t, uint8(0xEE), newState.V[0x2], "Vx should have the value 0xEE")
	})

	t.Run("(ADD Vx, byte) Instruction 7xkk should add the kk value into the current Vx value", func(t *testing.T) {
		c := New()
		c.CurrState.V[0x3] = 0x01
		newState := c.ExecuteOpcode(0x73EE)
		assert.Equal(t, uint8(0xEF), newState.V[0x3], "Vx should have the value 0xEF")
	})

	t.Run("(LD Vx, Vy) Instruction 8xy0 should load the Vy value into Vx", func(t *testing.T) {
		c := New()
		c.CurrState.V[0x2] = 0x22
		c.CurrState.V[0x1] = 0x33
		newState := c.ExecuteOpcode(0x8210)
		assert.Equal(t, uint8(0x33), newState.V[0x2], "Vx should have the value 0x33")
		assert.Equal(t, uint8(0x33), newState.V[0x1], "Vy should keep the value 0x33")
	})

	t.Run("(OR Vx, Vy) Instruction 8xy1 should load the Vy BITWISE OR Vx value into Vx", func(t *testing.T) {
		c := New()
		c.CurrState.V[0x0] = 0b00001000
		c.CurrState.V[0x1] = 0b00000100
		newState := c.ExecuteOpcode(0x8011)

		assert.Equal(t, uint8(0b00001100), newState.V[0x0], "Vx should have the value 0b00001100")
		assert.Equal(t, uint8(0b00000100), newState.V[0x1], "Vy should keep the value 0b00000100")
	})

	t.Run("(AND Vx, Vy) Instruction 8xy2 should load the Vy BITWISE AND Vx value into Vx", func(t *testing.T) {
		c := New()
		c.CurrState.V[0x0] = 0b10101000
		c.CurrState.V[0x1] = 0b01001100
		newState := c.ExecuteOpcode(0x8012)

		assert.Equal(t, uint8(0b00001000), newState.V[0x0], "Vx should have the value 0b00001000")
		assert.Equal(t, uint8(0b01001100), newState.V[0x1], "Vy should keep the value 0b01001100")
	})

	t.Run("(XOR Vx, Vy) Instruction 8xy3 should load the Vy BITWISE XOR Vx value into Vx", func(t *testing.T) {
		c := New()
		c.CurrState.V[0x0] = 0b00001000
		c.CurrState.V[0x1] = 0b01001100
		newState := c.ExecuteOpcode(0x8013)

		assert.Equal(t, uint8(0b01000100), newState.V[0x0], "Vx should have the value 0b01000100")
		assert.Equal(t, uint8(0b01001100), newState.V[0x1], "Vy should keep the value 0b01001100")
	})

	t.Run("(ADD Vx, Vy) Instruction 8xy4 should add the Vy value into the current Vx value", func(t *testing.T) {
		c := New()
		c.CurrState.V[0x3] = 0x01
		c.CurrState.V[0xE] = 0x01
		newState := c.ExecuteOpcode(0x83E4)
		assert.Equal(t, uint8(0x02), newState.V[0x3], "Vx should have the value of Vx + Vy")
		assert.NotEqual(t, uint8(0x01), newState.V[0xF], "VF should NOT be set to 1, since there wasn't a carry")
	})

	t.Run("(ADD Vx, Vy) Instruction 8xy4 should add the Vy value into the current Vx value and VF should be set to 1 when carry", func(t *testing.T) {
		c := New()
		c.CurrState.V[0x3] = 0xFF
		c.CurrState.V[0xE] = 0xFF
		newState := c.ExecuteOpcode(0x83E4)
		assert.Equal(t, uint8(0xFE), newState.V[0x3], "Vx should have the value of Vx + Vy")
		assert.Equal(t, uint8(0x01), newState.V[0xF], "VF should be set to 1, since there was a carry")
	})

	t.Run("(SUB Vx, Vy) Instruction 8xy5 should subtract the Vy value into the current Vx value", func(t *testing.T) {
		c := New()
		c.CurrState.V[0x3] = 0x01
		c.CurrState.V[0xE] = 0x01
		newState := c.ExecuteOpcode(0x83E5)
		assert.Equal(t, uint8(0x00), newState.V[0x3], "Vx should have the value of Vx - Vy")
		assert.NotEqual(t, uint8(0x01), newState.V[0xF], "VF should have the value 1, since NO borrow was made")
	})

	t.Run("(SUB Vx, Vy) Instruction 8xy5 should subtract the Vy value into the current Vx value and VF should be set to 1 when there was NO borrow", func(t *testing.T) {
		c := New()
		c.CurrState.V[0x3] = 0x01
		c.CurrState.V[0xE] = 0xFF
		newState := c.ExecuteOpcode(0x83E5)
		assert.Equal(t, uint8(0x02), newState.V[0x3], "Vx should have the value of Vx - Vy")
		assert.Equal(t, uint8(0x00), newState.V[0xF], "VF should have the value 0, since a borrow was made")
	})

	t.Run("(SHR Vx {, Vy}) Instruction 8xy6 should shift right the bits on Vx", func(t *testing.T) {
		c := New()
		c.CurrState.V[0x0] = 0b00010001
		newState := c.ExecuteOpcode(0x8016)
		assert.Equal(t, uint8(0b00001000), newState.V[0x0], "Vx bits should be shifted right once")
		assert.Equal(t, uint8(0x00), newState.V[0xF], "VF should be set to 0, since least significant byte is 0")
	})

	t.Run("(SHR Vx {, Vy}) Instruction 8xy6 should shift right the bits on Vx and VF should be set to 1 if least significant bit is 1", func(t *testing.T) {
		c := New()
		c.CurrState.V[0x0] = 0b00010011
		newState := c.ExecuteOpcode(0x8016)
		assert.Equal(t, uint8(0b00001001), newState.V[0x0], "Vx bits should be shifted right once")
		assert.Equal(t, uint8(0x01), newState.V[0xF], "VF should be set to 1, since least significant bit is 1")
	})

	t.Run("(SUBN Vx, Vy) Instruction 8xy7 should subtract the Vy value by the current Vx value into Vx", func(t *testing.T) {
		c := New()
		c.CurrState.V[0x0] = 0xFF
		c.CurrState.V[0x1] = 0x01
		newState := c.ExecuteOpcode(0x8017)
		assert.Equal(t, uint8(0x02), newState.V[0x0], "Vx should have the value of Vy - Vx")
		assert.Equal(t, uint8(0x00), newState.V[0xF], "VF should have the value 0, since a borrow was made")
	})

	t.Run("(SUBN Vx, Vy) Instruction 8xy7 should subtract the Vy value by the current Vx value into Vx and VF should be set to 1 when there was NO borrow", func(t *testing.T) {
		c := New()
		c.CurrState.V[0x0] = 0x01
		c.CurrState.V[0x1] = 0xFF
		newState := c.ExecuteOpcode(0x8017)
		assert.Equal(t, uint8(0xFE), newState.V[0x0], "Vx should have the value of Vy - Vx")
		assert.Equal(t, uint8(0x01), newState.V[0xF], "VF should have the value 1, since NO borrow was made")
	})

	t.Run("(SHL Vx {, Vy}) Instruction 8xyE should shift left the bits on Vx", func(t *testing.T) {
		c := New()
		c.CurrState.V[0x0] = 0b00010001
		newState := c.ExecuteOpcode(0x801E)
		assert.Equal(t, uint8(0b00100010), newState.V[0x0], "Vx bits should be shifted left once")
		assert.Equal(t, uint8(0x00), newState.V[0xF], "VF should be set to 0, since most significant byte is 0")
	})

	t.Run("(SHL Vx {, Vy}) Instruction 8xyE should shift left the bits on Vx and VF should be set to 1 if most significant bit is 1", func(t *testing.T) {
		c := New()
		c.CurrState.V[0x0] = 0b01010011
		newState := c.ExecuteOpcode(0x801E)
		assert.Equal(t, uint8(0b10100110), newState.V[0x0], "Vx bits should be shifted left once")
		assert.Equal(t, uint8(0x01), newState.V[0xF], "VF should be set to 1, since most significant bit is 1")
	})
}

/*
Todo:
	- Delay Timer e Sound timer
	- Sound interface
Rever os comandos:
	- 8xy4: ADD tem q mexer com flag e tal
	- 8xy5: SUB tem q mexer com flag e tal tbm
	- 8xy6, 8xy7, 8xyE mexem com flag
	- Dxyn: Esse vai ser complexo, o mais complexo até agora.
	- Ex9E e ExA1: Mexem com tecla, n sei como vou fazer
	- Fx0A: tecla tbm
	- Fx29: usa os digitos sprite
	- Fx33: chatinho só
	- Fx55 e Fx65: chatos tbm
*/

/*
	- Rever SHIFT RIGHT E SHIFT LEFT pra ver se VY precisa ser igual ao VX
*/
