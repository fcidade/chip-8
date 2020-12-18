# ! This is a WORK IN PROGRESS!!! !

# Chip-8 Emulator

Memory map: 4096 (from 0x000 to 0xFFF)
chip8: 512 (0x000 to 0x1FF)
programs: 
  Mostly start at 0x200 to 0xFFF
  But some begin at 0x600 to 0xFFF
display refresh: 0xF00 -> 0xFFF
call stack, internal use and other variables: 0xEA0-0xEFF

Registers:
  Chip-8 has 16 general porpuse 8-bit registers
  V0 to VF
  VF (alias "I") can be a flag to some instructions, this, it should be avoided
  In ADD, VF is the carry flag
  in SUB, it's the "no borrow" flag
  and in the draw instruction VF is ser upon pixel collision
  There's also two 8-bit registers for timers, when they're not 0, they're decremented in a 60hz rate.
  PC -> Program Counter (16-bit) -> stores the current execuring address
  SP -> Stack pointer (8-bit) -> points to the topmost level of the stack

The stack
Return addresses when subroutines are called
Allocated 48 bytes for up to 12 levels of nesting

Timers:
Both count down at 60 hetz, until they reach 0
- Delay timer: timing game events, READ & WRITE
- Sound timer: Sound effects, when it's value is nonzero, a beeping sound is made.

Input:
Hex Keyboard -> 16 keys from 0 to F
6, 4, 6 and 2 are typically used for directional input
[PUT INFO ABOUT OPCODES]

Graphics:
Resolution is 64x32 px
Color is monochrome
Graphics are drawn by drawing sprites
sprites are 8 pixels wide and may be from 1 to 16 px in height

! Opcodes:
Adresses are always 12 bits (?)
Can have 8 or 4 bit constants

http://devernay.free.fr/hacks/chip8/C8TECH10.HTM