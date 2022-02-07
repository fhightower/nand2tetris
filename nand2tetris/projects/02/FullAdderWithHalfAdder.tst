// This file is part of www.nand2tetris.org
// and the book "The Elements of Computing Systems"
// by Nisan and Schocken, MIT Press.
// File name: projects/02/FullAdderWithHalfAdder.tst

load FullAdderWithHalfAdder.hdl,
output-file FullAdderWithHalfAdder.out,
compare-to FullAdderWithHalfAdder.cmp,
output-list a%B3.1.3 b%B3.1.3 c%B3.1.3 sum%B3.1.3 carry%B3.1.3;

set a 0,
set b 0,
set c 0,
eval,
output;

set c 1,
eval,
output;

set b 1,
set c 0,
eval,
output;

set c 1,
eval,
output;

set a 1,
set b 0,
set c 0,
eval,
output;

set c 1,
eval,
output;

set b 1,
set c 0,
eval,
output;

set c 1,
eval,
output;
