package assembler

import (
	"fmt"
	"testing"
)

var exampleCode = `ONE
TWO
THR
TWO
THR
ONE`

var exampleAssembly = []byte{0x01, 0x02, 0x03, 0x01, 0x01, 0x03, 0x03, 0x02, 0x02, 0x04, 0x11, 0x05, 0xAA, 0xBB}

func TestAssemble(t *testing.T) {
	result := Assemble(exampleCode)
	fmt.Println(result)
}

func TestDisassemble(t *testing.T) {
	result := Disassemble(exampleAssembly)
	fmt.Println(result)
}
