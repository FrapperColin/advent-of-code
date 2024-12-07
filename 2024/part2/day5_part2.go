package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Main entry point
func main() {
	// Open the file
	file, err := os.Open("../inputs/inputs.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	// Parse ordering rules
	orderingRules := make(map[int][]int)
	isRulesPart := true
	var totalMiddleSum int // Variable to accumulate the sum of middle numbers
	totalMiddleSumOfFixedLines := 0
	// Parse rules and sequences
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			isRulesPart = false
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
			// Process sequences
			sequence := toIntSlice(strings.Split(line, ","))

			if isValid, middleSum := validateAndCalculateMiddleSum(sequence, orderingRules); isValid {
				totalMiddleSum += middleSum
			} else {
				fixedSequence := fixSequenceWithTopologicalSort(sequence, orderingRules)
				if fixedSequence != nil {
					middleSum := calculateMiddleSum(*fixedSequence)
					totalMiddleSumOfFixedLines += middleSum
				} else {
					fmt.Printf("Failed to fix sequence: %v\n", sequence)
				}
			}
		}
	}

	// Output the total sum after processing all sequences
	fmt.Printf("Total sum of middle numbers (odd-length sequences only): %d\n", totalMiddleSum)
	fmt.Printf("Total sum of middle numbers (fixed lines and odd-length sequences only): %d\n", totalMiddleSumOfFixedLines)

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}
}

func validateAndCalculateMiddleSum(sequence []int, orderingRules map[int][]int) (bool, int) {
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

// Helper to calculate the middle number for odd-length sequences
func calculateMiddleSum(sequence []int) int {
	if len(sequence)%2 == 1 {
		return sequence[len(sequence)/2]
	}
	return 0
}

// Helper to fix invalid sequence using topological sort
func fixSequenceWithTopologicalSort(sequence []int, rules map[int][]int) *([]int) {
	// Create the graph BASED ON SEQUENCE
	graph, inDegree := buildSubGraph(sequence, rules)

	fmt.Printf("graph: %v\n", graph)
	fmt.Printf("degree: %v\n", inDegree)

	// Perform Kahn's algorithm for topological sort
	// Attempt to find an initial node even if no degree == 0 exists
	// Queue to store nodes with in-degree 0
	queue := []int{}

	occurrences := make(map[int]int)
	for _, num := range sequence {
		occurrences[num]++
	}


	// Add all nodes with in-degree 0 to the queue
	for _, node := range sequence {
		if inDegree[node] == 0 {
			queue = append(queue, node)
		}
	}

	sorted := []int{} // Final sorted order

	// Process nodes iteratively
	for len(queue) > 0 {
		fmt.Printf("queue inside list: %v\n", queue)

		// Pop the first node from the queue
		current := queue[0]
		queue = queue[1:]

		// Add current node to sorted order
		sorted = append(sorted, current)

		// Reduce the in-degree of dependent nodes
		for _, neighbor := range graph[current] {
			fmt.Printf("decreasing : %v, %v\n", neighbor, inDegree[neighbor])
			inDegree[neighbor]-- // Decrease degree by 1
			// If in-degree becomes 0, add to queue
			if inDegree[neighbor] == 0 {
				for i := 0; i < occurrences[neighbor]; i++ {
					queue = append(queue, neighbor)
				}

			}
		}

		fmt.Printf("queue after : %v\n", queue)

	}

	// If we processed all nodes, the sort is valid
	if len(sorted) == len(sequence) {
		fmt.Printf("sorted: %v\n", sorted)

		return &sorted
	}
	fmt.Printf("sorted: %v\n", sorted)

	return nil // Topological sort failed

}

func buildSubGraph(sequence []int, rules map[int][]int) (map[int][]int, map[int]int) {
	graph := make(map[int][]int)
	inDegree := make(map[int]int)
	fmt.Printf("rules : %v\n", rules)

	// Count the occurrences of each number in the sequence
	occurrences := make(map[int]int)
	for _, num := range sequence {
		occurrences[num]++
	}

	fmt.Printf("occurrences : %v\n", occurrences)

	// Initialize in-degree for every number in the sequence
	for num := range occurrences {
		inDegree[num] = 0
	}

	fmt.Printf("inDegree : %v\n", inDegree)

	// Build the graph based on rules and respect repeated dependencies
	for _, num := range sequence {
		fmt.Printf("num : %v\n", num)
		fmt.Printf("rules[num] : %v\n", rules[num])

		if neighbors, exists := rules[num]; exists {
			// For every dependency repeated, map it correctly
			for _, neighbor := range neighbors {
				for i := 0; i < occurrences[neighbor]; i++ {
					graph[num] = append(graph[num], neighbor)

					inDegree[neighbor]++
				}
			}
		}
	}

	return graph, inDegree
}

func contains(slice []int, num int) bool {
	for _, v := range slice {
		if v == num {
			return true
		}
	}
	return false
}

// Helper to convert a string to an integer
func toInt(s string) int {
	var num int
	fmt.Sscanf(s, "%d", &num)
	return num
}

// Helper to convert a slice of strings to a slice of integers
func toIntSlice(strings []string) []int {
	result := []int{}
	for _, s := range strings {
		result = append(result, toInt(s))
	}
	return result
}
