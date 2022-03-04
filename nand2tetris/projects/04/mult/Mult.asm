// This file is part of www.nand2tetris.org
// and the book "The Elements of Computing Systems"
// by Nisan and Schocken, MIT Press.
// File name: projects/04/Mult.asm

// Multiplies R0 and R1 and stores the result in R2.
// (R0, R1, R2 refer to RAM[0], RAM[1], and RAM[2], respectively.)
//
// This program only needs to handle arguments that satisfy
// R0 >= 0, R1 >= 0, and R0*R1 < 32768.

// In this code:
// `i` is the counter

    // i = 0
    @i
    M=0

    // R2 = 0
    @2
    M=0

    // check if R0 is zero (and end if so)
    @0
    D=M
    @END
    D;JEQ

    // check if R1 is zero (and end if so)
    @1
    D=M
    @END
    D;JEQ
(LOOP)
    // i++
    @i
    M=M+1

    // D = R0
    @0
    D=M

    // R2 += D
    @2
    M=M+D

    // D = R1
    @1
    D=M

    // D -= i
    @i
    D=D-M

    // end if i == R1
    @END
    D;JEQ
    // otherwise, continue looping
    @LOOP
    0;JMP
(END)
    // infinite loop to signal end of program
    @END
    0;JMP
