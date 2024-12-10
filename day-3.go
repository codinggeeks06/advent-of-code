package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"regexp"
)

func main() {
	// Open the file
	file, err := os.Open("input-text-3.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	// Loop through each line of the file and parse the integers
	count := 0
	isMulEnabled := true
	for scanner.Scan() {
		line := scanner.Text()
		// Regex pattern to match mul(x, y) or similar variants
		re := regexp.MustCompile(`mul\((\d+),(\d+)\)|do\(\)|don't\(\)`)

		// Find all matches
		matches := re.FindAllString(line, -1)

		// Output all matches
		for _, match := range matches {

			if match == "do()" {
				isMulEnabled = true
			}
			if match == "don't()" {
				isMulEnabled = false
			}
			if isMulEnabled{
				ans, _ := multiplication(match)
				count += ans
			}
		}
	}
	fmt.Println(count)
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}
}

func multiplication(str string) (int, error) {
	// Remove "mul(" at the start and ")" at the end
	str = strings.TrimPrefix(str, "mul(")
	str = strings.TrimSuffix(str, ")")

	// Split the string by the comma to separate the numbers
	parts := strings.Split(str, ",")

	// Convert the parts into integers
	if len(parts) != 2 {
		return 0, fmt.Errorf("invalid format: expected two numbers separated by a comma")
	}

	// Convert the strings to integers
	num1, err := strconv.Atoi(parts[0])
	if err != nil {
		return 0, fmt.Errorf("error converting first number: %v", err)
	}

	num2, err := strconv.Atoi(parts[1])
	if err != nil {
		return 0, fmt.Errorf("error converting second number: %v", err)
	}

	// Return the product of the two numbers
	return num1 * num2, nil
}