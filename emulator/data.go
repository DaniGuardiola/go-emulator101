package emulator

// GetHL retrieves the address stored in the HL register pair
func GetHL(state *State) uint16 {
	return uint16(state.l)<<8 | uint16(state.h)
}

// GetMemoryHL retrieves the byte from memory at the address stored in the HL register pair
func GetMemoryHL(state *State) byte {
	return state.memory[GetHL(state)]
}

// SetMemoryHL sets a value in memory at the address stored in the HL register pair
func SetMemoryHL(state *State, value byte) {
	state.memory[GetHL(state)] = value
}
