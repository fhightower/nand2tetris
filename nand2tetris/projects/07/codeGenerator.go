package main

import (
	"fmt"
)

// comparison emits asm for eq/gt/lt. jump is the conditional jump on D
// after computing x-y (e.g. "D;JEQ"). Pushes -1 (true) or 0 (false).
func comparison(jump string, labelCounter int) []string {
	trueLabel := fmt.Sprintf("TRUE_%d", labelCounter)
	endLabel := fmt.Sprintf("END_%d", labelCounter)

	return []string{
		"@SP",
		"AM=M-1",
		"D=M",
		"A=A-1",
		"D=M-D",
		// this ^ is subtraction
		fmt.Sprintf("@%s", trueLabel),
		jump,
		// False path:
		"@SP",
		"A=M-1",
		"M=0",
		fmt.Sprintf("@%s", endLabel),
		"0;JMP",
		// True path:
		fmt.Sprintf("(%s)", trueLabel),
		"@SP",
		"A=M-1",
		"M=-1",
		fmt.Sprintf("(%s)", endLabel),
	}
}

func handleArithmeticCommand(cmd VmCommand, labelCounter int) ([]string, int, error) {
	switch cmd.Arg1 {
	case "add":
		return []string{
			"@SP",
			"AM=M-1", // SP--, A = address of top value (y)
			"D=M",    // D = y
			"A=A-1",  // A = address of next value (x)
			"M=M+D",  // x = x + y
		}, labelCounter, nil
	case "sub":
		return []string{
			"@SP",
			"AM=M-1",
			"D=M",
			"A=A-1",
			"M=M-D",
		}, labelCounter, nil
	case "neg":
		return []string{
			"@SP",
			"A=M-1",
			"M=-M",
		}, labelCounter, nil
	case "eq":
		return comparison("D;JEQ", labelCounter), labelCounter + 1, nil
	case "gt":
		return comparison("D;JGT", labelCounter), labelCounter + 1, nil
	case "lt":
		return comparison("D;JLT", labelCounter), labelCounter + 1, nil
	case "and":
		return []string{
			"@SP",
			"AM=M-1",
			"D=M",
			"A=A-1",
			"M=M&D",
		}, labelCounter, nil
		// todo: start here and build this case statement out
	}
	return nil, labelCounter, fmt.Errorf("unsupported arithmetic command: %s", cmd.Arg1)
}

func GenerateAssembly(cmds []VmCommand) ([]string, error) {
	var assembly []string
	labelCounter := 0

	// todo: start here and write this function
	for _, cmd := range cmds {
		switch cmd.Type {
		case C_ARITHMETIC:
			lines, newLabelCounter, err := handleArithmeticCommand(cmd, labelCounter)
			labelCounter = newLabelCounter
			if err != nil {
				return assembly, err
			}
			assembly = append(assembly, lines...)
		default:
			return assembly, fmt.Errorf("unable to convert command %+v: %s", cmd.Type, cmd.Raw)
		}
	}

	return assembly, nil
}
