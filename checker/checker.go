package checker

import (
	"fmt"
	"strings"
)

// isValidInput validates the input and returns a bool
func IsValidInput(args []string) bool {
	valid := true

	if len(args) < 2 || len(args) > 3 {
		fmt.Println("ERROR: wrong number of arguments\nUsage: go run . [COLOR] [STRING] [SUBSTRING]\n\nEX: go run . --color=<color> something <substring>")
		return false
	}

	if len(args[1]) == 0 {
		valid = false
	}

	if len(args) <= 3 {
		if !strings.Contains(args[0], "--color=") {
			fmt.Println("ERROR:invalid color flag")
			valid = false
		} else if !AllowedColors(strings.TrimPrefix(args[0], "--color=")) {
			fmt.Println("ERROR: only colors red, green, yellow, blue, purple, cyan, white, grey are allowed")
			valid = false
		}

		if !AllowedChars(args[1]) {
			fmt.Println("ERROR: only characters between 32 and 126 are allowed")
			valid = false
		}

		if len(args) == 3 {
			if !isAlphabet(args[2]) {
				fmt.Println("ERROR: invalid substring")
				valid = false
			} else {
				if !IsNumeralOrColon(args[2]) {
					if !strings.Contains(args[1], args[2]) {
						fmt.Println("ERROR: substring invalid")
						valid = false
					}
				}
			}
		}

	}

	if !valid {
		fmt.Println("Usage: go run . [COLOR] [STRING] [SUBSTRING]\n\nEX: go run . --color=<color> something <substring>")
		return false
	}
	return valid
}

// AllEmpty checks if all elements in a slice are empty and returns a bool
func AllEmpty(inputSlice []string) bool {
	for i := 0; i < len(inputSlice); i++ {
		if inputSlice[i] != "" {
			return false
		}
	}
	return true
}

// AllowedChars checks that each letter in a string contains only character 32 to 126 and returns a bool
func AllowedChars(s string) bool {
	for i := 0; i < len(s); i++ {
		if s[i] < 32 || s[i] > 126 {
			return false
		}
	}
	return true
}

// AllowedColors checks that only 8 specific colors (red, green, yellow, blue, purple, cyan, white, and grey) are allowed
func AllowedColors(s string) bool {
	s = strings.ToLower(s)
	if s == "red" || s == "green" || s == "yellow" || s == "blue" || s == "purple" || s == "cyan" || s == "white" || s == "grey" {
		return true
	}
	return false
}

// isAlphabet checks if s is an alphabet
func isAlphabet(s string) bool {
	for i := 0; i < len(s); i++ {
		if s[i] <= 'a' && s[i] >= 'z' || s[i] <= 'A' && s[i] >= 'Z' {
			return false
		}
	}
	return true
}

// Contains checks if s contains substr in the same order
func Contains(s string, substr string) bool {
	if len(substr) == 1 {
		for i := 0; i < len(s); i++ {
			if s[i] == substr[0] {
				return true
			}
		}
		return false
	}
	for _, c := range substr {
		i := strings.IndexRune(s, c)
		if i == -1 {
			return false
		}
		s = s[i+1:]
	}
	return true
}

// IsNumeral checks if s is a number
func IsNumeral(s string) bool {
	for _, c := range s {
		if c < '0' || c > '9' {
			return false
		}
	}
	return true
}

// IsNumeralOrColon checks if s is a number or colon
func IsNumeralOrColon(s string) bool {
	for _, c := range s {
		if c < '0' || c > '9' && c != ':' {
			return false
		}
	}
	return true
}
