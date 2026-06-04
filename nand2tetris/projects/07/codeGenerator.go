package main

import (
	"fmt"
)

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
		trueLabel := fmt.Sprintf("TRUE_%d", labelCounter)
		endLabel := fmt.Sprintf("END_%d", labelCounter)

		result := []string{
			"@SP",
			"AM=M-1",
			"D=M",
			"A=A-1",
			"D=M-D",
			// this ^ is subtraction
			fmt.Sprintf("@%s", trueLabel),
			"D;JEQ",
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
		return result, labelCounter + 1, nil
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
