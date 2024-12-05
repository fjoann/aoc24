package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

func main() {
	corruptedMemory, err := os.ReadFile("03/input")
	if err != nil {
		log.Fatal(err)
	}

	multiplyInstruction := regexp.MustCompile(`mul\((\d+),(\d+)\)`)
	sanitizedMemory := multiplyInstruction.FindAllStringSubmatch(string(corruptedMemory), -1)

	// Part 1: sum of multiplications
	result, err := multiplyAndSum(sanitizedMemory)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Sum of multiplications: %d\n", result)

	// Part 2: sum of enabled multiplications
	enabledInstruction := regexp.MustCompile(`(?s)(^|do\(\))(.*?)(don't\(\)|$)`)
	enabledMemory := enabledInstruction.FindAllString(string(corruptedMemory), -1)

	var sanitizedMemoryEnabled [][]string
	for _, match := range enabledMemory {
		sanitizedMemoryPart := multiplyInstruction.FindAllStringSubmatch(string(match), -1)
		sanitizedMemoryEnabled = append(sanitizedMemoryEnabled, sanitizedMemoryPart...)
	}

	resultEnabled, err := multiplyAndSum(sanitizedMemoryEnabled)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Sum of enabled multiplications: %d\n", resultEnabled)
}

func multiplyAndSum(sanitizedMemory [][]string) (int, error) {
	var totalSum int

	for _, instruction := range sanitizedMemory {
		leftOperand, err := strconv.Atoi(instruction[1])
		if err != nil {
			return 0, err
		}
		rightOperand, err := strconv.Atoi(instruction[2])
		if err != nil {
			return 0, err
		}
		totalSum += leftOperand * rightOperand
	}

	return totalSum, nil
}
