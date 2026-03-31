package parser

import (
	"bufio"
	"errors"
	"io"
	"log"
	"strconv"
	"strings"
)

const startOfFreeMemory = 16
const screenMemoryLoc = 16384

var symbolTable = map[string]int{
	"SP":     0,
	"LCL":    1,
	"ARG":    2,
	"THIS":   3,
	"THAT":   4,
	"R0":     0,
	"R1":     1,
	"R2":     2,
	"R3":     3,
	"R4":     4,
	"R5":     5,
	"R6":     6,
	"R7":     7,
	"R8":     8,
	"R9":     9,
	"R10":    10,
	"R11":    11,
	"R12":    12,
	"R13":    13,
	"R14":    14,
	"R15":    15,
	"SCREEN": screenMemoryLoc,
	"KBD":    24576,
}

type AsmCommand struct {
	Raw string

	IsACommand bool
	IsCCommand bool
	IsLCommand bool

	LSymbol string
	ASymbol int

	Dest string
	Comp string
	Jump string
}

func handleACommand(line string, ac AsmCommand) AsmCommand {
	ac.IsACommand = true
	symbol, err := strconv.Atoi(line[1:])
	if err != nil {
		memoryLoc, exists := GetSymbolMemoryLoc(line[1:])
		if !exists {
			// log.Fatalf("Invalid A-command value %q", symbol)
			return ac
		}
		symbol = memoryLoc
	}
	ac.ASymbol = symbol

	return ac
}

func handleLCommand(line string, ac AsmCommand) AsmCommand {
	ac.IsLCommand = true
	ac.LSymbol = line[1 : len(line)-1]
	return ac
}

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

func handleCCommand(line string, ac AsmCommand) AsmCommand {
	ac.IsCCommand = true
	ac.Dest = findDest(line)
	ac.Jump = findJump(line)
	ac.Comp = findComp(line, ac.Dest, ac.Jump)
	return ac
}

func parseLine(line string) (AsmCommand, error) {
	ac := AsmCommand{}
	ac.Raw = line

	if strings.HasPrefix(line, "@") {
		ac = handleACommand(line, ac)
	} else if strings.HasPrefix(line, "(") && strings.HasSuffix(line, ")") {
		ac = handleLCommand(line, ac)
	} else {
		ac = handleCCommand(line, ac)
	}

	return ac, nil
}

func GetSymbolMemoryLoc(symbol string) (int, bool) {
	value, exists := symbolTable[symbol]
	return value, exists
}

func findNextAvailableMemLocation() (int, error) {
	used := make(map[int]bool)

	for _, value := range symbolTable {
		used[value] = true
	}

	for i := startOfFreeMemory; i < screenMemoryLoc; i++ {
		if !used[i] {
			return i, nil
		}
	}
	return 0, errors.New("no available memory")
}

func isASCIILetter(c byte) bool {
	return (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z')
}

func Parse(r io.Reader) ([]AsmCommand, error) {
	var cmds []AsmCommand
	scanner := bufio.NewScanner(r)
	// This may not be the best/most efficient way to do this, but it works for now
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
			log.Fatalf("Unable to parse before pre-processing: %q", text)
		}
		aCommandHasSymbol := cmd.IsACommand && isASCIILetter(cmd.Raw[1])
		if cmd.IsLCommand || aCommandHasSymbol {
			relevantSymbol := cmd.LSymbol
			if cmd.IsACommand {
				relevantSymbol = cmd.Raw[1:]
			}
			_, exists := GetSymbolMemoryLoc(relevantSymbol)
			if !exists {
				nextLoc, err := findNextAvailableMemLocation()
				if err != nil {
					panic("no available memory location")
				}
				symbolTable[relevantSymbol] = nextLoc
			}
		}
		cmd, err = parseLine(text)
		if err != nil {
			log.Fatalf("Unable to parse after pre-processing: %q", text)
		}
		cmds = append(cmds, cmd)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return cmds, nil
}
