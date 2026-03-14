package parser

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

type AsmCommand struct {
	Raw string

	IsACommand bool
	IsCCommand bool
	IsLCommand bool

	Symbol string
	Dest   string
	Comp   string
	Jump   string
}

func handleACommand(line string, ac AsmCommand) AsmCommand {
	ac.IsACommand = true
	ac.Symbol = line[1:]
	return ac
}

func handleLCommand(line string, ac AsmCommand) AsmCommand {
	ac.IsLCommand = true
	ac.Symbol = line[1 : len(line)-1]
	return ac
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
	for i, cmd := range cmds {
		fmt.Printf("%d: %+v\n", i, cmd)
	}
	return cmds, nil
}
