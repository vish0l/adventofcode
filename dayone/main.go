package main

import (
	"bufio"
	"fmt"
	"log/slog"
	"math"
	"os"
	"sort"
)

const inputFileName = "input.txt"

func main() {
	inputLines, err := readInputFile(inputFileName)
	if err != nil {
		slog.Error("Error reading input file", "error", err.Error())
		return
	}

	listOne := make([]int, len(inputLines))
	listTwo := make([]int, len(inputLines))

	for i, v := range inputLines {
		_, err := fmt.Sscanf(v, "%d %d", &listOne[i], &listTwo[i])
		if err != nil {
			return
		}
	}

	fmt.Println("Part One Result:", CalculateMinDifferenceSum(listOne, listTwo))
	fmt.Println("Part Two Result:", CalculateSumWithOccurrences(listOne, listTwo))
}

func CalculateMinDifferenceSum(listOne, listTwo []int) float64 {
	if !sort.IntsAreSorted(listOne) {
		sort.Ints(listOne)
	}

	if !sort.IntsAreSorted(listTwo) {
		sort.Ints(listTwo)
	}

	var results float64
	for i := 0; i < len(listOne); i++ {
		results += math.Abs(float64(listOne[i] - listTwo[i]))
	}

	return results
}

func CalculateSumWithOccurrences(listOne, listTwo []int) int {
	listTwoMap := make(map[int]int)
	for _, v := range listTwo {
		listTwoMap[v]++
	}

	var results int

	for _, v := range listOne {
		results += listTwoMap[v] * v
	}

	return results
}

func readInputFile(fileName string) ([]string, error) {
	file, err := os.Open(fileName)
	if err != nil {
		slog.Error("Error opening file", "error", err.Error())
		return nil, err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			slog.Error("Error closing file", "error", err.Error())
		}
	}(file)

	var lines []string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		slog.Error("Error reading file", "error", err.Error())
		return nil, err
	}

	return lines, nil
}
