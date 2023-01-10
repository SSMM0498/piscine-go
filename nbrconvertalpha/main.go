package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) >= 2 {
		arguments := os.Args[1:]
		isUpper := false
		if arguments[0] == "--upper" {
			isUpper = true
			arguments = arguments[1:]
		}
		for _, arg := range arguments {
			number := 0
			number = basicAtoi(arg)
			if number >= 1 && number <= 26 {
				if !isUpper {
					z01.PrintRune(rune(number + 96))
				} else {
					z01.PrintRune(rune(number + 64))
				}
			} else {
				z01.PrintRune(' ')
			}
		}
		z01.PrintRune('\n')
	}
}

func basicAtoi(str string) int {
	number := 0
	factor := 1
	for i := len(str) - 1; i >= 0; i-- {
		number += (int(rune(str[i])) - 48) * factor
		factor = factor * 10
	}
	return number
}
