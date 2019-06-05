package assembler

import (
	"fmt"
	"testing"
)

var exampleAssembly = []byte{0x01, 0xFA, 0xCC, 0x11, 0xBB, 0xCC, 0x02}

// TODO...
func TestDisassemble(t *testing.T) {
	result := Disassemble(exampleAssembly)
	fmt.Println(result)
}
