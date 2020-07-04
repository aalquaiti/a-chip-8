// Package chip provide functionality for the chip-8 interpreter
package chip

import (
	"fmt"
	"math/rand"
	"time"
)

// ram should be access from Ox200 (512) and upward. The first 512 bytes are
// reseverd for the interpreter
var ram [0x1000]uint8

//General purpose Registers. Commonly refered to as Vx, where x is a
// hexademical digit (0 to F)
var regs [0x10]uint8

/*
Used to store memore addresses
*/
var iReg uint16

/*
Special purpose register used as a delay time register (DT)
*/
var dtReg uint8

/*
Special purpse register used as sound timer register (ST)
*/
var stReg uint8

/*
Program Counter
*/
var pc uint16

/*
Stack Pointer
*/
var sp uint8

/*
Stack allows up to 16 levels of nested subroutines
*/
var stack [0x10]uint16

// Offical Instruction set
var inst [0x10]func(uint8, uint8, uint8)

// Initialie interpreter
func init() {
	// TODO: Add pre defined sprites to first part of memory
	// (0 up to before 0x200)
	rand.Seed(time.Now().Local().UnixNano())

	inst[0x0] = func(n1, n2, n3 uint8) {
		switch {
		// 00E0
		case createNipple3(n1, n2, n3) == 0x00E0:
			screen.Reset()
		// 00EE
		case createNipple3(n1, n2, n3) == 0x00EE:
			pc = stack[sp]
			sp--

		// op 0x0NNN is considered invalid
		default:
			panic("Unsupported op code")
		}

	}

	inst[0x1] = func(n1, n2, n3 uint8) {
		pc = createNipple3(n1, n2, n3)
	}

	inst[0x2] = func(n1, n2, n3 uint8) {
		sp++
		stack[sp] = pc
		pc = createNipple3(n1, n2, n3)
	}

	inst[0x3] = func(n1, n2, n3 uint8) {
		if regs[n1] == createNipple2(n2, n3) {
			pc += 2
		}
	}

	inst[0x4] = func(n1, n2, n3 uint8) {
		if regs[n1] != createNipple2(n2, n3) {
			pc += 2
		}

	}

	inst[0x5] = func(n1, n2, n3 uint8) {
		if n3 != 0 {
			panic("Unsupported op code")
		}

		if regs[n1] == regs[n2] {
			pc += 2
		}

	}

	inst[0x6] = func(n1, n2, n3 uint8) {
		regs[n1] = createNipple2(n2, n3)
	}

	inst[0x7] = func(n1, n2, n3 uint8) {
		regs[n1] += createNipple2(n2, n3)
	}

	inst[0x8] = func(n1, n2, n3 uint8) {
		switch n3 {
		case 0:
			regs[n1] = regs[n2]

		case 1:
			regs[n1] |= regs[n2]

		case 2:
			regs[n1] &= regs[n2]

		case 3:
			regs[n1] ^= regs[n2]

		case 4:
			regs[n1], regs[0xF] = add8(regs[n1], regs[n2])

		case 5:
			regs[n1], regs[0xF] = sub8(regs[n1], regs[n2])

		case 6:
			regs[n1] = regs[n2] >> 1
			regs[0xF] = regs[n2] & 0x1

		case 7:
			regs[n1], regs[0xF] = sub8(regs[n2], regs[n1])

		case 0xE:
			regs[n1] = regs[n2] << 1
			regs[0xF] = (regs[n2] & 0b1000) >> 3

		default:
			panic("Unsupported op code")
		}
	}

	inst[0x9] = func(n1, n2, n3 uint8) {
		if n3 != 0 {
			panic("Unsupported op code")
		}

		if regs[n1] != regs[n2] {
			pc += 2
		}

	}

	inst[0xA] = func(n1, n2, n3 uint8) {
		iReg = createNipple3(n1, n2, 3)
	}

	inst[0xB] = func(n1, n2, n3 uint8) {
		pc = createNipple3(n1, n2, n3) + uint16(regs[0])
	}

	inst[0xC] = func(n1, n2, n3 uint8) {
		random := uint8(rand.Int31n(256))
		value := createNipple2(n2, n3)
		regs[n1] = value & random
	}

	inst[0xD] = func(n1, n2, n3 uint8) {

		// Draw a sprite on screen at posiont x (n1) and y (n2)
		// Draw happens in Xor Mode, i.e:
		// pixel = current state XOR new state
		// Each sprite width is 8 bits with height n3
		// Set VF = 1 if collision happens
		// Collision is when a written pixel (white) is to be unwritten (black)
		regs[0xF] = 0

		for j := 0; j < int(n3); j++ {
			value := int(ram[int(iReg)+j])

			for i := 0; i < 8; i-- {
				x := (int(regs[n1]) + i) % screenWidth
				y := (int(regs[n2]) + j) % screenHeight

				// Use a mask to extract the bit
				mask := 0x8 >> i
				pixelSet := value&mask == mask

				// if pixel is to be written (i.e bit is set to 1)
				if pixelSet {

					// By now, we know pixel is set
					// If current pixel drawn is set, then a flip will occur
					// indicating a collision
					if screen.PixelAt(x, y) {
						regs[0xF] = 1
						pixelSet = false
					}

					screen.Draw(x+i, y+j, pixelSet)
				}
			}
		}
	}

	inst[0xE] = func(n1, n2, n3 uint8) {
		switch createNipple2(n2, n3) {
		case 0x9E:
			// TODO complete
			fmt.Println("Skipping next instruction when key is pressed")
		case 0xA1:
			// TODO complete
			fmt.Println("Skipping next instruction when key is not pressed")
		default:
			panic("Unsupported op code")
		}
	}

	inst[0xF] = func(n1, n2, n3 uint8) {
		switch createNipple2(n2, n3) {
		case 0x07:
			regs[n1] = dtReg
		case 0xA:
			// TODO complete
			fmt.Println("Waiting for a key to be pressed")
		case 0x15:
			dtReg = regs[n1]
		case 0x18:
			stReg = regs[n1]
		case 0x1E:
			iReg = uint16(regs[n1])
		case 0x29:
			// TODO complete
			fmt.Println("Updating I to a sprite memory address")
		case 0x33:
			ram[iReg], ram[iReg+1], ram[iReg+2] = toBCD(regs[n1])
		case 0x55:
			for i := uint8(0); i < n1; i++ {
				ram[iReg] = regs[i]
				iReg++
			}
			iReg++
		case 0x65:
			for i := uint8(0); i < n1; i++ {
				regs[i] = ram[iReg]
				iReg++
			}
			iReg++
		default:
			panic("Unsupported op code")
		}
	}
}

func tick() {
	if dtReg > 0 {
		dtReg--
	}
	// Add sound effect for when sound timer is active and decrement by one
}
