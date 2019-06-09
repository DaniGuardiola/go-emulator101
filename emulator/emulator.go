package emulator

import "os"

/*
Original (C?, C++?) code from emulator101.com looks like this:
typedef struct ConditionCodes {
	uint8_t    z:1;
	uint8_t    s:1;
	uint8_t    p:1;
	uint8_t    cy:1;
	uint8_t    ac:1;
	uint8_t    pad:3;
} ConditionCodes;

TODO: find out what that ":1" thing means and how to do that in go
*/

// ConditionBits contains the current state of the condition bits of the emulator
type ConditionBits struct {
	zero           byte // zero
	sign           byte // sign
	parity         byte // parity
	carry          byte // carry
	auxiliaryCarry byte // auxiliary carry
}

// State contains the state of the processor emulator
type State struct {
	// registers
	a byte
	b byte
	c byte
	d byte
	e byte
	h byte
	l byte

	sp        uint16        // stack pointer
	pc        uint16        // program counter
	memory    []byte        // memory
	cc        ConditionBits // condition bits
	intEnable byte          // TODO: idk
}

func unimplementedInstruction(state *State) {
	print("Unimplemented instruction! :(")
	state.pc--
	os.Exit(1)
}

func tick(state *State) {
	opcode := state.memory[state.pc]
	switch opcode {
	case 0x00: // NOP
		break
	case 0x01: // LXI B,D16
		state.c = state.memory[state.pc+1]
		state.b = state.memory[state.pc+2]
		state.pc += 2
	case 0x02:
		unimplementedInstruction(state)
	case 0x03:
		unimplementedInstruction(state)
	case 0x04:
		unimplementedInstruction(state)
	case 0x05:
		unimplementedInstruction(state)
	case 0x06:
		unimplementedInstruction(state)
	case 0x07:
		unimplementedInstruction(state)
	case 0x08:
		unimplementedInstruction(state)
	case 0x09:
		unimplementedInstruction(state)
	case 0x0a:
		unimplementedInstruction(state)
	case 0x0b:
		unimplementedInstruction(state)
	case 0x0c:
		unimplementedInstruction(state)
	case 0x0d:
		unimplementedInstruction(state)
	case 0x0e:
		unimplementedInstruction(state)
	case 0x0f:
		unimplementedInstruction(state)
	case 0x10:
		unimplementedInstruction(state)
	case 0x11: // LXI D,D16
		state.e = state.memory[state.pc+1]
		state.d = state.memory[state.pc+2]
		state.pc += 2
	case 0x12:
		unimplementedInstruction(state)
	case 0x13:
		unimplementedInstruction(state)
	case 0x14:
		unimplementedInstruction(state)
	case 0x15:
		unimplementedInstruction(state)
	case 0x16:
		unimplementedInstruction(state)
	case 0x17:
		unimplementedInstruction(state)
	case 0x18:
		unimplementedInstruction(state)
	case 0x19:
		unimplementedInstruction(state)
	case 0x1a:
		unimplementedInstruction(state)
	case 0x1b:
		unimplementedInstruction(state)
	case 0x1c:
		unimplementedInstruction(state)
	case 0x1d:
		unimplementedInstruction(state)
	case 0x1e:
		unimplementedInstruction(state)
	case 0x1f:
		unimplementedInstruction(state)
	case 0x20:
		unimplementedInstruction(state)
	case 0x21: // LXI D,D16
		state.l = state.memory[state.pc+1]
		state.h = state.memory[state.pc+2]
		state.pc += 2
	case 0x22:
		unimplementedInstruction(state)
	case 0x23:
		unimplementedInstruction(state)
	case 0x24:
		unimplementedInstruction(state)
	case 0x25:
		unimplementedInstruction(state)
	case 0x26:
		unimplementedInstruction(state)
	case 0x27:
		unimplementedInstruction(state)
	case 0x28:
		unimplementedInstruction(state)
	case 0x29:
		unimplementedInstruction(state)
	case 0x2a:
		unimplementedInstruction(state)
	case 0x2b:
		unimplementedInstruction(state)
	case 0x2c:
		unimplementedInstruction(state)
	case 0x2d:
		unimplementedInstruction(state)
	case 0x2e:
		unimplementedInstruction(state)
	case 0x2f:
		unimplementedInstruction(state)
	case 0x30:
		unimplementedInstruction(state)
	case 0x31: // LXI SP,D16
		state.sp = uint16(state.memory[state.pc+2])<<8 | uint16(state.memory[state.pc+1])
		state.pc += 2
	case 0x32:
		unimplementedInstruction(state)
	case 0x33:
		unimplementedInstruction(state)
	case 0x34:
		unimplementedInstruction(state)
	case 0x35:
		unimplementedInstruction(state)
	case 0x36:
		unimplementedInstruction(state)
	case 0x37:
		unimplementedInstruction(state)
	case 0x38:
		unimplementedInstruction(state)
	case 0x39:
		unimplementedInstruction(state)
	case 0x3a:
		unimplementedInstruction(state)
	case 0x3b:
		unimplementedInstruction(state)
	case 0x3c:
		unimplementedInstruction(state)
	case 0x3d:
		unimplementedInstruction(state)
	case 0x3e:
		unimplementedInstruction(state)
	case 0x3f:
		unimplementedInstruction(state)
	case 0x40: // MOV B,B (NOP)
		break
	case 0x41: // MOV B,C
		state.b = state.c
	case 0x42: // MOV B,D
		state.b = state.d
	case 0x43: // MOV B,E
		state.b = state.e
	case 0x44: // MOV B,H
		state.b = state.h
	case 0x45: // MOV B,L
		state.b = state.l
	case 0x46: // MOV B,M
		state.b = GetMemoryHL(state)
	case 0x47: // MOV B,A
		state.b = state.a
	case 0x48: // MOV C,B
		state.c = state.b
	case 0x49: // MOV C,C (NOP)
		break
	case 0x4a: // MOV C,D
		state.c = state.d
	case 0x4b: // MOV C,E
		state.c = state.e
	case 0x4c: // MOV C,H
		state.c = state.h
	case 0x4d: // MOV C,L
		state.c = state.l
	case 0x4e: // MOV C,M
		state.c = GetMemoryHL(state)
	case 0x4f: // MOV C,A
		state.c = state.a
	case 0x50: // MOV D,B
		state.d = state.b
	case 0x51: // MOV D,C
		state.d = state.c
	case 0x52: // MOV D,D (NOP)
		break
	case 0x53: // MOV D,E
		state.d = state.e
	case 0x54: // MOV D,H
		state.d = state.h
	case 0x55: // MOV D,L
		state.d = state.l
	case 0x56: // MOV D,M
		state.d = GetMemoryHL(state)
	case 0x57: // MOV D,A
		state.d = state.a
	case 0x58: // MOV E,B
		state.e = state.b
	case 0x59: // MOV E,C
		state.e = state.c
	case 0x5a: // MOV E,D
		state.e = state.d
	case 0x5b: // MOV E,E (NOP)
		break
	case 0x5c: // MOV E,H
		state.e = state.h
	case 0x5d: // MOV E,L
		state.e = state.l
	case 0x5e: // MOV E,M
		state.e = GetMemoryHL(state)
	case 0x5f: // MOV E,A
		state.e = state.a
	case 0x60: // MOV H,B
		state.h = state.b
	case 0x61: // MOV H,C
		state.h = state.c
	case 0x62: // MOV H,D
		state.h = state.d
	case 0x63: // MOV H,E
		state.h = state.e
	case 0x64: // MOV H,H (NOP)
		break
	case 0x65: // MOV H,L
		state.h = state.l
	case 0x66: // MOV H,M
		state.h = GetMemoryHL(state)
	case 0x67: // MOV H,A
		state.h = state.a
	case 0x68: // MOV L,B
		state.l = state.b
	case 0x69: // MOV L,C
		state.l = state.c
	case 0x6a: // MOV L,D
		state.l = state.d
	case 0x6b: // MOV L,E
		state.l = state.e
	case 0x6c: // MOV L,H
		state.l = state.h
	case 0x6d: // MOV L,L (NOP)
		break
	case 0x6e: // MOV L,M
		state.l = GetMemoryHL(state)
	case 0x6f: // MOV L,A
		state.l = state.a
	case 0x70: // MOV M,B
		SetMemoryHL(state, state.b)
	case 0x71: // MOV M,C
		SetMemoryHL(state, state.c)
	case 0x72: // MOV M,D
		SetMemoryHL(state, state.d)
	case 0x73: // MOV M,E
		SetMemoryHL(state, state.e)
	case 0x74: // MOV M,H
		SetMemoryHL(state, state.h)
	case 0x75: // MOV M,L
		SetMemoryHL(state, state.l)
	case 0x76:
		unimplementedInstruction(state)
	case 0x77: // MOV M,A
		SetMemoryHL(state, state.a)
	case 0x78: // MOV A,B
		state.a = state.b
	case 0x79: // MOV A,C
		state.a = state.c
	case 0x7a: // MOV A,D
		state.a = state.d
	case 0x7b: // MOV A,E
		state.a = state.e
	case 0x7c: // MOV A,H
		state.a = state.h
	case 0x7d: // MOV A,L
		state.a = state.l
	case 0x7e: // MOV A,M
		state.a = state.memory[uint16(state.l)<<8|uint16(state.h)]
	case 0x7f: // MOV A,A (NOP)
		break
	case 0x80: // ADD B
		ADD(state, state.b)
	case 0x81: // ADD C
		ADD(state, state.c)
	case 0x82: // ADD D
		ADD(state, state.d)
	case 0x83: // ADD E
		ADD(state, state.e)
	case 0x84: // ADD H
		ADD(state, state.h)
	case 0x85: // ADD L
		ADD(state, state.l)
	case 0x86: // ADD M
		ADD(state, GetMemoryHL(state))
	case 0x87: // ADD A
		ADD(state, state.a)
	case 0x88: // ADC B
		ADC(state, state.b)
	case 0x89: // ADC C
		ADC(state, state.c)
	case 0x8a: // ADC D
		ADC(state, state.d)
	case 0x8b: // ADC E
		ADC(state, state.e)
	case 0x8c: // ADC H
		ADC(state, state.h)
	case 0x8d: // ADC L
		ADC(state, state.l)
	case 0x8e: // ADC M
		ADC(state, GetMemoryHL(state))
	case 0x8f: // ADC A
		ADC(state, state.a)
	case 0x90:
		unimplementedInstruction(state)
	case 0x91:
		unimplementedInstruction(state)
	case 0x92:
		unimplementedInstruction(state)
	case 0x93:
		unimplementedInstruction(state)
	case 0x94:
		unimplementedInstruction(state)
	case 0x95:
		unimplementedInstruction(state)
	case 0x96:
		unimplementedInstruction(state)
	case 0x97:
		unimplementedInstruction(state)
	case 0x98:
		unimplementedInstruction(state)
	case 0x99:
		unimplementedInstruction(state)
	case 0x9a:
		unimplementedInstruction(state)
	case 0x9b:
		unimplementedInstruction(state)
	case 0x9c:
		unimplementedInstruction(state)
	case 0x9d:
		unimplementedInstruction(state)
	case 0x9e:
		unimplementedInstruction(state)
	case 0x9f:
		unimplementedInstruction(state)
	case 0xa0:
		unimplementedInstruction(state)
	case 0xa1:
		unimplementedInstruction(state)
	case 0xa2:
		unimplementedInstruction(state)
	case 0xa3:
		unimplementedInstruction(state)
	case 0xa4:
		unimplementedInstruction(state)
	case 0xa5:
		unimplementedInstruction(state)
	case 0xa6:
		unimplementedInstruction(state)
	case 0xa7:
		unimplementedInstruction(state)
	case 0xa8:
		unimplementedInstruction(state)
	case 0xa9:
		unimplementedInstruction(state)
	case 0xaa:
		unimplementedInstruction(state)
	case 0xab:
		unimplementedInstruction(state)
	case 0xac:
		unimplementedInstruction(state)
	case 0xad:
		unimplementedInstruction(state)
	case 0xae:
		unimplementedInstruction(state)
	case 0xaf:
		unimplementedInstruction(state)
	case 0xb0:
		unimplementedInstruction(state)
	case 0xb1:
		unimplementedInstruction(state)
	case 0xb2:
		unimplementedInstruction(state)
	case 0xb3:
		unimplementedInstruction(state)
	case 0xb4:
		unimplementedInstruction(state)
	case 0xb5:
		unimplementedInstruction(state)
	case 0xb6:
		unimplementedInstruction(state)
	case 0xb7:
		unimplementedInstruction(state)
	case 0xb8:
		unimplementedInstruction(state)
	case 0xb9:
		unimplementedInstruction(state)
	case 0xba:
		unimplementedInstruction(state)
	case 0xbb:
		unimplementedInstruction(state)
	case 0xbc:
		unimplementedInstruction(state)
	case 0xbd:
		unimplementedInstruction(state)
	case 0xbe:
		unimplementedInstruction(state)
	case 0xbf:
		unimplementedInstruction(state)
	case 0xc0:
		unimplementedInstruction(state)
	case 0xc1:
		unimplementedInstruction(state)
	case 0xc2:
		unimplementedInstruction(state)
	case 0xc3:
		unimplementedInstruction(state)
	case 0xc4:
		unimplementedInstruction(state)
	case 0xc5:
		unimplementedInstruction(state)
	case 0xc6:
		unimplementedInstruction(state)
	case 0xc7:
		unimplementedInstruction(state)
	case 0xc8:
		unimplementedInstruction(state)
	case 0xc9:
		unimplementedInstruction(state)
	case 0xca:
		unimplementedInstruction(state)
	case 0xcb:
		unimplementedInstruction(state)
	case 0xcc:
		unimplementedInstruction(state)
	case 0xcd:
		unimplementedInstruction(state)
	case 0xce:
		unimplementedInstruction(state)
	case 0xcf:
		unimplementedInstruction(state)
	case 0xd0:
		unimplementedInstruction(state)
	case 0xd1:
		unimplementedInstruction(state)
	case 0xd2:
		unimplementedInstruction(state)
	case 0xd3:
		unimplementedInstruction(state)
	case 0xd4:
		unimplementedInstruction(state)
	case 0xd5:
		unimplementedInstruction(state)
	case 0xd6:
		unimplementedInstruction(state)
	case 0xd7:
		unimplementedInstruction(state)
	case 0xd8:
		unimplementedInstruction(state)
	case 0xd9:
		unimplementedInstruction(state)
	case 0xda:
		unimplementedInstruction(state)
	case 0xdb:
		unimplementedInstruction(state)
	case 0xdc:
		unimplementedInstruction(state)
	case 0xdd:
		unimplementedInstruction(state)
	case 0xde:
		unimplementedInstruction(state)
	case 0xdf:
		unimplementedInstruction(state)
	case 0xe0:
		unimplementedInstruction(state)
	case 0xe1:
		unimplementedInstruction(state)
	case 0xe2:
		unimplementedInstruction(state)
	case 0xe3:
		unimplementedInstruction(state)
	case 0xe4:
		unimplementedInstruction(state)
	case 0xe5:
		unimplementedInstruction(state)
	case 0xe6:
		unimplementedInstruction(state)
	case 0xe7:
		unimplementedInstruction(state)
	case 0xe8:
		unimplementedInstruction(state)
	case 0xe9:
		unimplementedInstruction(state)
	case 0xea:
		unimplementedInstruction(state)
	case 0xeb:
		unimplementedInstruction(state)
	case 0xec:
		unimplementedInstruction(state)
	case 0xed:
		unimplementedInstruction(state)
	case 0xee:
		unimplementedInstruction(state)
	case 0xef:
		unimplementedInstruction(state)
	case 0xf0:
		unimplementedInstruction(state)
	case 0xf1:
		unimplementedInstruction(state)
	case 0xf2:
		unimplementedInstruction(state)
	case 0xf3:
		unimplementedInstruction(state)
	case 0xf4:
		unimplementedInstruction(state)
	case 0xf5:
		unimplementedInstruction(state)
	case 0xf6:
		unimplementedInstruction(state)
	case 0xf7:
		unimplementedInstruction(state)
	case 0xf8:
		unimplementedInstruction(state)
	case 0xf9:
		unimplementedInstruction(state)
	case 0xfa:
		unimplementedInstruction(state)
	case 0xfb:
		unimplementedInstruction(state)
	case 0xfc:
		unimplementedInstruction(state)
	case 0xfd:
		unimplementedInstruction(state)
	case 0xfe:
		unimplementedInstruction(state)
	case 0xff:
		unimplementedInstruction(state)
	}
	state.pc++
}
