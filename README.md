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
  The stack is an array of 16 16-bit values, used to store the address of the return point on subroutines call
  Chip-8 alows for up to 16 levels of nested subroutines

The stack
Return addresses when subroutines are called
Allocated 48 bytes for up to (*12 or 16? idk, is wikipedia wrong?*) levels of nesting

Timers:
Both count down at 60 hetz, until they reach 0
- Delay timer: timing game events, READ & WRITE
- Sound timer: Sound effects, when it's value is nonzero, a beeping sound is made. The beeping frequence is defined by the interpreter's writer (me)

Input/Keyboard:
Hex Keyboard -> 16 keys from 0 to F
6, 4, 6 and 2 are typically used for directional input

Original 16-key hexadecimal keyboard layout:
|1|2|3|C|
|4|5|6|D|
|7|8|9|E|
|A|0|B|F|

[PUT INFO ABOUT OPCODES]

Graphics:
Resolution is 64x32 px
Color is monochrome
Graphics are drawn by drawing sprites
sprites are 8 pixels wide and may be from 1 to 16 px in height
A sprite is a group of bytes which are a binary representation of the desired picture
There's a group of sprites that can be references on memory 0x000 to 0x1ff, you can see more [here](http://devernay.free.fr/hacks/chip8/C8TECH10.HTM#2.4)

! Opcodes:
The original implementation includes 36 instructions
Allows for math, graphics and flow control functions
All instructions are 2 bytes long and are stored most-significant-byte first (big endian i guess)
In memory, the first byte of each instruction should be located at an even addresses.


Adresses are always 12 bits (?)
Can have 8 or 4 bit constants

http://devernay.free.fr/hacks/chip8/C8TECH10.HTM
https://tobiasvl.github.io/blog/write-a-chip-8-emulator/