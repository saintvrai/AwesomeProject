package main

import "fmt"

func main() {
	str := "Это ( \"кавычки\" ) внутри скобки, а это {\"еще кавычки\"} внутри фигурных скобок"
	if GetBalance(str) {
		fmt.Println("Скобки сбалансированы")
	} else {
		fmt.Println("Скобки несбалансированы")
	}
}

// GetBalance Function to determine a balanced string
func GetBalance(input string) bool {
	var stack []rune

	inQuotes := false

	for _, char := range input {
		if char == '"' {
			inQuotes = !inQuotes
		}

		if inQuotes {
			continue
		}

		switch char {
		case '(', '{':
			stack = append(stack, char)
		case ')':
			if len(stack) == 0 || stack[len(stack)-1] != '(' {
				return false
			}
			stack = stack[:len(stack)-1]
		case '}':
			if len(stack) == 0 || stack[len(stack)-1] != '{' {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}

	return len(stack) == 0
}
