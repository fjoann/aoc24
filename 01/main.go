package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"

	"github.com/fjoann/aoc24/aoc"
)

func main() {
	file, err := os.Open("01/input")
	if err != nil {
		log.Panic(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.Comma = ' '
	reader.TrimLeadingSpace = true

	records, err := reader.ReadAll()
	if err != nil {
		log.Panic(err)
	}

	groupOneList := buildList(records, 0)
	groupTwoList := buildList(records, 1)

	// Part 1: total distance
	var totalDistance int

	for i := 0; i < len(groupOneList); i++ {
		elementDistance := aoc.AbsInt(groupOneList[i] - groupTwoList[i])
		totalDistance += elementDistance
	}

	fmt.Printf("Total distance: %d\n", totalDistance)

	// Part 2: similarity score
	var similarityScore int

	for _, g1Val := range groupOneList {
		g2Count := 0

		for _, g2Val := range groupTwoList {
			if g1Val == g2Val {
				g2Count++
			}
		}

		elementSimilarity := g1Val * g2Count
		similarityScore += elementSimilarity
	}

	fmt.Printf("Similarity score: %d\n", similarityScore)
}

func buildList(records [][]string, colIdx int) []int {
	var list []int

	for _, record := range records {
		val, err := strconv.Atoi(record[colIdx])
		if err != nil {
			log.Panic(err)
		}
		list = append(list, val)
	}

	sort.Ints(list)

	return list
}
