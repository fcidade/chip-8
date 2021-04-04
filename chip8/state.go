package chip8

import "math"

type State struct {
	V          [0x10]uint8
	I          uint16
	DelayTimer uint8
	SoundTimer uint8
	Memory     [0xFFF]uint8
	Stack      [0xF]uint16
	SP         uint8
	PC         uint16
}

func (c *State) Opcode() uint16 {
	mostSignificantByte := uint16(c.Memory[c.PC]) << 8
	lessSignificantByte := uint16(c.Memory[c.PC+1])
	return mostSignificantByte | lessSignificantByte
}

func (c *State) FetchNext() {
	c.PC += 0x2
	c.DelayTimer = uint8(math.Max(0, float64(c.DelayTimer-1)))
	c.SoundTimer = uint8(math.Max(0, float64(c.DelayTimer-1)))
}
