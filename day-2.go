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
	file, err := os.Open("mod.txt")
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
		isIncreasing := true
		isDecreasing := true

		// Check both conditions
		for i := 0; i < len(part)-1; i++ {
			// Condition 1: Check if adjacent elements differ by at least 1 and at most 3
			diff := math.Abs(float64(part[i+1] - part[i]))
			if diff < 1 || diff > 3 {
				break // If any adjacent difference is out of bounds, return false
			}
	
			// Condition 2: Check if the sequence is increasing or decreasing
			if part[i+1] > part[i] {
				isDecreasing = false
			} else if part[i+1] < part[i] {
				isIncreasing = false
			}
		}
		if isDecreasing || isIncreasing {
			count++
		}
	}
	fmt.Print("count", count)
}