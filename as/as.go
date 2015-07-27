package as

import (
	"bufio"
	"encoding/binary"
	"errors"
	"io"
	"strconv"
	"strings"
	"unicode"

	"github.com/kladd/risc16/spec"
)

func parseReg(arg string) (uint16, error) {
	if arg[0] != '%' {
		return 0, errors.New("Invalid register format.")
	}

	n, err := strconv.Atoi(arg[1:])
	if n < 0 || n > spec.NRegs {
		err = errors.New("Register out of range.")
	}

	return uint16(n), err
}

func parseRRR(args []string) (uint16, error) {
	regA, err := parseReg(args[1])
	regB, err := parseReg(args[2])
	regC, err := parseReg(args[3])

	return spec.EncodeRRR(regA, regB, regC), err
}

func parseRRI(args []string) (uint16, error) {
	regA, err := parseReg(args[1])
	regB, err := parseReg(args[2])
	imm, err := strconv.Atoi(args[3])

	return spec.EncodeRRI(regA, regB, imm), err
}

func parse(args []string) (uint16, error) {
	var out uint16

	switch args[0] {
	case "add":
		out, err := parseRRR(args)
		return spec.EncodeOp(spec.Add) | out, err
	case "addi":
		out, err := parseRRI(args)
		return spec.EncodeOp(spec.Addi) | out, err
	case "beq":
		out, err := parseRRI(args)
		return spec.EncodeOp(spec.Beq) | out, err
	case "jalr":
		out, err := parseRRI(append(args, "0"))
		return spec.EncodeOp(spec.Jalr) | out, err
	}

	return out, errors.New("Invalid opcode")
}

// AssembleFile converts assembly to machine code
func AssembleFile(r io.Reader, w io.Writer) {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		ins := strings.Split(strings.TrimSpace(scanner.Text()), "!")[0]

		if ins != "" {
			var enc uint16

			enc, _ = parse(strings.FieldsFunc(ins, func(c rune) bool {
				return c == ',' || unicode.IsSpace(c)
			}))

			binary.Write(w, binary.BigEndian, enc)
		}
	}
}
