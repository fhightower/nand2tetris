package main

import (
	"bufio"
	"io"
	"log"
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

func determineTypeFromOperation(operation string) string {
	switch operation {

	case "push":
		return C_PUSH
	case "pop":
		return C_POP
	// todo: start here and expand this case
	case "add", "sub":
		return C_ARITHMETIC
		// todo: start here and add more cases
		// C_LABEL      CommandType = "C_LABEL"
		// C_GOTO       CommandType = "C_GOTO"
		// C_IF         CommandType = "C_IF"
		// C_FUNCTION   CommandType = "C_FUNCTION"
		// C_RETURN     CommandType = "C_RETURN"
		// C_CALL       CommandType = "C_CALL"
	}
	log.Fatalf("Unable to determine type for %+v", operation)
	return ""
}

func parseCommand(rawCommand string) (VmCommand, error) {
	vc := VmCommand{}
	vc.Raw = rawCommand

	elements := strings.Split(rawCommand, " ")
	operation := elements[0]
	vc.Type = determineTypeFromOperation(operation)

	// todo: continue updating parsing logic
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

	return cmds, nil
}
