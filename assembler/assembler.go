package assembler

import (
	"bytes"
	"fmt"
	s "strings"
)

// Assemble assembles assembly code into bytes
func Assemble(code string) (assembly []byte) {
	lines := s.Split(code, "\n")
	for _, line := range lines {
		tokens := s.Split(line, " ")
		instruction := tokens[0]
		assembly = append(assembly, InstructionSetByOpcode[instruction].byteN)
	}
	return assembly
}

// Disassemble disassembles bytes into assembly code
func Disassemble(assembly []byte) string {
	var code bytes.Buffer
	for i := 0; i < len(assembly); i++ {
		byteN := assembly[i]
		inst := InstructionSetByByte[byteN]
		code.WriteString(inst.opcode)
		if inst.paramsN != 0 {
			code.WriteString(" ")
			if inst.paramsN == 1 {
				code.WriteString(fmt.Sprintf(inst.paramsTemplate, assembly[i+1]))
			} else if inst.paramsN == 2 {
				code.WriteString(fmt.Sprintf(inst.paramsTemplate, assembly[i+1], assembly[i+2]))
			}
			i += inst.paramsN
		}
		code.WriteString("\n")
	}
	return code.String()
}
