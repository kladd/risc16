add     %0, %0, %0  ! dummy instruction (for now)
addi    %1, %0, 5   ! put 5 in reg 1
jalr    %2, %1      ! jump to instruction 4

addi    %7, %0, 1   ! [skipped] put 1 in reg 7
jalr    %2, %0      ! [skipped] jump to instruction 0, halting the program

addi    %6, %0, 1   ! put 1 in reg 6
beq     %0, %0, 1   ! skip next instruction if 0 == 0
jalr    %2, 0       ! halts the program if beq doesn't work

beq     %0, %1, 1   ! skip next instruction if 0 == 5
addi    %5, %0, 4   ! put 5 in reg 5
jalr    %2, %5      ! jump to the jump that exits the program

! this program should result in 1 being in reg 6, and not in reg 7
