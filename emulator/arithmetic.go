package emulator

// condition bits
// --------------

// SetZero sets the zero condition bit depending on a value
func SetZero(cc ConditionBits, value uint16) {
	if value&0xff == 0 {
		cc.zero = 1
	} else {
		cc.zero = 0
	}
}

// SetSign sets the sign condition bit depending on a value
func SetSign(cc ConditionBits, value uint16) {
	if value&0x80 != 0 {
		cc.sign = 1
	} else {
		cc.sign = 0
	}
}

// SetCarry sets the carry condition bit depending on a value
func SetCarry(cc ConditionBits, value uint16) {
	if value > 0xff {
		cc.carry = 1
	} else {
		cc.carry = 0
	}
}

// SetParity sets the parity condition bit depending on a value
func SetParity(cc ConditionBits, value uint16) {
	// based on https://stackoverflow.com/a/21618038 - by user TypeIA
	x := byte(value & 0xff)
	x ^= x >> 4
	x ^= x >> 2
	x ^= x >> 1
	result := x & 1
	if result == 0 {
		cc.parity = 1
	} else {
		cc.parity = 0
	}
}

// operations
// --------------

// PersistArithmeticResult persists the result of most arithmetic operations in the emulator
// state, by setting the zero, sign, carry, auxiliary carry (TODO) and parity condition
// flags and setting the accumulator registry to the least significant byte of the result
func PersistArithmeticResult(state *State, result uint16) {
	// condition bits
	SetZero(state.cc, result)
	SetSign(state.cc, result)
	SetCarry(state.cc, result)
	// TODO: set auxiliary carry
	SetParity(state.cc, result)

	state.a = byte(result & 0xff)

}

// ADD adds a value to the accumulator register
func ADD(state *State, value byte) {
	result := uint16(state.a) + uint16(value)
	PersistArithmeticResult(state, result)
}

// ADC adds a value and the carry condition bit to the accumulator register
func ADC(state *State, value byte) {
	result := uint16(state.a) + uint16(value) + uint16(state.cc.carry)
	PersistArithmeticResult(state, result)
}

// SUB substracts a value from the accumulator register
func SUB(state *State, value byte) {
	// TODO: does this work like two's complement arithmetic??
	result := uint16(state.a) - uint16(value)
	PersistArithmeticResult(state, result)
}
