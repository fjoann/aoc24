package main

import (
	"encoding/csv"
	"log"
	"os"
	"sort"
	"strconv"
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

	var totalDistance int

	for i := 0; i < len(groupOneList); i++ {
		elementDistance := abs(groupOneList[i] - groupTwoList[i])
		totalDistance += elementDistance
	}

	log.Println(totalDistance)
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

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
