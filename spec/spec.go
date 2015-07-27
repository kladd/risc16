// Package spec contains constants and utility functions that depend on the spec
// of the VM's instruction set, architecture, etc.
package spec

const (
	// Word defines word size in bits
	Word = 16

	// NRegs defines the number of registers
	NRegs = 8

	// s = shift, m = mask
	sop   = 13
	sra   = 10
	srb   = 7
	ssimm = 9
	mr    = 0x0007
	msimm = 0x007F

	Add  = 0
	Addi = 1
	Nand = 2
	Lui  = 3
	Sw   = 4
	Lw   = 5
	Beq  = 6
	Jalr = 7
	Ext  = 7
)

// Op decodes opcode from an instruction
func Op(ins uint16) int {
	return int(ins >> sop)
}

// RegA decodes regA from an instruction
func RegA(ins uint16) int {
	return int(ins >> sra & mr)
}

// RegB decodes regB from an instruction
func RegB(ins uint16) int {
	return int(ins >> srb & mr)
}

// RegC decodes regC from an instruction
func RegC(ins uint16) int {
	return int(ins & mr)
}

// Simm decodes signed imm from an instruction
func Simm(ins uint16) int {
	return int(int16(ins<<ssimm) >> ssimm)
}

// EncodeOp returns an opcode in instruction format
func EncodeOp(op uint16) uint16 {
	return op << sop
}

// EncodeRRR returns an RRR encoded instruction
func EncodeRRR(regA uint16, regB uint16, regC uint16) uint16 {
	return regA<<sra | regB<<srb | regC
}

// EncodeRRI returns an RRI encoded instruction
func EncodeRRI(regA uint16, regB uint16, simm int) uint16 {
	return regA<<sra | regB<<srb | (uint16(simm) & msimm)
}
