// This file is part of www.nand2tetris.org
// and the book "The Elements of Computing Systems"
// by Nisan and Schocken, MIT Press.
// File name: projects/04/Fill.asm

// Runs an infinite loop that listens to the keyboard input.
// When a key is pressed (any key), the program blackens the screen,
// i.e. writes "black" in every pixel;
// the screen should remain fully black as long as the key is pressed. 
// When no key is pressed, the program clears the screen, i.e. writes
// "white" in every pixel;
// the screen should remain fully clear as long as no key is pressed.

(LISTENER)
    @KBD
    D=M

    // fill screen
    @FILL_SCREEN
    D;JGT
    // otherwise, uncolor the screen
    @UNFILL_SCREEN
    0;JMP
(END)

(FILL_SCREEN)
    // put the value we want to write to the screen in R0
    @pixel
    M=1

    @COLOR_SCREEN
    0;JMP
(END)

(UNFILL_SCREEN)
    // put the value we want to write to the screen in R0
    @pixel
    M=0

    @COLOR_SCREEN
    0;JMP
(END)

(COLOR_SCREEN)
    @16384
    D=A
    @loc
    M=D

    (FILL_PIXEL)
        // D = pixel
        @pixel
        D=M

        @loc
        A=M
        M=D

        @loc
        M=M+1
        D=M

        @16414
        D=D-A

        // keep drawing if i - 255 is < zero
        @FILL_PIXEL
        D;JLT
    (END)

    @LISTENER
    0;JMP
(END)
