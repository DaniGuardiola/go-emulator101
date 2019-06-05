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
		skip = 0
		switch mc[i] {
		case 0x00:
			code.WriteString(fmt.Sprintf("NOP"))
		case 0x01:
			code.WriteString(fmt.Sprintf("LXI    B,$%02X%02X", mc[i+1], mc[i+2]))
			skip = 2
		case 0x02:
			code.WriteString(fmt.Sprintf("STAX   B"))
		case 0x03:
			code.WriteString(fmt.Sprintf("INX    B"))
		case 0x04:
			code.WriteString(fmt.Sprintf("INR    B"))
		case 0x05:
			code.WriteString(fmt.Sprintf("DCR    B"))
		case 0x06:
			code.WriteString(fmt.Sprintf("MVI    B,$%02X", mc[i+1]))
			skip = 1
		case 0x07:
			code.WriteString(fmt.Sprintf("RLC"))
		case 0x09:
			code.WriteString(fmt.Sprintf("DAD    B"))
		case 0x0A:
			code.WriteString(fmt.Sprintf("LDAX   B"))
		case 0x0B:
			code.WriteString(fmt.Sprintf("DCX    B"))
		case 0x0C:
			code.WriteString(fmt.Sprintf("INR    C"))
		case 0x0D:
			code.WriteString(fmt.Sprintf("DCR    C"))
		case 0x0E:
			code.WriteString(fmt.Sprintf("MVI    C,$%02X", mc[i+1]))
			skip = 1
		case 0x0F:
			code.WriteString(fmt.Sprintf("RRC"))
		case 0x11:
			code.WriteString(fmt.Sprintf("LXI    D,$%02X%02X", mc[i+1], mc[i+2]))
			skip = 2
		case 0x12:
			code.WriteString(fmt.Sprintf("STAX   D"))
		case 0x13:
			code.WriteString(fmt.Sprintf("INX    D"))
		case 0x14:
			code.WriteString(fmt.Sprintf("INR    D"))
		case 0x15:
			code.WriteString(fmt.Sprintf("DCR    D"))
		case 0x16:
			code.WriteString(fmt.Sprintf("MVI    D,$%02X", mc[i+1]))
			skip = 1
		case 0x17:
			code.WriteString(fmt.Sprintf("RAL"))
		case 0x19:
			code.WriteString(fmt.Sprintf("DAD    D"))
		case 0x1A:
			code.WriteString(fmt.Sprintf("LDAX   D"))
		case 0x1B:
			code.WriteString(fmt.Sprintf("DCX    D"))
		case 0x1C:
			code.WriteString(fmt.Sprintf("INR    E"))
		case 0x1D:
			code.WriteString(fmt.Sprintf("DCR    E"))
		case 0x1E:
			code.WriteString(fmt.Sprintf("MVI    E,$%02X", mc[i+1]))
			skip = 1
		case 0x1F:
			code.WriteString(fmt.Sprintf("RAR"))
		case 0x21:
			code.WriteString(fmt.Sprintf("LXI    H,$%02X%02X", mc[i+1], mc[i+2]))
			skip = 2
		case 0x22:
			code.WriteString(fmt.Sprintf("SHLD   $%02X%02X", mc[i+1], mc[i+2]))
			skip = 2
		case 0x23:
			code.WriteString(fmt.Sprintf("INX    H"))
		case 0x24:
			code.WriteString(fmt.Sprintf("INR    H"))
		case 0x25:
			code.WriteString(fmt.Sprintf("DCR    H"))
		case 0x26:
			code.WriteString(fmt.Sprintf("MVI    H,$%02X", mc[i+1]))
			skip = 1
		case 0x27:
			code.WriteString(fmt.Sprintf("DAA"))
		case 0x29:
			code.WriteString(fmt.Sprintf("DAD    H"))
		case 0x2A:
			code.WriteString(fmt.Sprintf("LHLD   $%02X%02X", mc[i+1], mc[i+2]))
			skip = 2
		case 0x2B:
			code.WriteString(fmt.Sprintf("DCX    H"))
		case 0x2C:
			code.WriteString(fmt.Sprintf("INR    L"))
		case 0x2D:
			code.WriteString(fmt.Sprintf("DCR    L"))
		case 0x2E:
			code.WriteString(fmt.Sprintf("MVI    L,$%02X", mc[i+1]))
			skip = 1
		case 0x2F:
			code.WriteString(fmt.Sprintf("CMA"))
		case 0x31:
			code.WriteString(fmt.Sprintf("LXI    SP,$%02X%02X", mc[i+1], mc[i+2]))
			skip = 2
		case 0x32:
			code.WriteString(fmt.Sprintf("STA    $%02X%02X", mc[i+1], mc[i+2]))
			skip = 2
		case 0x33:
			code.WriteString(fmt.Sprintf("INX    SP"))
		case 0x34:
			code.WriteString(fmt.Sprintf("INR    M"))
		case 0x35:
			code.WriteString(fmt.Sprintf("DCR    M"))
		case 0x36:
			code.WriteString(fmt.Sprintf("MVI    M,$%02X", mc[i+1]))
			skip = 1
		case 0x37:
			code.WriteString(fmt.Sprintf("STC"))
		case 0x39:
			code.WriteString(fmt.Sprintf("DAD    SP"))
		case 0x3A:
			code.WriteString(fmt.Sprintf("LDA    $%02X%02X", mc[i+1], mc[i+2]))
			skip = 2
		case 0x3B:
			code.WriteString(fmt.Sprintf("DCX    SP"))
		case 0x3C:
			code.WriteString(fmt.Sprintf("INR    A"))
		case 0x3D:
			code.WriteString(fmt.Sprintf("DCR    A"))
		case 0x3E:
			code.WriteString(fmt.Sprintf("MVI    A,$%02X", mc[i+1]))
			skip = 1
		case 0x3F:
			code.WriteString(fmt.Sprintf("CMC"))
		case 0x40:
			code.WriteString(fmt.Sprintf("MOV    B,B"))
		case 0x41:
			code.WriteString(fmt.Sprintf("MOV    B,C"))
		case 0x42:
			code.WriteString(fmt.Sprintf("MOV    B,D"))
		case 0x43:
			code.WriteString(fmt.Sprintf("MOV    B,E"))
		case 0x44:
			code.WriteString(fmt.Sprintf("MOV    B,H"))
		case 0x45:
			code.WriteString(fmt.Sprintf("MOV    B,L"))
		case 0x46:
			code.WriteString(fmt.Sprintf("MOV    B,M"))
		case 0x47:
			code.WriteString(fmt.Sprintf("MOV    B,A"))
		case 0x48:
			code.WriteString(fmt.Sprintf("MOV    C,B"))
		case 0x49:
			code.WriteString(fmt.Sprintf("MOV    C,C"))
		case 0x4A:
			code.WriteString(fmt.Sprintf("MOV    C,D"))
		case 0x4B:
			code.WriteString(fmt.Sprintf("MOV    C,E"))
		case 0x4C:
			code.WriteString(fmt.Sprintf("MOV    C,H"))
		case 0x4D:
			code.WriteString(fmt.Sprintf("MOV    C,L"))
		case 0x4E:
			code.WriteString(fmt.Sprintf("MOV    C,M"))
		case 0x4F:
			code.WriteString(fmt.Sprintf("MOV    C,A"))
		case 0x50:
			code.WriteString(fmt.Sprintf("MOV    D,B"))
		case 0x51:
			code.WriteString(fmt.Sprintf("MOV    D,C"))
		case 0x52:
			code.WriteString(fmt.Sprintf("MOV    D,D"))
		case 0x53:
			code.WriteString(fmt.Sprintf("MOV    D,E"))
		case 0x54:
			code.WriteString(fmt.Sprintf("MOV    D,H"))
		case 0x55:
			code.WriteString(fmt.Sprintf("MOV    D,L"))
		case 0x56:
			code.WriteString(fmt.Sprintf("MOV    D,M"))
		case 0x57:
			code.WriteString(fmt.Sprintf("MOV    D,A"))
		case 0x58:
			code.WriteString(fmt.Sprintf("MOV    E,B"))
		case 0x59:
			code.WriteString(fmt.Sprintf("MOV    E,C"))
		case 0x5A:
			code.WriteString(fmt.Sprintf("MOV    E,D"))
		case 0x5B:
			code.WriteString(fmt.Sprintf("MOV    E,E"))
		case 0x5C:
			code.WriteString(fmt.Sprintf("MOV    E,H"))
		case 0x5D:
			code.WriteString(fmt.Sprintf("MOV    E,L"))
		case 0x5E:
			code.WriteString(fmt.Sprintf("MOV    E,M"))
		case 0x5F:
			code.WriteString(fmt.Sprintf("MOV    E,A"))
		case 0x60:
			code.WriteString(fmt.Sprintf("MOV    H,B"))
		case 0x61:
			code.WriteString(fmt.Sprintf("MOV    H,C"))
		case 0x62:
			code.WriteString(fmt.Sprintf("MOV    H,D"))
		case 0x63:
			code.WriteString(fmt.Sprintf("MOV    H,E"))
		case 0x64:
			code.WriteString(fmt.Sprintf("MOV    H,H"))
		case 0x65:
			code.WriteString(fmt.Sprintf("MOV    H,L"))
		case 0x66:
			code.WriteString(fmt.Sprintf("MOV    H,M"))
		case 0x67:
			code.WriteString(fmt.Sprintf("MOV    H,A"))
		case 0x68:
			code.WriteString(fmt.Sprintf("MOV    L,B"))
		case 0x69:
			code.WriteString(fmt.Sprintf("MOV    L,C"))
		case 0x6A:
			code.WriteString(fmt.Sprintf("MOV    L,D"))
		case 0x6B:
			code.WriteString(fmt.Sprintf("MOV    L,E"))
		case 0x6C:
			code.WriteString(fmt.Sprintf("MOV    L,H"))
		case 0x6D:
			code.WriteString(fmt.Sprintf("MOV    L,L"))
		case 0x6E:
			code.WriteString(fmt.Sprintf("MOV    L,M"))
		case 0x6F:
			code.WriteString(fmt.Sprintf("MOV    L,A"))
		case 0x70:
			code.WriteString(fmt.Sprintf("MOV    M,B"))
		case 0x71:
			code.WriteString(fmt.Sprintf("MOV    M,C"))
		case 0x72:
			code.WriteString(fmt.Sprintf("MOV    M,D"))
		case 0x73:
			code.WriteString(fmt.Sprintf("MOV    M,E"))
		case 0x74:
			code.WriteString(fmt.Sprintf("MOV    M,H"))
		case 0x75:
			code.WriteString(fmt.Sprintf("MOV    M,L"))
		case 0x76:
			code.WriteString(fmt.Sprintf("MOV    M,M"))
		case 0x77:
			code.WriteString(fmt.Sprintf("MOV    M,A"))
		case 0x78:
			code.WriteString(fmt.Sprintf("MOV    A,B"))
		case 0x79:
			code.WriteString(fmt.Sprintf("MOV    A,C"))
		case 0x7A:
			code.WriteString(fmt.Sprintf("MOV    A,D"))
		case 0x7B:
			code.WriteString(fmt.Sprintf("MOV    A,E"))
		case 0x7C:
			code.WriteString(fmt.Sprintf("MOV    A,H"))
		case 0x7D:
			code.WriteString(fmt.Sprintf("MOV    A,L"))
		case 0x7E:
			code.WriteString(fmt.Sprintf("MOV    A,M"))
		case 0x7F:
			code.WriteString(fmt.Sprintf("MOV    A,A"))
		case 0x80:
			code.WriteString(fmt.Sprintf("ADD    B"))
		case 0x81:
			code.WriteString(fmt.Sprintf("ADD    C"))
		case 0x82:
			code.WriteString(fmt.Sprintf("ADD    D"))
		case 0x83:
			code.WriteString(fmt.Sprintf("ADD    E"))
		case 0x84:
			code.WriteString(fmt.Sprintf("ADD    H"))
		case 0x85:
			code.WriteString(fmt.Sprintf("ADD    L"))
		case 0x86:
			code.WriteString(fmt.Sprintf("ADD    M"))
		case 0x87:
			code.WriteString(fmt.Sprintf("ADD    A"))
		case 0x88:
			code.WriteString(fmt.Sprintf("ADC    B"))
		case 0x89:
			code.WriteString(fmt.Sprintf("ADC    C"))
		case 0x8A:
			code.WriteString(fmt.Sprintf("ADC    D"))
		case 0x8B:
			code.WriteString(fmt.Sprintf("ADC    E"))
		case 0x8C:
			code.WriteString(fmt.Sprintf("ADC    H"))
		case 0x8D:
			code.WriteString(fmt.Sprintf("ADC    L"))
		case 0x8E:
			code.WriteString(fmt.Sprintf("ADC    M"))
		case 0x8F:
			code.WriteString(fmt.Sprintf("ADC    A"))
		case 0x90:
			code.WriteString(fmt.Sprintf("SUB    B"))
		case 0x91:
			code.WriteString(fmt.Sprintf("SUB    C"))
		case 0x92:
			code.WriteString(fmt.Sprintf("SUB    D"))
		case 0x93:
			code.WriteString(fmt.Sprintf("SUB    E"))
		case 0x94:
			code.WriteString(fmt.Sprintf("SUB    H"))
		case 0x95:
			code.WriteString(fmt.Sprintf("SUB    L"))
		case 0x96:
			code.WriteString(fmt.Sprintf("SUB    M"))
		case 0x97:
			code.WriteString(fmt.Sprintf("SUB    A"))
		case 0x98:
			code.WriteString(fmt.Sprintf("SBB    B"))
		case 0x99:
			code.WriteString(fmt.Sprintf("SBB    C"))
		case 0x9A:
			code.WriteString(fmt.Sprintf("SBB    D"))
		case 0x9B:
			code.WriteString(fmt.Sprintf("SBB    E"))
		case 0x9C:
			code.WriteString(fmt.Sprintf("SBB    H"))
		case 0x9D:
			code.WriteString(fmt.Sprintf("SBB    L"))
		case 0x9E:
			code.WriteString(fmt.Sprintf("SBB    M"))
		case 0x9F:
			code.WriteString(fmt.Sprintf("SBB    A"))
		case 0xA0:
			code.WriteString(fmt.Sprintf("ANA    B"))
		case 0xA1:
			code.WriteString(fmt.Sprintf("ANA    C"))
		case 0xA2:
			code.WriteString(fmt.Sprintf("ANA    D"))
		case 0xA3:
			code.WriteString(fmt.Sprintf("ANA    E"))
		case 0xA4:
			code.WriteString(fmt.Sprintf("ANA    H"))
		case 0xA5:
			code.WriteString(fmt.Sprintf("ANA    L"))
		case 0xA6:
			code.WriteString(fmt.Sprintf("ANA    M"))
		case 0xA7:
			code.WriteString(fmt.Sprintf("ANA    A"))
		case 0xA8:
			code.WriteString(fmt.Sprintf("XRA    B"))
		case 0xA9:
			code.WriteString(fmt.Sprintf("XRA    C"))
		case 0xAA:
			code.WriteString(fmt.Sprintf("XRA    D"))
		case 0xAB:
			code.WriteString(fmt.Sprintf("XRA    E"))
		case 0xAC:
			code.WriteString(fmt.Sprintf("XRA    H"))
		case 0xAD:
			code.WriteString(fmt.Sprintf("XRA    L"))
		case 0xAE:
			code.WriteString(fmt.Sprintf("XRA    M"))
		case 0xAF:
			code.WriteString(fmt.Sprintf("XRA    A"))
		case 0xB0:
			code.WriteString(fmt.Sprintf("ORA    B"))
		case 0xB1:
			code.WriteString(fmt.Sprintf("ORA    C"))
		case 0xB2:
			code.WriteString(fmt.Sprintf("ORA    D"))
		case 0xB3:
			code.WriteString(fmt.Sprintf("ORA    E"))
		case 0xB4:
			code.WriteString(fmt.Sprintf("ORA    H"))
		case 0xB5:
			code.WriteString(fmt.Sprintf("ORA    L"))
		case 0xB6:
			code.WriteString(fmt.Sprintf("ORA    M"))
		case 0xB7:
			code.WriteString(fmt.Sprintf("ORA    A"))
		case 0xB8:
			code.WriteString(fmt.Sprintf("CMP    B"))
		case 0xB9:
			code.WriteString(fmt.Sprintf("CMP    C"))
		case 0xBA:
			code.WriteString(fmt.Sprintf("CMP    D"))
		case 0xBB:
			code.WriteString(fmt.Sprintf("CMP    E"))
		case 0xBC:
			code.WriteString(fmt.Sprintf("CMP    H"))
		case 0xBD:
			code.WriteString(fmt.Sprintf("CMP    L"))
		case 0xBE:
			code.WriteString(fmt.Sprintf("CMP    M"))
		case 0xBF:
			code.WriteString(fmt.Sprintf("CMP    A"))
		case 0xC0:
			code.WriteString(fmt.Sprintf("RNZ"))
		case 0xC1:
			code.WriteString(fmt.Sprintf("POP    B"))
		case 0xC2:
			code.WriteString(fmt.Sprintf("JNZ    $%02X%02X", mc[i+1], mc[i+2]))
			skip = 2
		case 0xC3:
			code.WriteString(fmt.Sprintf("JMP    $%02X%02X", mc[i+1], mc[i+2]))
			skip = 2
		case 0xC4:
			code.WriteString(fmt.Sprintf("CNZ    $%02X%02X", mc[i+1], mc[i+2]))
			skip = 2
		case 0xC5:
			code.WriteString(fmt.Sprintf("PUSH   B"))
		case 0xC6:
			code.WriteString(fmt.Sprintf("ADI    $%02X", mc[i+1]))
			skip = 1
		case 0xC7:
			code.WriteString(fmt.Sprintf("RST    0"))
		case 0xC8:
			code.WriteString(fmt.Sprintf("RZ"))
		case 0xC9:
			code.WriteString(fmt.Sprintf("RET"))
		case 0xCA:
			code.WriteString(fmt.Sprintf("JZ     $%02X%02X", mc[i+1], mc[i+2]))
			skip = 2
		case 0xCC:
			code.WriteString(fmt.Sprintf("CZ     $%02X%02X", mc[i+1], mc[i+2]))
			skip = 2
		case 0xCD:
			code.WriteString(fmt.Sprintf("CALL   $%02X%02X", mc[i+1], mc[i+2]))
			skip = 2
		case 0xCE:
			code.WriteString(fmt.Sprintf("ACI    $%02X", mc[i+1]))
			skip = 1
		case 0xCF:
			code.WriteString(fmt.Sprintf("RST    1"))
		case 0xD0:
			code.WriteString(fmt.Sprintf("RNC"))
		case 0xD1:
			code.WriteString(fmt.Sprintf("POP    D"))
		case 0xD2:
			code.WriteString(fmt.Sprintf("JNC    $%02X%02X", mc[i+1], mc[i+2]))
			skip = 2
		case 0xD3:
			code.WriteString(fmt.Sprintf("OUT    $%02X", mc[i+1]))
			skip = 1
		case 0xD4:
			code.WriteString(fmt.Sprintf("CNC    $%02X%02X", mc[i+1], mc[i+2]))
			skip = 2
		case 0xD5:
			code.WriteString(fmt.Sprintf("PUSH   D"))
		case 0xD6:
			code.WriteString(fmt.Sprintf("SUI    $%02X", mc[i+1]))
			skip = 1
		case 0xD7:
			code.WriteString(fmt.Sprintf("RST    2"))
		case 0xD8:
			code.WriteString(fmt.Sprintf("RC"))
		case 0xDA:
			code.WriteString(fmt.Sprintf("JC     $%02X%02X", mc[i+1], mc[i+2]))
			skip = 2
		case 0xDB:
			code.WriteString(fmt.Sprintf("IN     $%02X", mc[i+1]))
			skip = 1
		case 0xDC:
			code.WriteString(fmt.Sprintf("CC     $%02X%02X", mc[i+1], mc[i+2]))
			skip = 2
		case 0xDE:
			code.WriteString(fmt.Sprintf("SBI    $%02X", mc[i+1]))
			skip = 1
		case 0xDF:
			code.WriteString(fmt.Sprintf("RST    3"))
		case 0xE0:
			code.WriteString(fmt.Sprintf("RPO"))
		case 0xE1:
			code.WriteString(fmt.Sprintf("POP    H"))
		case 0xE2:
			code.WriteString(fmt.Sprintf("JPO    $%02X%02X", mc[i+1], mc[i+2]))
			skip = 2
		case 0xE3:
			code.WriteString(fmt.Sprintf("XTHL"))
		case 0xE4:
			code.WriteString(fmt.Sprintf("CPO    $%02X%02X", mc[i+1], mc[i+2]))
			skip = 2
		case 0xE5:
			code.WriteString(fmt.Sprintf("PUSH   H"))
		case 0xE6:
			code.WriteString(fmt.Sprintf("ANI    $%02X", mc[i+1]))
			skip = 1
		case 0xE7:
			code.WriteString(fmt.Sprintf("RST    4"))
		case 0xE8:
			code.WriteString(fmt.Sprintf("RPE"))
		case 0xE9:
			code.WriteString(fmt.Sprintf("PCHL"))
		case 0xEA:
			code.WriteString(fmt.Sprintf("JPE    $%02X%02X", mc[i+1], mc[i+2]))
			skip = 2
		case 0xEB:
			code.WriteString(fmt.Sprintf("XCHG"))
		case 0xEC:
			code.WriteString(fmt.Sprintf("CPE    $%02X%02X", mc[i+1], mc[i+2]))
			skip = 2
		case 0xEE:
			code.WriteString(fmt.Sprintf("XRI    $%02X", mc[i+1]))
			skip = 1
		case 0xEF:
			code.WriteString(fmt.Sprintf("RST    5"))
		case 0xF0:
			code.WriteString(fmt.Sprintf("RP"))
		case 0xF1:
			code.WriteString(fmt.Sprintf("POP    PSW"))
		case 0xF2:
			code.WriteString(fmt.Sprintf("JP     $%02X%02X", mc[i+1], mc[i+2]))
			skip = 2
		case 0xF3:
			code.WriteString(fmt.Sprintf("DI"))
		case 0xF4:
			code.WriteString(fmt.Sprintf("CP     $%02X%02X", mc[i+1], mc[i+2]))
			skip = 2
		case 0xF5:
			code.WriteString(fmt.Sprintf("PUSH   PSW"))
		case 0xF6:
			code.WriteString(fmt.Sprintf("ORI    $%02X", mc[i+1]))
			skip = 1
		case 0xF7:
			code.WriteString(fmt.Sprintf("RST    6"))
		case 0xF8:
			code.WriteString(fmt.Sprintf("RM"))
		case 0xF9:
			code.WriteString(fmt.Sprintf("SPHL"))
		case 0xFA:
			code.WriteString(fmt.Sprintf("JM     $%02X%02X", mc[i+1], mc[i+2]))
			skip = 2
		case 0xFB:
			code.WriteString(fmt.Sprintf("EI"))
		case 0xFC:
			code.WriteString(fmt.Sprintf("CM     $%02X%02X", mc[i+1], mc[i+2]))
			skip = 2
		case 0xFE:
			code.WriteString(fmt.Sprintf("CPI    $%02X", mc[i+1]))
			skip = 1
		case 0xFF:
			code.WriteString(fmt.Sprintf("RST    7"))
		}
		code.WriteString("\n")
	}
	return code.String()
}
