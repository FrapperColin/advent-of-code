package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	file, err := os.Open("../inputs.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var text string
	for scanner.Scan() {
		text += scanner.Text()
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	pattern := `mul\((\d{1,3}),(\d{1,3})\)`

	re := regexp.MustCompile(pattern)

	matches := re.FindAllStringSubmatch(text, -1)

	sum := 0

	for _, match := range matches {
		// Convert captures to integers
		num1, _ := strconv.Atoi(match[1])
		num2, _ := strconv.Atoi(match[2])
		sum += num1 * num2
	}

	fmt.Println(sum)
}
