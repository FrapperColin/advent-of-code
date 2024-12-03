package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	// Open the text file
	file, err := os.Open("../inputs.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Read the content from the file
	fileInfo, err := os.Stat("../inputs.txt")
	if err != nil {
		fmt.Println("Error reading file info:", err)
		return
	}
	size := fileInfo.Size()
	data := make([]byte, size)
	_, err = file.Read(data)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	text := string(data)

	// Regular expressions
	mulPattern := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)
	dontPattern := regexp.MustCompile(`don't\(\)`)
	doPattern := regexp.MustCompile(`do\(\)`)

	// Track state using indices
	enabled := true
	sum := 0

	matches := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)|don't\(\)|do\(\)`).FindAllStringIndex(text, -1)

	for _, match := range matches {

		segment := text[match[0]:match[1]]

		if dontPattern.MatchString(segment) {
			enabled = false
		} else if doPattern.MatchString(segment) {
			enabled = true
		} else if enabled {
			fmt.Println("seg", segment)

			mulMatch := mulPattern.FindStringSubmatch(segment)

			num1, _ := strconv.Atoi(mulMatch[1])
			num2, _ := strconv.Atoi(mulMatch[2])
			sum += num1 * num2

		}
	}

	// Print the total sum
	fmt.Println(sum)

}
