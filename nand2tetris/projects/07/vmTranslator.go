package main

import (
	"fmt"
	"log"
	"os"
)

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

	// todo: generate output file
}
