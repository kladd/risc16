package main

import (
	"bytes"
	"os"

	"github.com/kladd/risc16/as"
	"github.com/kladd/risc16/vm"
)

func main() {
	buf := new(bytes.Buffer)

	as.AssembleFile(os.Stdin, buf)
	vm.Exec(buf)

	vm.Dump()
}
