package chip8

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

func (s *State) Opcode() uint16 {
	mostSignificantByte := uint16(s.Memory[s.PC]) << 8
	lessSignificantByte := uint16(s.Memory[s.PC+1])
	return mostSignificantByte | lessSignificantByte
}
