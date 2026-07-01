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
	case "or":
		return []string{
			"@SP",
			"AM=M-1",
			"D=M",
			"A=A-1",
			"M=M|D",
		}, labelCounter, nil
	case "not":
		return []string{
			"@SP",
			"A=M-1",
			"M=!M",
		}, labelCounter, nil
	}
	return nil, labelCounter, fmt.Errorf("unsupported arithmetic command: %s", cmd.Arg1)
}

// pushD emits asm that pushes the value in D onto the stack and bumps SP.
var pushD = []string{
	"@SP",
	"A=M", // A = top of stack
	"M=D", // *SP = D
	"@SP",
	"M=M+1", // SP++
}

// pointerSegments maps a VM segment name to its base-pointer symbol.
// These segments are indirect: base held in RAM, address = *base + i.
var pointerSegments = map[string]string{
	"local":    "LCL",
	"argument": "ARG",
	"this":     "THIS",
	"that":     "THAT",
}

func handlePushCommand(cmd VmCommand) ([]string, error) {
	if base, ok := pointerSegments[cmd.Arg1]; ok {
		lines := []string{
			fmt.Sprintf("@%d", cmd.Arg2),
			"D=A", // D = i
			fmt.Sprintf("@%s", base),
			"A=M+D", // A = base + i
			"D=M",   // D = *(base + i)
		}
		return append(lines, pushD...), nil
	}

	switch cmd.Arg1 {
	case "constant":
		lines := []string{
			fmt.Sprintf("@%d", cmd.Arg2),
			"D=A", // D = constant
		}
		return append(lines, pushD...), nil
	case "temp":
		// temp is a fixed 8-slot region at RAM[5..12]; address = 5 + i (direct).
		lines := []string{
			fmt.Sprintf("@%d", 5+cmd.Arg2),
			"D=M", // D = RAM[5 + i]
		}
		return append(lines, pushD...), nil
	case "pointer":
		// pointer 0 -> THIS (RAM[3]), pointer 1 -> THAT (RAM[4]); direct.
		lines := []string{
			fmt.Sprintf("@%d", 3+cmd.Arg2),
			"D=M", // D = THIS/THAT
		}
		return append(lines, pushD...), nil
	}
	return nil, fmt.Errorf("unsupported push segment: %s", cmd.Arg1)
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
		case C_PUSH:
			lines, err := handlePushCommand(cmd)
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
