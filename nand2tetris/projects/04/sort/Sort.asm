// Given 3, unique numbers (between [0-9] in decimal), sort them

// We assume that the input numbers are given in R10-R12
// We use R0-9 as the output

// define the starting point
@10
D=A
@start
M=D 
@current
M=D

// define the ending point
@12
D=A
@end
M=D 

(LOOP)
        // check to see if we are done...
        @current
        D=M
        @end
        D=M-D
        @END
        D;JLT
        
	// b/c we are not done, update the register for the current value to 1
        @current
	// this took me a while to figure out :)
        A=M
	A=M
        M=1
        
	// increment the current index
        @current
        M=M+1

	// go around again...
	@LOOP
	0;JMP
(END)

// infinite loop to signal end of program
@END        
0;JMP

