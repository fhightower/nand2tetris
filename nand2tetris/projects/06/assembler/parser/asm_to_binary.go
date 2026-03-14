package parser

import (
	"fmt"
)

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

func convertCommandToBinary(command AsmCommand) (string, error) {
	pass
}

func ConvertAsmToBinary(asmCommands []AsmCommand) ([]string, error) {
	var binaryCommands []string
	for i, cmd := range asmCommands {
		fmt.Printf("%d: %+v\n", i, cmd)
		// todo: start here...
	}
	return cmds, nil
}
