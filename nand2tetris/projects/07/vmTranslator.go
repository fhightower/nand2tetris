package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

// todo(later): handle directory input
func writeAssembly(fileName string, assembly []string) error {
	assemblyFileName := strings.Replace(fileName, ".vm", ".asm", 1)
	outputFile, err := os.Create(assemblyFileName)
	if err != nil {
		log.Fatal(err)
	}
	defer outputFile.Close()

	outputFile.WriteString(strings.Join(assembly, "\n"))
	return nil
}

func main() {
	if len(os.Args) < 2 {
		log.Fatal("usage: VMtranslator <source>")
	}

	// todo(later): handle directory
	fileName := os.Args[1]
	f, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	prog, err := ParseFile(f)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(prog)

	assembly, err := GenerateAssembly(prog)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(assembly)

	err = writeAssembly(fileName, assembly)
	if err != nil {
		log.Fatal(err)
	}
}
