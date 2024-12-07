package main

import (
	"bufio"
	"fmt"
	"os"
)

// Directions for traversal: dx, dy
var directions = [8][2]int{
	{0, 1}, {1, 0}, {1, 1}, {-1, 1}, // Right, Down, Diagonal Down-Right, Diagonal Up-Right
	{0, -1}, {-1, 0}, {-1, -1}, {1, -1}, // Left, Up, Diagonal Up-Left, Diagonal Down-Left
}

func main() {
	// Read the file into a 2D array
	grid := readGrid("inputs.txt")
	word := "XMAS"

	// Find the word in the grid
	findWord(grid, word)
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

func findWord(grid [][]rune, word string) {
	wordRunes := []rune(word)
	fmt.Printf("rune %d \n", wordRunes)

	numberOfLine := len(grid)

	totalOccurence := 0
	for lineIndex := 0; lineIndex < numberOfLine; lineIndex++ {
		numberOfCharactersPerLine := len(grid[lineIndex])
		for characterIndex := 0; characterIndex < numberOfCharactersPerLine; characterIndex++ {
			for _, dir := range directions {
				if checkDirection(grid, lineIndex, characterIndex, dir[0], dir[1], wordRunes) {
					fmt.Printf("Found %s at (%d, %d) direction (%d, %d)\n", word, lineIndex, characterIndex, dir[0], dir[1])
					totalOccurence++
				}
			}
		}
	}
	fmt.Printf("total occurence of %s %d \n", word, totalOccurence)

}

func checkDirection(grid [][]rune, lineIndex, characterIndex, dx, dy int, wordRunes []rune) bool {
	runesLength := len(wordRunes)

	for k := 0; k < runesLength; k++ {
		positionX, positionY := lineIndex+dx*k, characterIndex+dy*k
		if positionStillInGrid(grid, positionX, positionY) || grid[positionX][positionY] != wordRunes[k] {
			return false
		}
	}
	return true
}

func positionStillInGrid(grid [][]rune, positionX, positionY int) bool {
	return positionX < 0 || positionY < 0 || positionX >= len(grid) || positionY >= len(grid[0])
}