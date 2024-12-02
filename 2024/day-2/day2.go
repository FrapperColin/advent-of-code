package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	// File name
	// fileName := "inputs_copy.txt"
	fileName := "inputs.txt"

	// Open the file
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Scanner to read the file line by line
	scanner := bufio.NewScanner(file)

	lineNumber := 0
	totalValidLine := 0
	for scanner.Scan() {
		lineNumber++

		// Read and process the line
		line := scanner.Text()
		stringNumbers := strings.Fields(line)

		// Convert strings to integers
		var numbers []int
		for _, strNum := range stringNumbers {
			num, err := strconv.Atoi(strNum)
			if err != nil {
				fmt.Printf("Error converting string to number at line %d: %v\n", lineNumber, err)
				continue
			}
			numbers = append(numbers, num)
		}

		// Make copies of the original slice for sorting
		ascSorted := make([]int, len(numbers))
		descSorted := make([]int, len(numbers))
		copy(ascSorted, numbers)
		copy(descSorted, numbers)

		// Sort the copies
		sort.Ints(ascSorted)                               // Ascending
		sort.Sort(sort.Reverse(sort.IntSlice(descSorted))) // Descending

		// Check if they match the original line
		isAscMatch := slicesEqual(numbers, ascSorted)
		isDescMatch := slicesEqual(numbers, descSorted)

		// Print the results
		if (isAscMatch || isDescMatch) && checkAdjacentDifferences(numbers, 3) {
			totalValidLine++
		}
	}
	fmt.Printf("total line %d", totalValidLine)

	// Check for errors during scanning
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

	// remove list if ordered list desc or asc is not corresponding

}

func slicesEqual(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func checkAdjacentDifferences(numbers []int, diffMax int) bool {
	tolerateError := 0
	for i := 0; i < len(numbers)-1; i++ {
		diff := numbers[i+1] - numbers[i]
		if diff == 0 || diff > diffMax || diff < -diffMax {
			tolerateError++
		}
		if tolerateError > 1 {
			return false
		}
	}
	return true
}

func countDifferences(original, sorted []int) int {
    if len(original) != len(sorted) {
        return len(original) // If lengths differ, all elements are mismatches
    }
    differences := 0
    for i := range original {
        if original[i] != sorted[i] {
            differences++
        }
    }
    return differences
}
