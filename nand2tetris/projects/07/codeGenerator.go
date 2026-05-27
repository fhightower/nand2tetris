package main

import (
	"fmt"
)

func handleArithmeticCommand(cmd VmCommand) ([]string, error) {
	switch cmd.Arg1 {
	case "add":
		return []string{
			"@SP",
			"AM=M-1", // SP--, A = address of top value (y)
			"D=M",    // D = y
			"A=A-1",  // A = address of next value (x)
			"M=D+M",  // x = y + x
		}, nil
		// todo: start here and build this case statement out
	}
	return nil, fmt.Errorf("unsupported arithmetic command: %s", cmd.Arg1)
}

func GenerateAssembly(cmds []VmCommand) ([]string, error) {
	var assembly []string

	// todo: start here and write this function
	for _, cmd := range cmds {
		switch cmd.Type {
		case C_ARITHMETIC:
			lines, err := handleArithmeticCommand(cmd)
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
