package vm

import (
	"encoding/binary"
	"io"
	"log"

	"github.com/kladd/risc16/spec"
)

var (
	regs [8]int
	prog []uint16
	pc   int
)

func exec(ins uint16) {
	regs[0] = 0

	switch spec.Op(ins) {
	case spec.Add:
		ra := spec.RegA(ins)
		rb := spec.RegB(ins)
		rc := spec.RegC(ins)

		regs[ra] = regs[rb] + regs[rc]
	case spec.Addi:
		ra := spec.RegA(ins)
		rb := spec.RegB(ins)
		simm := spec.Simm(ins)

		regs[ra] = regs[rb] + simm
	case spec.Beq:
		ra := spec.RegA(ins)
		rb := spec.RegB(ins)
		simm := spec.Simm(ins)

		if regs[ra] == regs[rb] {
			pc += simm
		}
	}

	pc++
}

// Exec executes an "executable" file
func Exec(r io.Reader) {
	var p uint16

	err := binary.Read(r, binary.BigEndian, &p)
	for err == nil {
		prog = append(prog, p)
		err = binary.Read(r, binary.BigEndian, &p)
	}

	for {
		exec(prog[pc])

		if pc == 0 || pc >= len(prog) {
			break
		}
	}
}

// Dump dumps VM state to log
func Dump() {
	log.Printf("%5s %d", "PC:", pc)
	for k, v := range regs {
		log.Printf("r[%d]: 0x%04X (%d)", k, v, v)
	}
}
