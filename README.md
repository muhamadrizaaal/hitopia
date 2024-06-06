# Balanced Brackets Checker

## Deskripsi
Program ini memeriksa apakah tanda kurung dalam sebuah string seimbang (balanced). Tanda kurung yang diperbolehkan adalah `(`, `)`, `{`, `}`, `[`, dan `]`.

## Aturan
1. Tanda kurung harus memiliki pasangan yang sesuai dan urutannya benar.
2. Tanda kurung bisa dipisahkan dengan atau tanpa whitespace.

## Kompleksitas
- **Waktu (Time Complexity):** O(n)
  - Program melakukan satu pass melalui string input dan setiap operasi pada stack (push dan pop) adalah O(1).
- **Ruang (Space Complexity):** O(n)
  - Dalam kasus terburuk, semua tanda kurung buka akan disimpan dalam stack.

## Cara Menggunakan
1. Import package `fmt`.
2. Definisikan fungsi `isBalanced(s string) string`.
3. Panggil fungsi ini dengan string input untuk memeriksa apakah seimbang.

## Contoh
```go
package main

import (
	"fmt"
)

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
