package chip8

type Chip8State struct {
	V          [0xF]uint8
	I          uint16
	DelayTimer uint8
	SoundTimer uint8
	Memory     [0xFFF]uint8
	Stack      [0xF]uint16
	SP         uint8
	PC         uint16
}

func (c *Chip8State) FetchOpcode() uint16 {
	mostSignificantByte := uint16(c.Memory[c.PC]) << 8
	c.PC++
	lessSignificantByte := uint16(c.Memory[c.PC])
	c.PC++
	return mostSignificantByte | lessSignificantByte
}
