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

// the max screen size is: 16384 + (32 * 256)
// where 16384 is the starting address of the screen,
// 32 is the number of words per row,
// and 256 is the number of rows (so 32 * 256 finds the total number of words on the screen)
@24576
D=A
@SCREEN_MAX
M=D

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
    @pixel
    M=-1

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
    @SCREEN
    D=A
    // define "loc" which will keep track of the current word on the screen
    @loc
    M=D

    (FILL_PIXEL)
        @pixel
        D=M

        @loc
        A=M
        M=D

        @loc
        M=M+1
        D=M

        @SCREEN_MAX
        D=D-M

        // keep drawing if i - 255 is < zero
        @FILL_PIXEL
        D;JLT
    (END)

    @LISTENER
    0;JMP
(END)
