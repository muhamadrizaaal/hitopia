package main

import (
	"fmt"
)

// fungsi untuk menghitung bobot karakter
func charWeight(c rune) int {
	return int(c - 'a' + 1)
}

// fungsi untuk menghitung bobot substring dalam string
func calculateWeights(s string) map[int]bool {
	weights := make(map[int]bool)
	n := len(s)

	for i := 0; i < n; i++ {
		sum := 0
		for j := i; j < n && s[i] == s[j]; j++ {
			sum += charWeight(rune(s[i]))
			weights[sum] = true
		}
	}

	return weights
}

// fungsi untuk memproses queries berdasarkan bobot substring
func weightedStrings(s string, queries []int) []string {
	weights := calculateWeights(s)
	results := make([]string, len(queries))

	for i, query := range queries {
		if weights[query] {
			results[i] = "Yes"
		} else {
			results[i] = "No"
		}
	}

	return results
}

// fungsi utama untuk menguji hasil
func main() {
	s := "abbcccd"
	queries := []int{1, 3, 9, 8}
	result := weightedStrings(s, queries)
	fmt.Println(result) 
}
