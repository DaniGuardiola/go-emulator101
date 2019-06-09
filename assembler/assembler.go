package assembler

import (
	"bytes"
	"fmt"
)

// Disassemble converts machine code into assembly code
func Disassemble(mc []byte) string {
	var code bytes.Buffer
	var skip int
	for i := 0; i < len(mc); i += 1 + skip {
		if i != 0 {
			code.WriteString("\n")
		}
		code.WriteString(fmt.Sprintf("%04x ", i))
		skip = 0
		switch mc[i] {
		case 0x00:
			code.WriteString("NOP")
		case 0x01:
			code.WriteString(fmt.Sprintf("LXI    B,#$%02x%02x", mc[i+2], mc[i+1]))
			skip = 2
		case 0x02:
			code.WriteString("STAX   B")
		case 0x03:
			code.WriteString("INX    B")
		case 0x04:
			code.WriteString("INR    B")
		case 0x05:
			code.WriteString("DCR    B")
		case 0x06:
			code.WriteString(fmt.Sprintf("MVI    B,#$%02x", mc[i+1]))
			skip = 1
		case 0x07:
			code.WriteString("RLC")
		case 0x09:
			code.WriteString("DAD    B")
		case 0x0A:
			code.WriteString("LDAX   B")
		case 0x0B:
			code.WriteString("DCX    B")
		case 0x0C:
			code.WriteString("INR    C")
		case 0x0D:
			code.WriteString("DCR    C")
		case 0x0E:
			code.WriteString(fmt.Sprintf("MVI    C,#$%02x", mc[i+1]))
			skip = 1
		case 0x0F:
			code.WriteString("RRC")
		case 0x11:
			code.WriteString(fmt.Sprintf("LXI    D,#$%02x%02x", mc[i+2], mc[i+1]))
			skip = 2
		case 0x12:
			code.WriteString("STAX   D")
		case 0x13:
			code.WriteString("INX    D")
		case 0x14:
			code.WriteString("INR    D")
		case 0x15:
			code.WriteString("DCR    D")
		case 0x16:
			code.WriteString(fmt.Sprintf("MVI    D,#$%02x", mc[i+1]))
			skip = 1
		case 0x17:
			code.WriteString("RAL")
		case 0x19:
			code.WriteString("DAD    D")
		case 0x1A:
			code.WriteString("LDAX   D")
		case 0x1B:
			code.WriteString("DCX    D")
		case 0x1C:
			code.WriteString("INR    E")
		case 0x1D:
			code.WriteString("DCR    E")
		case 0x1E:
			code.WriteString(fmt.Sprintf("MVI    E,#$%02x", mc[i+1]))
			skip = 1
		case 0x1F:
			code.WriteString("RAR")
		case 0x21:
			code.WriteString(fmt.Sprintf("LXI    H,#$%02x%02x", mc[i+2], mc[i+1]))
			skip = 2
		case 0x22:
			code.WriteString(fmt.Sprintf("SHLD   $%02x%02x", mc[i+2], mc[i+1]))
			skip = 2
		case 0x23:
			code.WriteString("INX    H")
		case 0x24:
			code.WriteString("INR    H")
		case 0x25:
			code.WriteString("DCR    H")
		case 0x26:
			code.WriteString(fmt.Sprintf("MVI    H,#$%02x", mc[i+1]))
			skip = 1
		case 0x27:
			code.WriteString("DAA")
		case 0x29:
			code.WriteString("DAD    H")
		case 0x2A:
			code.WriteString(fmt.Sprintf("LHLD   $%02x%02x", mc[i+2], mc[i+1]))
			skip = 2
		case 0x2B:
			code.WriteString("DCX    H")
		case 0x2C:
			code.WriteString("INR    L")
		case 0x2D:
			code.WriteString("DCR    L")
		case 0x2E:
			code.WriteString(fmt.Sprintf("MVI    L,#$%02x", mc[i+1]))
			skip = 1
		case 0x2F:
			code.WriteString("CMA")
		case 0x31:
			code.WriteString(fmt.Sprintf("LXI    SP,#$%02x%02x", mc[i+2], mc[i+1]))
			skip = 2
		case 0x32:
			code.WriteString(fmt.Sprintf("STA    $%02x%02x", mc[i+2], mc[i+1]))
			skip = 2
		case 0x33:
			code.WriteString("INX    SP")
		case 0x34:
			code.WriteString("INR    M")
		case 0x35:
			code.WriteString("DCR    M")
		case 0x36:
			code.WriteString(fmt.Sprintf("MVI    M,#$%02x", mc[i+1]))
			skip = 1
		case 0x37:
			code.WriteString("STC")
		case 0x39:
			code.WriteString("DAD    SP")
		case 0x3A:
			code.WriteString(fmt.Sprintf("LDA    $%02x%02x", mc[i+2], mc[i+1]))
			skip = 2
		case 0x3B:
			code.WriteString("DCX    SP")
		case 0x3C:
			code.WriteString("INR    A")
		case 0x3D:
			code.WriteString("DCR    A")
		case 0x3E:
			code.WriteString(fmt.Sprintf("MVI    A,#$%02x", mc[i+1]))
			skip = 1
		case 0x3F:
			code.WriteString("CMC")
		case 0x40:
			code.WriteString("MOV    B,B")
		case 0x41:
			code.WriteString("MOV    B,C")
		case 0x42:
			code.WriteString("MOV    B,D")
		case 0x43:
			code.WriteString("MOV    B,E")
		case 0x44:
			code.WriteString("MOV    B,H")
		case 0x45:
			code.WriteString("MOV    B,L")
		case 0x46:
			code.WriteString("MOV    B,M")
		case 0x47:
			code.WriteString("MOV    B,A")
		case 0x48:
			code.WriteString("MOV    C,B")
		case 0x49:
			code.WriteString("MOV    C,C")
		case 0x4A:
			code.WriteString("MOV    C,D")
		case 0x4B:
			code.WriteString("MOV    C,E")
		case 0x4C:
			code.WriteString("MOV    C,H")
		case 0x4D:
			code.WriteString("MOV    C,L")
		case 0x4E:
			code.WriteString("MOV    C,M")
		case 0x4F:
			code.WriteString("MOV    C,A")
		case 0x50:
			code.WriteString("MOV    D,B")
		case 0x51:
			code.WriteString("MOV    D,C")
		case 0x52:
			code.WriteString("MOV    D,D")
		case 0x53:
			code.WriteString("MOV    D,E")
		case 0x54:
			code.WriteString("MOV    D,H")
		case 0x55:
			code.WriteString("MOV    D,L")
		case 0x56:
			code.WriteString("MOV    D,M")
		case 0x57:
			code.WriteString("MOV    D,A")
		case 0x58:
			code.WriteString("MOV    E,B")
		case 0x59:
			code.WriteString("MOV    E,C")
		case 0x5A:
			code.WriteString("MOV    E,D")
		case 0x5B:
			code.WriteString("MOV    E,E")
		case 0x5C:
			code.WriteString("MOV    E,H")
		case 0x5D:
			code.WriteString("MOV    E,L")
		case 0x5E:
			code.WriteString("MOV    E,M")
		case 0x5F:
			code.WriteString("MOV    E,A")
		case 0x60:
			code.WriteString("MOV    H,B")
		case 0x61:
			code.WriteString("MOV    H,C")
		case 0x62:
			code.WriteString("MOV    H,D")
		case 0x63:
			code.WriteString("MOV    H,E")
		case 0x64:
			code.WriteString("MOV    H,H")
		case 0x65:
			code.WriteString("MOV    H,L")
		case 0x66:
			code.WriteString("MOV    H,M")
		case 0x67:
			code.WriteString("MOV    H,A")
		case 0x68:
			code.WriteString("MOV    L,B")
		case 0x69:
			code.WriteString("MOV    L,C")
		case 0x6A:
			code.WriteString("MOV    L,D")
		case 0x6B:
			code.WriteString("MOV    L,E")
		case 0x6C:
			code.WriteString("MOV    L,H")
		case 0x6D:
			code.WriteString("MOV    L,L")
		case 0x6E:
			code.WriteString("MOV    L,M")
		case 0x6F:
			code.WriteString("MOV    L,A")
		case 0x70:
			code.WriteString("MOV    M,B")
		case 0x71:
			code.WriteString("MOV    M,C")
		case 0x72:
			code.WriteString("MOV    M,D")
		case 0x73:
			code.WriteString("MOV    M,E")
		case 0x74:
			code.WriteString("MOV    M,H")
		case 0x75:
			code.WriteString("MOV    M,L")
		case 0x76:
			code.WriteString("HLT")
		case 0x77:
			code.WriteString("MOV    M,A")
		case 0x78:
			code.WriteString("MOV    A,B")
		case 0x79:
			code.WriteString("MOV    A,C")
		case 0x7A:
			code.WriteString("MOV    A,D")
		case 0x7B:
			code.WriteString("MOV    A,E")
		case 0x7C:
			code.WriteString("MOV    A,H")
		case 0x7D:
			code.WriteString("MOV    A,L")
		case 0x7E:
			code.WriteString("MOV    A,M")
		case 0x7F:
			code.WriteString("MOV    A,A")
		case 0x80:
			code.WriteString("ADD    B")
		case 0x81:
			code.WriteString("ADD    C")
		case 0x82:
			code.WriteString("ADD    D")
		case 0x83:
			code.WriteString("ADD    E")
		case 0x84:
			code.WriteString("ADD    H")
		case 0x85:
			code.WriteString("ADD    L")
		case 0x86:
			code.WriteString("ADD    M")
		case 0x87:
			code.WriteString("ADD    A")
		case 0x88:
			code.WriteString("ADC    B")
		case 0x89:
			code.WriteString("ADC    C")
		case 0x8A:
			code.WriteString("ADC    D")
		case 0x8B:
			code.WriteString("ADC    E")
		case 0x8C:
			code.WriteString("ADC    H")
		case 0x8D:
			code.WriteString("ADC    L")
		case 0x8E:
			code.WriteString("ADC    M")
		case 0x8F:
			code.WriteString("ADC    A")
		case 0x90:
			code.WriteString("SUB    B")
		case 0x91:
			code.WriteString("SUB    C")
		case 0x92:
			code.WriteString("SUB    D")
		case 0x93:
			code.WriteString("SUB    E")
		case 0x94:
			code.WriteString("SUB    H")
		case 0x95:
			code.WriteString("SUB    L")
		case 0x96:
			code.WriteString("SUB    M")
		case 0x97:
			code.WriteString("SUB    A")
		case 0x98:
			code.WriteString("SBB    B")
		case 0x99:
			code.WriteString("SBB    C")
		case 0x9A:
			code.WriteString("SBB    D")
		case 0x9B:
			code.WriteString("SBB    E")
		case 0x9C:
			code.WriteString("SBB    H")
		case 0x9D:
			code.WriteString("SBB    L")
		case 0x9E:
			code.WriteString("SBB    M")
		case 0x9F:
			code.WriteString("SBB    A")
		case 0xA0:
			code.WriteString("ANA    B")
		case 0xA1:
			code.WriteString("ANA    C")
		case 0xA2:
			code.WriteString("ANA    D")
		case 0xA3:
			code.WriteString("ANA    E")
		case 0xA4:
			code.WriteString("ANA    H")
		case 0xA5:
			code.WriteString("ANA    L")
		case 0xA6:
			code.WriteString("ANA    M")
		case 0xA7:
			code.WriteString("ANA    A")
		case 0xA8:
			code.WriteString("XRA    B")
		case 0xA9:
			code.WriteString("XRA    C")
		case 0xAA:
			code.WriteString("XRA    D")
		case 0xAB:
			code.WriteString("XRA    E")
		case 0xAC:
			code.WriteString("XRA    H")
		case 0xAD:
			code.WriteString("XRA    L")
		case 0xAE:
			code.WriteString("XRA    M")
		case 0xAF:
			code.WriteString("XRA    A")
		case 0xB0:
			code.WriteString("ORA    B")
		case 0xB1:
			code.WriteString("ORA    C")
		case 0xB2:
			code.WriteString("ORA    D")
		case 0xB3:
			code.WriteString("ORA    E")
		case 0xB4:
			code.WriteString("ORA    H")
		case 0xB5:
			code.WriteString("ORA    L")
		case 0xB6:
			code.WriteString("ORA    M")
		case 0xB7:
			code.WriteString("ORA    A")
		case 0xB8:
			code.WriteString("CMP    B")
		case 0xB9:
			code.WriteString("CMP    C")
		case 0xBA:
			code.WriteString("CMP    D")
		case 0xBB:
			code.WriteString("CMP    E")
		case 0xBC:
			code.WriteString("CMP    H")
		case 0xBD:
			code.WriteString("CMP    L")
		case 0xBE:
			code.WriteString("CMP    M")
		case 0xBF:
			code.WriteString("CMP    A")
		case 0xC0:
			code.WriteString("RNZ")
		case 0xC1:
			code.WriteString("POP    B")
		case 0xC2:
			code.WriteString(fmt.Sprintf("JNZ    $%02x%02x", mc[i+2], mc[i+1]))
			skip = 2
		case 0xC3:
			code.WriteString(fmt.Sprintf("JMP    $%02x%02x", mc[i+2], mc[i+1]))
			skip = 2
		case 0xC4:
			code.WriteString(fmt.Sprintf("CNZ    $%02x%02x", mc[i+2], mc[i+1]))
			skip = 2
		case 0xC5:
			code.WriteString("PUSH   B")
		case 0xC6:
			code.WriteString(fmt.Sprintf("ADI    #$%02x", mc[i+1]))
			skip = 1
		case 0xC7:
			code.WriteString("RST    0")
		case 0xC8:
			code.WriteString("RZ")
		case 0xC9:
			code.WriteString("RET")
		case 0xCA:
			code.WriteString(fmt.Sprintf("JZ     $%02x%02x", mc[i+2], mc[i+1]))
			skip = 2
		case 0xCC:
			code.WriteString(fmt.Sprintf("CZ     $%02x%02x", mc[i+2], mc[i+1]))
			skip = 2
		case 0xCD:
			code.WriteString(fmt.Sprintf("CALL   $%02x%02x", mc[i+2], mc[i+1]))
			skip = 2
		case 0xCE:
			code.WriteString(fmt.Sprintf("ACI    #$%02x", mc[i+1]))
			skip = 1
		case 0xCF:
			code.WriteString("RST    1")
		case 0xD0:
			code.WriteString("RNC")
		case 0xD1:
			code.WriteString("POP    D")
		case 0xD2:
			code.WriteString(fmt.Sprintf("JNC    $%02x%02x", mc[i+2], mc[i+1]))
			skip = 2
		case 0xD3:
			code.WriteString(fmt.Sprintf("OUT    #$%02x", mc[i+1]))
			skip = 1
		case 0xD4:
			code.WriteString(fmt.Sprintf("CNC    $%02x%02x", mc[i+2], mc[i+1]))
			skip = 2
		case 0xD5:
			code.WriteString("PUSH   D")
		case 0xD6:
			code.WriteString(fmt.Sprintf("SUI    #$%02x", mc[i+1]))
			skip = 1
		case 0xD7:
			code.WriteString("RST    2")
		case 0xD8:
			code.WriteString("RC")
		case 0xDA:
			code.WriteString(fmt.Sprintf("JC     $%02x%02x", mc[i+2], mc[i+1]))
			skip = 2
		case 0xDB:
			code.WriteString(fmt.Sprintf("IN     #$%02x", mc[i+1]))
			skip = 1
		case 0xDC:
			code.WriteString(fmt.Sprintf("CC     $%02x%02x", mc[i+2], mc[i+1]))
			skip = 2
		case 0xDE:
			code.WriteString(fmt.Sprintf("SBI    #$%02x", mc[i+1]))
			skip = 1
		case 0xDF:
			code.WriteString("RST    3")
		case 0xE0:
			code.WriteString("RPO")
		case 0xE1:
			code.WriteString("POP    H")
		case 0xE2:
			code.WriteString(fmt.Sprintf("JPO    $%02x%02x", mc[i+2], mc[i+1]))
			skip = 2
		case 0xE3:
			code.WriteString("XTHL")
		case 0xE4:
			code.WriteString(fmt.Sprintf("CPO    $%02x%02x", mc[i+2], mc[i+1]))
			skip = 2
		case 0xE5:
			code.WriteString("PUSH   H")
		case 0xE6:
			code.WriteString(fmt.Sprintf("ANI    #$%02x", mc[i+1]))
			skip = 1
		case 0xE7:
			code.WriteString("RST    4")
		case 0xE8:
			code.WriteString("RPE")
		case 0xE9:
			code.WriteString("PCHL")
		case 0xEA:
			code.WriteString(fmt.Sprintf("JPE    $%02x%02x", mc[i+2], mc[i+1]))
			skip = 2
		case 0xEB:
			code.WriteString("XCHG")
		case 0xEC:
			code.WriteString(fmt.Sprintf("CPE    $%02x%02x", mc[i+2], mc[i+1]))
			skip = 2
		case 0xEE:
			code.WriteString(fmt.Sprintf("XRI    #$%02x", mc[i+1]))
			skip = 1
		case 0xEF:
			code.WriteString("RST    5")
		case 0xF0:
			code.WriteString("RP")
		case 0xF1:
			code.WriteString("POP    PSW")
		case 0xF2:
			code.WriteString(fmt.Sprintf("JP     $%02x%02x", mc[i+2], mc[i+1]))
			skip = 2
		case 0xF3:
			code.WriteString("DI")
		case 0xF4:
			code.WriteString(fmt.Sprintf("CP     $%02x%02x", mc[i+2], mc[i+1]))
			skip = 2
		case 0xF5:
			code.WriteString("PUSH   PSW")
		case 0xF6:
			code.WriteString(fmt.Sprintf("ORI    #$%02x", mc[i+1]))
			skip = 1
		case 0xF7:
			code.WriteString("RST    6")
		case 0xF8:
			code.WriteString("RM")
		case 0xF9:
			code.WriteString("SPHL")
		case 0xFA:
			code.WriteString(fmt.Sprintf("JM     $%02x%02x", mc[i+2], mc[i+1]))
			skip = 2
		case 0xFB:
			code.WriteString("EI")
		case 0xFC:
			code.WriteString(fmt.Sprintf("CM     $%02x%02x", mc[i+2], mc[i+1]))
			skip = 2
		case 0xFE:
			code.WriteString(fmt.Sprintf("CPI    #$%02x", mc[i+1]))
			skip = 1
		case 0xFF:
			code.WriteString("RST    7")
		default:
			code.WriteString("NOP")
		}
	}
	return code.String()
}
