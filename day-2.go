package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"math"
)

func main() {
	// Open the file
	file, err := os.Open("input-text-2.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Create a scanner to read the file line by line
	var data [][]int
	scanner := bufio.NewScanner(file)

	// Loop through each line of the file and parse the integers
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Fields(line) // Split by whitespace (spaces or tabs)
		var numbers []int

		// Convert each part to an integer
		for _, part := range parts {
			num, err := strconv.Atoi(part)
			if err != nil {
				fmt.Println("Error converting string to integer:", err)
				continue
			}
			numbers = append(numbers, num)
		}
		data = append(data, numbers)
	}

	// day 2
	checkIfSafe(data)
	// Check for scanning errors
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}
}

func checkIfSafe(data [][]int) {
	count := 0
	for _,part := range data{
		if isSafe(part) {
			count++;
		}
	}
	fmt.Println("count", count)
}

func isSafe(report []int) bool {
	// Flag to check if all differences are between 1 and 3 and the sequence is monotonic
	isIncreasing := true
	isDecreasing := true
	// Iterate over the report's levels and check the conditions
	for i := 1; i < len(report); i++ {
		diff := report[i] - report[i-1]
		// Check if the difference is outside the allowed range [1, 3]
		if math.Abs(float64(diff)) == 0 || math.Abs(float64(diff)) > 3   {
			return false
		}

		// Track if the sequence is increasing or decreasing
		if diff > 0 {
			isDecreasing = false
		}
		if diff < 0 {
			isIncreasing = false
		}
	}
	return isIncreasing || isDecreasing
}