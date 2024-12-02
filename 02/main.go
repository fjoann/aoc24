package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"

	"github.com/fjoann/aoc24/aoc"
)

func main() {
	file, err := os.Open("02/input")
	if err != nil {
		log.Panic(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.Comma = ' '
	reader.FieldsPerRecord = -1

	records, err := reader.ReadAll()
	if err != nil {
		log.Panic(err)
	}

	var safeReports int
	var safeReportsWithProblemDampener int

	for _, record := range records {
		levels := make([]int, len(record))
		for i, val := range record {
			levels[i], err = strconv.Atoi(val)
			if err != nil {
				log.Panic(err)
			}
		}

		if isSafeReport(levels) {
			safeReports++
		}
		if isSafeReportWithProblemDampener(levels) {
			safeReportsWithProblemDampener++
		}
	}

	fmt.Printf("Safe reports: %d\n", safeReports)
	fmt.Printf("Safe reports with Problem Dampener: %d\n", safeReportsWithProblemDampener)
}

func isSafeReport(levels []int) bool {
	isIncreasing := true
	isDecreasing := true

	for i := 0; i < len(levels)-1; i++ {
		diff := levels[i+1] - levels[i]
		absDiff := aoc.AbsInt(diff)

		if absDiff < 1 || absDiff > 3 {
			return false
		}

		switch {
		case diff < 0:
			isIncreasing = false
		case diff > 0:
			isDecreasing = false
		}
	}

	return isIncreasing != isDecreasing
}

func isSafeReportWithProblemDampener(levels []int) bool {
	for i := 0; i < len(levels); i++ {
		adjustedLevels := slices.Delete(slices.Clone(levels), i, i+1)
		if isSafeReport(adjustedLevels) {
			return true
		}
	}

	return false
}
