package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) == 3 {
		s := noDoubleStr(os.Args[1])
		str := noDoubleStr(os.Args[2])
		for _, char := range s {
			for _, actual := range str {
				if char == actual {
					z01.PrintRune(char)
				}
			}
		}
		z01.PrintRune('\n')
	}
}

func noDoubleStr(s string) string {
	noOccur := []rune(nil)
	for _, char := range s {
		isFound := false
		for _, c := range noOccur {
			if char == c {
				isFound = true
			}
		}
		if !isFound {
			noOccur = append(noOccur, char)
		}
	}
	return string(noOccur)
}
