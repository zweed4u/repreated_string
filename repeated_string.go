package main

import (
	"fmt"
	"strings"
)

func repeatedString(s string, n int) (aCount int) {
	if s == "a" {
		return n
	}
	if !strings.Contains(s, "a") {
		return 0
	}
	lengthOfString := len(s)
	for _, letter := range s {
		if string(letter) == "a" {
			aCount += 1
		}
	}
	aCount *= n / lengthOfString
	leftover := n % lengthOfString
	remaining := s[:leftover]
	for _, letter := range remaining {
		if string(letter) == "a" {
			aCount += 1
		}
	}
	fmt.Println(aCount)
	return
}

func main() {
	repeatedString("a", 1000000000000)
	repeatedString("aba", 10)
	repeatedString("abcac", 10)
}
