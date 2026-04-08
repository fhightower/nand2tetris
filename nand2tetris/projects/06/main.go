package main

import (
	"log"
	"os"
	"strings"

	"projects06/assembler/parser"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("usage: assembler <input.asm>")
	}

	fileName := os.Args[1]
	f, err := os.Open(fileName)
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

	hackFileName := strings.Replace(fileName, ".asm", ".hack", 1)
	outputFile, err := os.Create(hackFileName)
	if err != nil {
		log.Fatal(err)
	}
	defer outputFile.Close()

	outputFile.WriteString(strings.Join(binaryStrings, "\n"))
}
