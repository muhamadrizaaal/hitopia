package main

import (
	"fmt"
)

// fungsi untuk mengecek apakah bracket seimbang
func isBalanced(s string) string {
	stack := []rune{}
	bracketMap := map[rune]rune{
		')': '(',
		'}': '{',
		']': '[',
	}

	for _, char := range s {
		switch char {
		case '(', '{', '[':
			stack = append(stack, char)
		case ')', '}', ']':
			if len(stack) == 0 || stack[len(stack)-1] != bracketMap[char] {
				return "NO"
			}
			stack = stack[:len(stack)-1]
		}
	}

	if len(stack) == 0 {
		return "YES"
	} else {
		return "NO"
	}
}

// fungsi utama untuk menguji hasil
func main() {
	testCases := []string{
		"{ [ ( ) ] }",
		"{ [ ( ] ) }",
		"{ ( ( [ ] ) [ ] ) [ ] }",
	}

	for _, testCase := range testCases {
		fmt.Printf("Input: %s\nOutput: %s\n\n", testCase, isBalanced(testCase))
	}
}
