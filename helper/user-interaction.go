package helper

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

func AskToContinue(scanner *bufio.Scanner) bool {
	fmt.Print("\nDo you want to continue browsing? (y/n): ")
	scanner.Scan()
	input := strings.TrimSpace(scanner.Text())
	return strings.ToLower(input) == "y" || strings.ToLower(input) == "yes"
}

func IsExitCommand(input string) bool {
	exitCommands := []string{"n", "no", "exit", "quit", "q"}
	lower := strings.ToLower(strings.TrimSpace(input))
	for _, cmd := range exitCommands {
		if lower == cmd {
			return true
		}
	}
	return false
}

func ReadInt(scanner *bufio.Scanner) (int, error) {
	scanner.Scan()
	input := strings.TrimSpace(scanner.Text())
	if IsExitCommand(input) {
		return 0, nil
	}
	return strconv.Atoi(input)
}
