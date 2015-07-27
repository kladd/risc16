# RISC-16 (Sort of) - VM in Go

The original plan was to make a RISC-16 VM. But, in the end I've made a hybrid of RISC-16 and MIPS with some syntax borrowed from SPARC. Words are 16 bits wide and there are 8 registers. Register 0 always holds the value 0.

## Usage

Assembly and execution happen in series, and the program reads from stdin.

```bash
cat test.as | ./risc16
```

## Syntax

```
! test.as
!
! Comment character is '!'
!
! registers start with '%', numbered 0-7
! literals, labels, and addresses have no prefix
!
! operation destination, operands...
addi   %1, %0, 2  ! put 2 in register 1
addi   %2, %0, 2  ! put 2 in register 2
add    %1, %2, %3 ! add 2 and 2, put the result in register 1
```

## Instructions implemented so far

* add
* addi
* beq

## License

See [LICENSE](./LICENSE)
