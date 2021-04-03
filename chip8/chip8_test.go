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
		fmt.Printf("0x%04x\n", c.Opcode())
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
		// TODO!
		t.Skip()
		c := New()
		oldState := c.CurrState
		newState := c.ExecuteOpcode(0x0000)
		assert.Equal(t, oldState, newState, "")
	})
}
