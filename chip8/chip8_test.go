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
		c := &Chip8State{}
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

	t.Run("(SNE Vx, byte) Instruction 3xkk skips next instruction if Vx is NOT equals kk", func(t *testing.T) {
		c := New()
		c.CurrState.PC = 0x200
		c.CurrState.V[0x0] = 0x00

		newState := c.ExecuteOpcode(0x40FF)

		assert.NotEqual(t, uint8(0xFF), newState.V[0x0], "Vx should NOT have the value 0xFF")
		assert.Equal(t, uint16(0x202), newState.PC, "Program Counter should increment by 2")
	})

	t.Run("(SNE Vx, byte) Instruction 3xkk should NOT skip next instruction if Vx equals kk", func(t *testing.T) {
		c := New()
		c.CurrState.PC = 0x200
		c.CurrState.V[0x0] = 0xFF

		newState := c.ExecuteOpcode(0x40FF)

		assert.Equal(t, uint8(0xFF), newState.V[0x0], "Vx should have the value 0xFF")
		assert.Equal(t, uint16(0x200), newState.PC, "Program Counter should remain the same")
	})
}
