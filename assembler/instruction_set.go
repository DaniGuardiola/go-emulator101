package assembler

type instruction struct {
	byteN          byte
	paramsN        int
	opcode         string
	paramsTemplate string
}

// InstructionSetByOpcode contains all CPU instructions indexed by opcode
var InstructionSetByOpcode = map[string]instruction{
	"ONE": {0x01, 0, "ONE", ""},
	"TWO": {0x02, 0, "TWO", ""},
	"THR": {0x03, 0, "THR", ""},
	"PA1": {0x04, 1, "PA1", "B%x"},
	"PA2": {0x05, 2, "PA2", "B%x,%x"},
}

// InstructionSetByByte contains all CPU instructions indexed by byte
var InstructionSetByByte = map[byte]instruction{
	0x01: InstructionSetByOpcode["ONE"],
	0x02: InstructionSetByOpcode["TWO"],
	0x03: InstructionSetByOpcode["THR"],
	0x04: InstructionSetByOpcode["PA1"],
	0x05: InstructionSetByOpcode["PA2"],
}
