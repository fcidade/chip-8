package chip8

import (
	"fmt"
	"math/rand"
	"testing"

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

		c.CurrState.Graphics = [32]uint64{0, 1, 2, 3, 4}
		newState := c.ExecuteOpcode(0x00E0)

		assert.Equal(t, newState.Graphics, [32]uint64{}, "Graphics should be all zeroes")
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

	t.Run("(SNE Vx, Vy) Instruction 9xkk skips next instruction if Vx is NOT equals Vy", func(t *testing.T) {
		c := New()
		c.CurrState.PC = 0x200
		c.CurrState.V[0x0] = 0x00
		c.CurrState.V[0x1] = 0x01

		newState := c.ExecuteOpcode(0x9010)

		assert.NotEqual(t, uint8(0xFF), newState.V[0x0], "Vx should NOT have the same value as the received")
		assert.Equal(t, uint16(0x202), newState.PC, "Program Counter should increment by 2")
	})

	t.Run("(SNE Vx, Vy) Instruction 9xkk should NOT skip next instruction if Vx equals Vy", func(t *testing.T) {
		c := New()
		c.CurrState.PC = 0x200
		c.CurrState.V[0x0] = 0xFF
		c.CurrState.V[0x1] = 0xFF

		newState := c.ExecuteOpcode(0x9010)

		assert.Equal(t, uint8(0xFF), newState.V[0x0], "Vx should have the value same value as the received")
		assert.Equal(t, uint16(0x200), newState.PC, "Program Counter should remain the same")
	})

	t.Run("(LD I, addr) Instruction Annn should load the received address into I", func(t *testing.T) {
		c := New()
		newState := c.ExecuteOpcode(0xA333)
		assert.Equal(t, uint16(0x333), newState.I, "I should have the received address")
	})

	t.Run("(JMP V0, addr) Instruction Bnnn should jump the program counter to the received address + V0", func(t *testing.T) {
		c := New()
		c.CurrState.V[0x0] = 0x10
		newState := c.ExecuteOpcode(0xB333)
		assert.Equal(t, uint16(0x343), newState.PC, "Program Counter should have the V0 value + the received address")
	})

	t.Run("(RND Vx, byte) Instruction Cxkk should load a random value into Vx BITWISE AND received value", func(t *testing.T) {
		c := New()
		rand.Seed(1)
		newState := c.ExecuteOpcode(0xC101)
		assert.Equal(t, uint8(0x01), newState.V[0x1], "Program Counter should have the V0 value + the received address")
	})

	t.Run("(LD Vx, DT) Instruction Fx07 should load the Delay Timer into Vx", func(t *testing.T) {
		c := New()
		c.CurrState.DelayTimer = 0x34
		newState := c.ExecuteOpcode(0xF007)
		assert.Equal(t, uint8(0x34), newState.V[0x0], "Vx should have the value of the Delay Timer")
	})

	t.Run("(LD Vx, DT) Instruction Fx15 should load the Vx value into Delay Timer", func(t *testing.T) {
		c := New()
		c.CurrState.V[0x1] = 0x34
		newState := c.ExecuteOpcode(0xF115)
		assert.Equal(t, uint8(0x34), newState.DelayTimer, "Delay Timer should have the value of Vx")
	})

	t.Run("(LD Vx, ST) Instruction Fx18 should load the Vx value into Sound Timer", func(t *testing.T) {
		c := New()
		c.CurrState.V[0x1] = 0x34
		newState := c.ExecuteOpcode(0xF118)
		assert.Equal(t, uint8(0x34), newState.SoundTimer, "Sound Timer should have the value of Vx")
	})

	t.Run("(ADD I, Vx) Instruction Fx1E should add the value of Vx into the existing value in I", func(t *testing.T) {
		c := New()
		c.CurrState.V[0x1] = 0x60
		c.CurrState.I = 0x100
		newState := c.ExecuteOpcode(0xF11E)
		assert.Equal(t, uint16(0x160), newState.I, "I should have the value of I + Vx")
	})

	t.Run("(LD F, Vx) Instruction Fx29 should load the address of the Vx character sprite into I", func(t *testing.T) {
		c := New()
		c.LoadFonts()
		c.CurrState.V[0x1] = 0xA
		newState := c.ExecuteOpcode(0xF129)
		assert.Equal(t, uint16(0x82), newState.I, "Should be at the right address")
		assert.Equal(t, []uint8{0xF0, 0x90, 0xF0, 0x90, 0x90}, newState.Memory[newState.I:newState.I+5], "Should have loaded the right sprite")
	})

	t.Run("(LD B, Vx) Instruction Fx33 should load the Vx digits into Memory at I, I+1 and I+3", func(t *testing.T) {
		c := New()
		c.CurrState.V[0x1] = 237
		c.CurrState.I = 0x210
		newState := c.ExecuteOpcode(0xF133)
		assert.Equal(t, uint8(2), newState.Memory[0x210], "Should have the right digit")
		assert.Equal(t, uint8(3), newState.Memory[0x211], "Should have the right digit")
		assert.Equal(t, uint8(7), newState.Memory[0x212], "Should have the right digit")
	})

	t.Run("(LD [I], Vx) Instruction Fx55 should loads the V[0:x] into memory starting by I", func(t *testing.T) {
		c := New()
		c.CurrState.V[0x0] = 0x00
		c.CurrState.V[0x1] = 0x01
		c.CurrState.V[0x2] = 0x02
		c.CurrState.V[0x3] = 0x03
		c.CurrState.V[0x4] = 0x04
		c.CurrState.V[0x5] = 0x05
		c.CurrState.V[0x6] = 0x06
		c.CurrState.I = 0x210
		newState := c.ExecuteOpcode(0xF655)
		assert.Equal(t, uint8(0x00), newState.Memory[0x210], "Should have the right value")
		assert.Equal(t, uint8(0x01), newState.Memory[0x211], "Should have the right value")
		assert.Equal(t, uint8(0x02), newState.Memory[0x212], "Should have the right value")
		assert.Equal(t, uint8(0x03), newState.Memory[0x213], "Should have the right value")
		assert.Equal(t, uint8(0x04), newState.Memory[0x214], "Should have the right value")
		assert.Equal(t, uint8(0x05), newState.Memory[0x215], "Should have the right value")
		assert.Equal(t, uint8(0x06), newState.Memory[0x216], "Should have the right value")
	})

	t.Run("(LD Vx, [I]) Instruction Fx65 should loads the into V[0:x] the memory values starting by I", func(t *testing.T) {
		c := New()
		c.CurrState.Memory[0x200] = 0x00
		c.CurrState.Memory[0x201] = 0x01
		c.CurrState.Memory[0x202] = 0x02
		c.CurrState.Memory[0x203] = 0x03
		c.CurrState.Memory[0x204] = 0x04
		c.CurrState.Memory[0x205] = 0x05
		c.CurrState.Memory[0x206] = 0x06
		c.CurrState.I = 0x200
		newState := c.ExecuteOpcode(0xF665)
		assert.Equal(t, uint8(0x00), newState.V[0x0], "Should have the right value")
		assert.Equal(t, uint8(0x01), newState.V[0x1], "Should have the right value")
		assert.Equal(t, uint8(0x02), newState.V[0x2], "Should have the right value")
		assert.Equal(t, uint8(0x03), newState.V[0x3], "Should have the right value")
		assert.Equal(t, uint8(0x04), newState.V[0x4], "Should have the right value")
		assert.Equal(t, uint8(0x05), newState.V[0x5], "Should have the right value")
		assert.Equal(t, uint8(0x06), newState.V[0x6], "Should have the right value")
	})

}

/*
Todo:
	- Delay Timer e Sound timer
Rever os comandos:
	- Fx29: usa os digitos sprite
	- Fx33: chatinho só
	- Fx55 e Fx65: chatos tbm

	- Ex9E e ExA1: Mexem com tecla, n sei como vou fazer
	- Fx0A: tecla tbm
*/
