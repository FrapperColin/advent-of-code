package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// Open the file
	file, err := os.Open("./inputs/inputs.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	// Step 1: Parse ordering rules
	orderingRules := make(map[int][]int)
	isRulesPart := true
	var totalMiddleSum int // Variable to accumulate the sum of middle numbers

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			isRulesPart = false // Empty line indicates the end of rules
			continue
		}

		if isRulesPart {
			// Parse ordering rule "4|3"
			parts := strings.Split(line, "|")
			if len(parts) != 2 {
				fmt.Println("Invalid rule:", line)
				continue
			}
			from := toInt(parts[0])
			to := toInt(parts[1])
			orderingRules[from] = append(orderingRules[from], to)
		} else {
			if isValid, middleSum := validateAndCalculateMiddleSum(line, orderingRules); isValid {
				totalMiddleSum += middleSum
			}
		}
	}

	fmt.Printf("Total sum of middle numbers (odd-length sequences only): %d\n", totalMiddleSum)

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}
}

// Helper function to validate a sequence and calculate the middle sum if valid
func validateAndCalculateMiddleSum(line string, orderingRules map[int][]int) (bool, int) {
	sequence := toIntSlice(strings.Split(line, ","))
	seen := make(map[int]bool)

	for _, num := range sequence {
		seen[num] = true
		if constraints, exists := orderingRules[num]; exists {
			for _, mustComeAfter := range constraints {
				if seen[mustComeAfter] {
					return false, 0 // Invalid sequence
				}
			}
		}
	}

	// Sequence is valid, calculate the middle sum for odd-length sequences
	if len(sequence)%2 == 1 {
		middleSum := calculateMiddleSum(sequence)
		return true, middleSum
	}

	return true, 0 // Even-length sequences contribute 0 to the sum
}

// Helper function to calculate the sum of the middle number(s) for odd-length sequences
func calculateMiddleSum(sequence []int) int {
	length := len(sequence)
	return sequence[length/2] // Return the single middle number for odd-length sequences
}

// Helper function to convert a string to an integer
func toInt(s string) int {
	var num int
	fmt.Sscanf(s, "%d", &num)
	return num
}

// Helper function to convert a slice of strings to a slice of integers
func toIntSlice(strings []string) []int {
	result := []int{}
	for _, s := range strings {
		result = append(result, toInt(s))
	}
	return result
}
