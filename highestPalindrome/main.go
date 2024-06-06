package main

import (
	"fmt"
)

// fungsi untuk memeriksa apakah string adalah palindrome
func isPalindrome(s string) bool {
	n := len(s)
	for i := 0; i < n/2; i++ {
		if s[i] != s[n-1-i] {
			return false
		}
	}
	return true
}

// fungsi rekursif untuk membuat string menjadi palindrome
func makePalindrome(s string, k int, l int, r int) (string, int) {
	if l >= r {
		return s, k
	}

	if s[l] == s[r] {
		return makePalindrome(s, k, l+1, r-1)
	}

	if k == 0 {
		return s, k
	}

	if s[l] > s[r] {
		s = s[:r] + string(s[l]) + s[r+1:]
	} else {
		s = s[:l] + string(s[r]) + s[l+1:]
	}

	return makePalindrome(s, k-1, l+1, r-1)
}

// fungsi rekursif untuk meningkatkan nilai string menjadi palindrome tertinggi
func maximizePalindrome(s string, k int, l int, r int) string {
	if l > r {
		return s
	}

	if s[l] == s[r] {
		s = maximizePalindrome(s, k, l+1, r-1)
	} else {
		if k > 0 {
			highestChar := max(s[l], s[r])
			if s[l] != highestChar {
				s = s[:l] + string(highestChar) + s[l+1:]
				k--
			}
			if s[r] != highestChar {
				s = s[:r] + string(highestChar) + s[r+1:]
				k--
			}
			s = maximizePalindrome(s, k, l+1, r-1)
		}
	}

	return s
}

// fungsi untuk mendapatkan karakter maksimum dari dua karakter
func max(a, b byte) byte {
	if a > b {
		return a
	}
	return b
}

// fungsi utama untuk mendapatkan highest palindrome
func highestPalindrome(s string, k int) string {
	n := len(s)
	// buat palindrome dengan perubahan minimal
	s, remainingChanges := makePalindrome(s, k, 0, n-1)

	// jika tidak bisa jadi palindrome, return -1
	if !isPalindrome(s) {
		return "-1"
	}

	// tingkatkan palindrome menjadi nilai tertinggi
	return maximizePalindrome(s, remainingChanges, 0, n-1)
}

func main() {
	fmt.Println(highestPalindrome("3943", 1))        
	fmt.Println(highestPalindrome("932239", 2))      
	fmt.Println(highestPalindrome("12321", 1))       
	fmt.Println(highestPalindrome("12345", 3))       
	fmt.Println(highestPalindrome("3943", 0))        
	fmt.Println(highestPalindrome("3943", 4))        
}
