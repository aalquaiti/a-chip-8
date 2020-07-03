package main

import (
	"fmt"
	"math/rand"
	"time"
)

/*
RAM should be access from Ox200 (512) and upward. The first 512 bytes are reseverd
for the interpreter
*/
var ram [0x1000]uint8

/*
General purpose Registers. Commonly refered to as Vx, where x is a hexademical digit (0 to F)
*/
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

// TODO: Add pre defined sprites to first part of memory (0 up to before 0x200)

var inst [0x10]func(uint8, uint8, uint8)

// Initialie interpreter
func init() {
	rand.Seed(time.Now().Local().UnixNano())

	inst[0x0] = func(n1, n2, n3 uint8) {
		switch {
		// 00E0
		case createNipple3(n1, n2, n3) == 0x00E0:
			fmt.Println("Clearn Screen")
		// 00EE
		case createNipple3(n1, n2, n3) == 0x00EE:
			fmt.Println("Return from subroutine")
		default:
			panic("Unsupported op code")
		}

	}

	inst[0x1] = func(n1, n2, n3 uint8) {
		pc = createNipple3(n1, n2, n3)
		fmt.Printf("PC jumped to %.4X\n", pc)
	}

	inst[0x2] = func(n1, n2, n3 uint8) {
		// TODO update stack
		pc = createNipple3(n1, n2, n3)
		fmt.Printf("Executing subroutine at %.4X\n", pc)
	}

	inst[0x3] = func(n1, n2, n3 uint8) {
		if regs[n1] == createNipple2(n2, n3) {
			fmt.Println("Skipping next instruction")
		}

	}

	inst[0x4] = func(n1, n2, n3 uint8) {
		if regs[n1] != createNipple2(n2, n3) {
			fmt.Println("Skipping next instruction")
		}

	}

	inst[0x5] = func(n1, n2, n3 uint8) {
		if n3 != 0 {
			panic("Unsupported op code")
		}

		if regs[n1] == regs[n2] {
			fmt.Println("Skipping next instruction")
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
			fmt.Println("Skipping next instruction")
		}

	}

	inst[0xA] = func(n1, n2, n3 uint8) {
		iReg = createNipple3(n1, n2, 3)
	}

	inst[0xB] = func(n1, n2, n3 uint8) {
		pc = createNipple3(n1, n2, n3) + uint16(regs[0])
	}

	inst[0xC] = func(n1, n2, n3 uint8) {
		value := int32(createNipple2(n2, n3))
		regs[n1] = uint8(rand.Int31n(value))
	}

	inst[0xD] = func(n1, n2, n3 uint8) {
		fmt.Println("Draw stuff on screen")
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

		}
	}

}

/*
Converts byte to Binary-Coded Decimal using double dabble algorithm
Returns hundreth, tenth and unit as (x, y, z) respectively
*/
func toBCD(n uint8) (x, y, z uint8) {
	value := uint32(n)

	// Get the position of the most significant bit
	i := 8
	for {
		if n&0x10 == 0x10 {
			break
		}
		i--
		n <<= 1
	}

	for i > 0 {
		// Checking if unit is above 5
		if (value & 0xF00) >= 0x500 {
			value += 0x300
		}

		// Checking if tenth is above 5
		if (value & 0xF000) >= 0x5000 {
			value += 0x3000
		}

		// Checking if hundreth is above 5
		if (value & 0xF0000) >= 0x50000 {
			value += 0x30000
		}

		value <<= 1
		i--
	}

	x = uint8((value & 0XF0000) >> 16)
	y = uint8((value & 0XF000) >> 12)
	z = uint8((value & 0XF00) >> 8)
	return
}

/*
Adds two bytes
Returns sum as byte with carry
*/
func add8(x, y uint8) (sum, carry uint8) {
	sum16 := uint16(x) + uint16(y)
	sum = uint8(sum16)
	carry = uint8(sum16 >> 8)

	return
}

/*
Subtracts two bytes
Returns difference as byte with borrow
*/
func sub8(x, y uint8) (diff, borrow uint8) {
	diff = x - y
	borrow = ((^x & y) | (^(x ^ y) & diff)) >> 7

	return
}

// Takes three nibbles (4 bits) and turn them into 12 bit value
// returned as 16 bits
func createNipple3(n1, n2, n3 uint8) uint16 {
	return (uint16(n1) << 8) + uint16((n2<<4)+n3)
}

// Takes two nibbles and turn them into 8 bit vlaue
// returned as a byte
func createNipple2(n1, n2 uint8) uint8 {
	return (n1 << 4) + n2
}

func main() {
	fmt.Println("Hello from a-chip-8")
	x, y, z := toBCD(20)

	fmt.Println(x, y, z)
}
