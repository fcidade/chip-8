package chip8

type State struct {
	V          [0x10]uint8
	Memory     [0xFFF]uint8
	I          uint16
	PC         uint16
	DelayTimer uint8
	SoundTimer uint8
	SP         uint8
	Stack      [0xF]uint16
	Graphics   [ScreenHeight]uint64
	Keyboard   [0x10]bool
}

func (s *State) Opcode() uint16 {
	mostSignificantByte := uint16(s.Memory[s.PC]) << 8
	lessSignificantByte := uint16(s.Memory[s.PC+1])
	return mostSignificantByte | lessSignificantByte
}

func (s *State) GetPixel(x, y uint8) bool {
	return s.Graphics[y]&(FirstScreenBitMask>>x) != 0
}

func (s *State) SetPixel(x, y uint8) {
	xBitShifting := (x % ScreenWidth)
	clampedX := uint64(FirstScreenBitMask) >> xBitShifting
	clampedY := y % ScreenHeight
	isAlreadyPainted := s.GetPixel(xBitShifting, clampedY)
	if isAlreadyPainted {
		s.Graphics[clampedY] &= ^uint64(clampedX)
	} else {
		s.Graphics[clampedY] |= clampedX
	}
}
