package parser

import (
	"bufio"
	"io"
	"log"
	"strconv"
	"strings"
)

type AsmCommand struct {
	Raw string

	IsACommand bool
	IsCCommand bool
	IsLCommand bool

	LSymbol string
	ASymbol int
	Dest    string
	Comp    string
	Jump    string
}

func handleACommand(line string, ac AsmCommand) AsmCommand {
	ac.IsACommand = true
	symbol, err := strconv.Atoi(line[1:])
	if err != nil {
		log.Fatal("invalid A-command value %q: %w", symbol, err)
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

func Parse(r io.Reader) ([]AsmCommand, error) {
	var cmds []AsmCommand
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
	return cmds, nil
}
