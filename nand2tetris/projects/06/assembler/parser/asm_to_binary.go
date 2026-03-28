package parser

import (
	"errors"
	"fmt"
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

func convertComp(command AsmCommand) string {
	switch command.Comp {
	// a=0
	case "0":
		return "101010"
	case "1":
		return "111111"
	case "-1":
		return "111010"
	case "D":
		return "001100"
	case "A":
		return "110000"
	case "!D":
		return "001101"
	case "!A":
		return "110001"
	case "-D":
		return "001111"
	case "-A":
		return "110011"
	case "D+1":
		return "011111"
	case "A+1":
		return "110111"
	case "D-1":
		return "001110"
	case "A-1":
		return "110010"
	case "D+A":
		return "000010"
	case "D-A":
		return "010011"
	case "A-D":
		return "000111"
	case "D&A":
		return "000000"
	case "D|A":
		return "010101"
	// a=1
	case "M":
		return "110000"
	case "!M":
		return "110001"
	case "-M":
		return "110011"
	case "M+1":
		return "110111"
	case "M-1":
		return "110010"
	case "D+M":
		return "000010"
	case "D-M":
		return "010011"
	case "M-D":
		return "000111"
	case "D&M":
		return "000000"
	case "D|M":
		return "010101"
	}
	return ""
}

func convertDest(command AsmCommand) string {
	switch command.Dest {
	case "M":
		return "001"
	case "D":
		return "010"
	case "MD":
		return "011"
	case "A":
		return "100"
	case "AM":
		return "101"
	case "AD":
		return "110"
	case "AMD":
		return "111"
	}
	// Unknown dest mnemonic; default to no dest.
	return "000"
}

func convertJump(command AsmCommand) string {
	switch command.Jump {
	case "JGT":
		return "001"
	case "JEQ":
		return "010"
	case "JGE":
		return "011"
	case "JLT":
		return "100"
	case "JNE":
		return "101"
	case "JLE":
		return "110"
	case "JMP":
		return "111"
	}
	return "000"
}

func convertCCommandToBinary(command AsmCommand) string {
	return fmt.Sprintf("111%s%s%s", convertComp(command), convertDest(command), convertJump(command))
}

func convertACommandToBinary(command AsmCommand) string {
	return fmt.Sprintf("0%015b", command.ASymbol)
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

func convertLCommandToBinary(command AsmCommand) string {
	memoryLoc, exists := GetSymbolMemoryLoc(command.LSymbol)
	if !exists {
		nextLoc, err := findNextAvailableMemLocation()
		if err != nil {
			panic("no available memory location")
		}
		symbolTable[command.LSymbol] = nextLoc
		memoryLoc = nextLoc
	}

	// todo: start here...
	return fmt.Sprintf("0%015b", memoryLoc)
}

func convertCommandToBinary(command AsmCommand) (string, error) {
	if command.IsCCommand {
		return convertCCommandToBinary(command), nil
	} else if command.IsACommand {
		return convertACommandToBinary(command), nil
	} else if command.IsLCommand {
		return convertLCommandToBinary(command), nil
	} else {
		return "", fmt.Errorf("Command of unknown type: %v", command)
	}
}

func ConvertAsmToBinary(asmCommands []AsmCommand) ([]string, error) {
	var binaryCommands []string
	for _, cmd := range asmCommands {
		newBinaryCommand, err := convertCommandToBinary(cmd)
		if err != nil {
			return nil, err
		}
		fmt.Printf("%+v\n", cmd)
		fmt.Printf("%+v\n", newBinaryCommand)
		binaryCommands = append(binaryCommands, newBinaryCommand)
	}
	return binaryCommands, nil
}
