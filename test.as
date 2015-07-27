! adds two numbers together
addi	%2, %0, 2
addi	%3, %0, 2
beq		%3, %2, 1
addi	%4, %0, 1 ! should be skipped
addi	%5, %0, 1
