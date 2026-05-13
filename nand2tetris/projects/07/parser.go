package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

type CommandType string

const (
	C_ARITHMETIC CommandType = "C_ARITHMETIC"
	C_PUSH       CommandType = "C_PUSH"
	C_POP        CommandType = "C_POP"
	C_LABEL      CommandType = "C_LABEL"
	C_GOTO       CommandType = "C_GOTO"
	C_IF         CommandType = "C_IF"
	C_FUNCTION   CommandType = "C_FUNCTION"
	C_RETURN     CommandType = "C_RETURN"
	C_CALL       CommandType = "C_CALL"
)

type VmCommand struct {
	Raw string

	Type CommandType
	// This is optional for some command types (e.g. C_RETURN)
	Arg1 string
	// Also optional - only used for C_PUSH, C_POP, C_FUNCTION and C_CALL; zero otherwise.
	Arg2 int
}

func parseCommand(rawCommand string) (VmCommand, error) {
	vc := VmCommand{}

	// todo: start here: update parsing logic
	return vc, nil
}

func ParseFile(r io.ReadSeeker) ([]VmCommand, error) {
	var cmds []VmCommand
	scanner := bufio.NewScanner(r)

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		// Skip comments and empty lines
		if strings.HasPrefix(line, "//") || line == "" {
			continue
		}

		// Handle in-line comments
		parts := strings.SplitN(line, "//", 2)
		rawCommand := strings.TrimSpace(parts[0])

		vmCommand, err := parseCommand(rawCommand)

		if err != nil {
			return cmds, nil
		}
		cmds = append(cmds, vmCommand)

	}

	fmt.Println(len(cmds))
	return cmds, nil
}
