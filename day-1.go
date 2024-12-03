package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"math"
)

func main() {
	// Open the file
	file, err := os.Open("input-text.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Create a scanner to read the file line by line
	var dataA []int
	var dataB []int
	scanner := bufio.NewScanner(file)

	// Loop through each line of the file and parse the integers
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Fields(line) // Split by whitespace (spaces or tabs)

		if len(parts) == 2 {
			a, err1 := strconv.Atoi(parts[0])
			b, err2 := strconv.Atoi(parts[1])

			if err1 != nil || err2 != nil {
				fmt.Println("Error parsing numbers:", err1, err2)
				continue
			}

			// Append the parsed values as a pair of integers
			dataA = append(dataA, a)
			dataB = append(dataB, b)
		} else {
			fmt.Println("Invalid line format")
		}
	}

	// Sort the data first by the first column, then by the second column
	sort.Ints(dataA)
	sort.Ints(dataB)
	count := 0
	// Print the sorted data
	for i := 0; i < len(dataA) ; i++{
		count += int(math.Abs(float64(dataA[i] - dataB[i])))
	}
	fmt.Println("ans", count)

	// Check for scanning errors
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}
}
