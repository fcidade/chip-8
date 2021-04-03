package chip8

import (
	"fmt"
	"testing"
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
	t.Run("", func(t *testing.T) {
		c := &Chip8{}
		r := c.ExecuteOpcode(0x00E0)
		fmt.Println(r)
	})
}
