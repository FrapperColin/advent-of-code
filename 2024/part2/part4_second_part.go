package main

import (
	"bufio"
	"fmt"
	"os"
)

var diagonalPairs = [][2][2]int{
	{{-1, -1}, {1, 1}}, // Top-left to bottom-right
	{{-1, 1}, {1, -1}}, // Top-right to bottom-left
}

func main() {
	// Example grid
	grid := readGrid("../inputs.txt")
	word := "MAS"

	findXShape(grid, word)
}

func readGrid(filename string) [][]rune {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var grid [][]rune
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		grid = append(grid, []rune(line))
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}
	return grid
}

func findXShape(grid [][]rune, word string) {
	wordRunes := []rune(word)
	numberOfLine := len(grid)
	totalOccurence := 0

	for lineIndex := 0; lineIndex < numberOfLine; lineIndex++ {
		numberOfCharactersPerLine := len(grid[lineIndex])

		for characterIndex := 0; characterIndex < numberOfCharactersPerLine; characterIndex++ {
			// Check if the current cell matches the middle character of the word
			if grid[lineIndex][characterIndex] == wordRunes[1] {
				if isValidXShape(grid, lineIndex, characterIndex, wordRunes) {
					totalOccurence++
				}
			}
		}
	}
	fmt.Printf("Found X shape for %s, %d times \n", word, totalOccurence)

}

func isValidXShape(grid [][]rune, lineIndex, characterIndex int, wordRunes []rune) bool {
	for _, pair := range diagonalPairs {

		// Check both diagonals for the required letters
		first := getLetter(grid, lineIndex+pair[0][0], characterIndex+pair[0][1])
		second := getLetter(grid, lineIndex+pair[1][0], characterIndex+pair[1][1])

		if !((first == wordRunes[0] && second == wordRunes[2]) || (first == wordRunes[2] && second == wordRunes[0])) {
			return false
		}
	}
	return true
}

func getLetter(grid [][]rune, lineIndex, characterIndex int) rune {
	if lineIndex < 0 || characterIndex < 0 || lineIndex >= len(grid) || characterIndex >= len(grid[0]) {
		return ' '
	}
	return grid[lineIndex][characterIndex]
}
