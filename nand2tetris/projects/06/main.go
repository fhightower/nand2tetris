package main

import (
	"log"
	"os"

	"projects06/assembler/parser"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("usage: assembler <input.asm>")
	}

	f, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	prog, err := parser.Parse(f)
	if err != nil {
		log.Fatal(err)
	}

	binaryStrings, err := parser.ConvertAsmToBinary(prog)
	if err != nil {
		log.Fatal(err)
	}

	_ = prog
	// TODO: pass prog to assembler/codegen stage.
}
