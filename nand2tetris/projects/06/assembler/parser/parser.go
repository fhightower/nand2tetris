package parser

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

type BinaryCommand struct {
	Raw string

	IsACommand bool
	IsCCommand bool
	IsLCommand bool

	Symbol string
	Dest   string
	Comp   string
	Jump   string
}

func handleACommand(line string, bc BinaryCommand) BinaryCommand {
	bc.IsACommand = true
	bc.Symbol = line[1:]
	return bc
}

func handleLCommand(line string, bc BinaryCommand) BinaryCommand {
	bc.IsLCommand = true
	bc.Symbol = line[1 : len(line)-1]
	return bc
}

// func findDest(line string) string {
// 	i := strings.Index(line, "=")
// 	if i == -1 {
// 		return "000"
// 	}
// 	switch strings.TrimSpace(line[:i]) {
// 	case "M":
// 		return "001"
// 	case "D":
// 		return "010"
// 	case "MD":
// 		return "011"
// 	case "A":
// 		return "100"
// 	case "AM":
// 		return "101"
// 	case "AD":
// 		return "110"
// 	case "AMD":
// 		return "111"
// 	}
// 	// Unknown dest mnemonic; default to no dest.
// 	return "000"
// }

func findDest(line string) string {
	i := strings.Index(line, "=")
	if i == -1 {
		return ""
	} else {
		return strings.TrimSpace(line[:i])
	}
}

func findComp(line string, dest string, jump string) string {
	comp := strings.TrimPrefix(line, dest+"=")
	comp = strings.TrimSuffix(comp, ";"+jump)
	return strings.TrimSpace(comp)
}

func findJump(line string) string {
	i := strings.Index(line, ";")
	if i == -1 {
		return ""
	} else {
		return strings.TrimSpace(line[i+1:])
	}
}

func handleCCommand(line string, bc BinaryCommand) BinaryCommand {
	bc.IsCCommand = true
	bc.Dest = findDest(line)
	bc.Jump = findJump(line)
	bc.Comp = findComp(line, bc.Dest, bc.Jump)
	return bc
}

func parseLine(line string) (BinaryCommand, error) {
	bc := BinaryCommand{}
	bc.Raw = line

	if strings.HasPrefix(line, "@") {
		bc = handleACommand(line, bc)
	} else if strings.HasPrefix(line, "(") && strings.HasSuffix(line, ")") {
		bc = handleLCommand(line, bc)
	} else {
		bc = handleCCommand(line, bc)
	}

	return bc, nil
}

func Parse(r io.Reader) ([]BinaryCommand, error) {
	var cmds []BinaryCommand
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		text := scanner.Text()

		// Skip comments and empty lines
		if strings.HasPrefix(text, "//") || text == "" {
			continue
		}

		// Handle in-line comments
		parts := strings.SplitN(text, "//", 2)
		text = strings.TrimSpace(parts[0])

		cmd, err := parseLine(text)
		if err != nil {
			return nil, err
		}
		cmds = append(cmds, cmd)
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	for i, cmd := range cmds {
		fmt.Printf("%d: %+v\n", i, cmd)
	}
	return cmds, nil
}
