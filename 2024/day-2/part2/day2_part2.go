package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// Open the file
	file, err := os.Open("../inputs.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Scanner to read the file line by line
	scanner := bufio.NewScanner(file)

	var lines [][]int
	lineNumber := 0

	// Read and parse the file into a list of lines
	for scanner.Scan() {
		lineNumber++
		line := scanner.Text()
		stringNumbers := strings.Fields(line)

		var numbers []int
		for _, strNum := range stringNumbers {
			num, err := strconv.Atoi(strNum)
			if err != nil {
				fmt.Printf("Error converting string to number at line %d: %v\n", lineNumber, err)
				continue
			}
			numbers = append(numbers, num)
		}

		lines = append(lines, numbers)
	}

	correctLines := 0

	for i := 0; i < len(lines); i++ {
		if validate(lines[i]) || validateByRemoving(lines[i]) {
			correctLines++
		}
	}

	fmt.Printf("\nTotal Correct Lines: %d\n", correctLines)

	// Check for errors during scanning
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}
}

func validateByRemoving(data []int) bool {
	for i := 0; i < len(data); i++ {
		data1 := append([]int{}, data[:i]...)
		data1 = append(data1, data[i+1:]...)
		if validate(data1) {
			return true
		}
	}
	return false
}

func validate(data []int) bool {
	desc := AllDecreasing(data)
	incr := AllIncreasing(data)
	diff := diff(data, 3, 1)
	return (desc || incr) && diff
}

func AllDecreasing(input []int) bool {
	for i := 1; i < len(input); i++ {
		if input[i] > input[i-1] {
			return false
		}
	}
	return true
}

func AllIncreasing(input []int) bool {
	for i := 1; i < len(input); i++ {
		if input[i] < input[i-1] {
			return false
		}
	}
	return true
}
func diff(input []int, max, min int) bool {
	for i := 1; i < len(input); i++ {
		if absValue(input[i]-input[i-1]) > max || absValue(input[i]-input[i-1]) < min {
			return false
		}
	}
	return true
}

func absValue(a int) int {
	if a < 0 {
		return -1 * a
	}
	return a
}