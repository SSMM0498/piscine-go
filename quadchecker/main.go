package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	quads := []string{"quadA", "quadB", "quadC", "quadD", "quadE"}
	// Récupérer le output de la derniere commande
	_lastOutput, _ := io.ReadAll(os.Stdin)
	lastOutput := string(_lastOutput)

	// Est ce que le dernier output n'est pas vide
	if len(lastOutput) != 0 {
		// Récupérer nombre max ligne et max colonne
		maxLine, maxColumn := getDimension(lastOutput)
		// Récupérer quad correspondant au dernier output
		matchingQuad := []string(nil)
		// On exécuter les quads 1 a 1 pour voir ce qui correspondent au lastoutput
		for _, quad := range quads {
			quadOutput := runCommand(quad, maxColumn, maxLine)
			if lastOutput == quadOutput {
				matchingQuad = append(matchingQuad, quad)
			}
		}
		// Est ce qu'il y'a au moins 1 quad qui match
		if len(matchingQuad) != 0 {
			// On parcours les quads qui match et on fait l'affichage
			for i, quad := range matchingQuad {
				fmt.Printf("[%v] [%v] [%v]", quad, maxColumn, maxLine)
				if i != len(matchingQuad)-1 {
					fmt.Print(" || ")
				} else {
					fmt.Print("\n")
				}
			}
		} else {
			// Si aucun quad ne correspond
			fmt.Println("Not a Raid function")
		}
	}
}

func getDimension(lastOutput string) (int, int) {
	maxLine := 0
	maxColumn := 0
	for i, char := range lastOutput {
		if char == '\n' {
			maxLine++
			if maxColumn == 0 {
				maxColumn = i
			}
		}
	}
	return maxLine, maxColumn
}

func runCommand(prg string, x, y int) string {
	str := ""
	if prg == "quadA" {
		str = QuadA(x, y)
	} else if prg == "quadB" {
		str = QuadB(x, y)
	} else if prg == "quadC" {
		str = QuadC(x, y)
	} else if prg == "quadD" {
		str = QuadD(x, y)
	} else if prg == "quadE" {
		str = QuadE(x, y)
	}
	return str
}

func drawQuad(
	maxColumn int,
	maxLine int,
	cornerTopRight rune,
	cornerTopLeft rune,
	cornerBottomRight rune,
	cornerBottomLeft rune,
	edgeLine rune,
	edgeColumn rune,
) string {
	if maxLine <= 0 || maxColumn <= 0 {
		return ""
	}
	whiteSpace := rune(' ')
	runes := []rune(nil)
	for line := 1; line <= maxLine; line++ {
		if line == 1 {
			runes = append(runes, drawRow(cornerTopLeft, cornerTopRight, edgeLine, maxColumn)...)
		} else if line == maxLine {
			runes = append(runes, drawRow(cornerBottomLeft, cornerBottomRight, edgeLine, maxColumn)...)
		} else {
			runes = append(runes, drawRow(edgeColumn, edgeColumn, whiteSpace, maxColumn)...)
		}
	}
	return string(runes)
}

func drawRow(
	start rune,
	end rune,
	middle rune,
	maxColumn int,
) []rune {
	runes := []rune(nil)
	for column := 1; column <= maxColumn; column++ {
		if column == 1 {
			runes = append(runes, start)
		} else if column == maxColumn {
			runes = append(runes, end)
		} else {
			runes = append(runes, middle)
		}
	}
	runes = append(runes, '\n')
	return runes
}

func QuadA(
	x int,
	y int,
) string {
	cornerTopRight := 'o'
	cornerTopLeft := cornerTopRight
	cornerBottomRight := cornerTopRight
	cornerBottomLeft := cornerTopRight
	edgeLine := '-'
	edgeColumn := '|'
	return drawQuad(x, y, cornerTopRight, cornerTopLeft, cornerBottomRight, cornerBottomLeft, edgeLine, edgeColumn)
}

func QuadB(
	x int,
	y int,
) string {
	cornerTopRight := '\\'
	cornerTopLeft := '/'
	cornerBottomRight := cornerTopLeft
	cornerBottomLeft := cornerTopRight
	edgeLine := '*'
	edgeColumn := edgeLine
	return drawQuad(x, y, cornerTopRight, cornerTopLeft, cornerBottomRight, cornerBottomLeft, edgeLine, edgeColumn)
}

func QuadC(
	x int,
	y int,
) string {
	cornerTopRight := 'A'
	cornerTopLeft := cornerTopRight
	cornerBottomRight := 'C'
	cornerBottomLeft := cornerBottomRight
	edgeLine := 'B'
	edgeColumn := edgeLine
	return drawQuad(x, y, cornerTopRight, cornerTopLeft, cornerBottomRight, cornerBottomLeft, edgeLine, edgeColumn)
}

func QuadD(
	x int,
	y int,
) string {
	cornerTopRight := 'C'
	cornerTopLeft := 'A'
	cornerBottomRight := cornerTopRight
	cornerBottomLeft := cornerTopLeft
	edgeLine := 'B'
	edgeColumn := edgeLine
	return drawQuad(x, y, cornerTopRight, cornerTopLeft, cornerBottomRight, cornerBottomLeft, edgeLine, edgeColumn)
}

func QuadE(
	x int,
	y int,
) string {
	cornerTopRight := 'C'
	cornerTopLeft := 'A'
	cornerBottomRight := cornerTopLeft
	cornerBottomLeft := cornerTopRight
	edgeLine := 'B'
	edgeColumn := edgeLine
	return drawQuad(x, y, cornerTopRight, cornerTopLeft, cornerBottomRight, cornerBottomLeft, edgeLine, edgeColumn)
}
