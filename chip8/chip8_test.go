package chip8

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestChip8State(t *testing.T) {
	t.Run("", func(t *testing.T) {
		t.Skip()
		c := &Chip8State{}
		c.Memory[0] = 0x12
		c.Memory[1] = 0x34
		fmt.Printf("0x%04x\n", c.FetchOpcode())
	})
}

func TestChip8(t *testing.T) {
	t.Run("(SYS) Instructions on range 0nnn should be ignored, as they are actually SYS calls", func(t *testing.T) {
		c := New()
		oldState := c.CurrState
		newState := c.ExecuteOpcode(0x0000)
		assert.Equal(t, oldState, newState, "State should not be altered")
	})

	t.Run("(CLS) Instruction 00E0 should clear the screen", func(t *testing.T) {
		// TODO!
		t.Skip()
		c := New()
		oldState := c.CurrState
		newState := c.ExecuteOpcode(0x0000)
		assert.Equal(t, oldState, newState, "")
	})

	t.Run("(RET) Instruction 00EE should return from a subroutine", func(t *testing.T) {
		c := New()
		c.CurrState.Stack[0x0] = 0x222
		c.CurrState.SP = 0x1
		newState := c.ExecuteOpcode(0x00EE)
		assert.Equal(t, uint16(0x222), newState.PC, "Should set Program Counter back to the value on the top of the stack")
		assert.Equal(t, uint8(0x0), newState.SP, "Stack pointer should decrement by 1")
	})

	t.Run("(JMP) Instruction 1nnn should set program counter to the be received address", func(t *testing.T) {
		c := New()
		newState := c.ExecuteOpcode(0x1333)
		assert.Equal(t, uint16(0x333), newState.PC, "Program Counter should be equal to the received address")
	})

	t.Run("(CALL) Instruction 2nnn should call subroutine on address received", func(t *testing.T) {
		c := New()
		c.CurrState.PC = 0x250
		newState := c.ExecuteOpcode(0x2325)
		assert.Equal(t, uint16(0x325), newState.PC, "Program Counter should be equal to the received address")
		assert.Equal(t, uint8(0x1), newState.SP, "Stack Pointer should be incremented by 1")
		assert.Equal(t, uint16(0x250), newState.Stack[0x0], "First element of the Stack should contain the previous Program Counter")
	})

}
