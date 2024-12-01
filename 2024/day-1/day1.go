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

	// Create slices to hold column data
	var col1, col2 []int

	// Create a scanner to read the file line by line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		// Split the line into columns
		line := scanner.Text()
		columns := strings.Fields(line)

		// Ensure there are at least two columns
		if len(columns) < 2 {
			fmt.Println("Invalid line:", line)
			return
		}

		// Convert columns to integers
		val1, err1 := strconv.Atoi(columns[0])
		val2, err2 := strconv.Atoi(columns[1])
		if err1 != nil || err2 != nil {
			fmt.Println("Error parsing line:", line)
			return
		}

		// Append to slices
		col1 = append(col1, val1)
		col2 = append(col2, val2)
	}

	// Check for scanner errors
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	if !areSlicesSameSize(col1, col2) {
		fmt.Println("Error, not same list length")
		return
	}

	sortedCol1 := sortedList(col1)
	sortedCol2 := sortedList(col2)

	fmt.Println("totalDistance:", calculTotalDistance(sortedCol1, sortedCol2))
	fmt.Println("similarities:", calculSimilarities(sortedCol1, sortedCol2))

}

func areSlicesSameSize(slice1, slice2 []int) bool {
	return len(slice1) == len(slice2)
}

func sortedList(list []int) []int {
	sortedCol1 := make([]int, len(list))
	copy(sortedCol1, list)
	sort.Ints(sortedCol1)
	return sortedCol1
}

func calculTotalDistance(sortedCol1, sortedCol2 []int) int {
	totalDistance := 0
	for i := 0; i < len(sortedCol1) && i < len(sortedCol2); i++ {
		totalDistance += absoluteDifference(sortedCol1[i], sortedCol2[i])
	}
	return totalDistance
}

func calculSimilarities(sortedCol1, sortedCol2 []int) int {
	// there is also a different way to do it, maybe more optimized
	// what if in ur first list you take all occurences, mlitple by occurence of the second list
	// but you have to remove after you find it 
	// it's maybe not necessary to go though that list if it's ordered, smth with a binary search might exist
	similarities := 0
	for i := 0; i < len(sortedCol1) && i < len(sortedCol2); i++ {
		occurence := countOccurrences(sortedCol2, sortedCol1[i])
		similarities += sortedCol1[i] * occurence
	}
	return similarities
}

func countOccurrences(nums []int, n int) int {
	count := 0
	for _, num := range nums {
		if num == n {
			count++
		}
	}
	return count
}

func absoluteDifference(num1, num2 int) int {
	if num1 > num2 {
		return num1 - num2
	}
	return num2 - num1
}
