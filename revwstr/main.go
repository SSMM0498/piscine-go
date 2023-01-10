package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) == 2 {
		words := SplitWhiteSpaces(os.Args[1])
		for i := len(words) - 1; i >= 0; i-- {
			PrintStr(words[i])
			if i != 0 {
				z01.PrintRune(' ')
			}
		}
		z01.PrintRune('\n')
	}
}

func PrintStr(s string) {
	for _, char := range s {
		z01.PrintRune(char)
	}
}

func SplitWhiteSpaces(s string) []string {
	arr := []string(nil)
	lastWhiteSpaceIndex := 0
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' {
			if lastWhiteSpaceIndex != i {
				arr = append(arr, s[lastWhiteSpaceIndex:i])
			}
			lastWhiteSpaceIndex = i + 1
		} else if i == len(s)-1 {
			arr = append(arr, s[lastWhiteSpaceIndex:i+1])
			lastWhiteSpaceIndex = i + 1
		}
	}
	return arr
}
