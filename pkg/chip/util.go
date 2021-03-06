// Package chip provide functionality for the chip-8 interpreter
package chip

/*
Converts byte to Binary-Coded Decimal using double dabble algorithm
Returns hundreth, tenth and unit as (x, y, z) respectively
*/
func toBCD(n uint8) (x, y, z uint8) {
	value := uint32(n)

	// Get the position of the most significant bit
	i := 8
	for i >= 0 {
		if n&0x80 == 0x80 {
			break
		}
		i--
		n >>= 1
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
	if x > y {
		borrow = 1
	}

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
