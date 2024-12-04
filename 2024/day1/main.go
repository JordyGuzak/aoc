package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

const listCount int = 2

func main() {
	// Open input file
	file, err := os.Open("input.txt")

	if err != nil {
		fmt.Println("Failed to read input file", err)
		return
	}

	defer file.Close()

	// Read lines and create lists
	var lists [listCount][]int

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		trimmed := strings.TrimSpace(line)
		parts := strings.Fields(trimmed)

		for i := 0; i < len(parts) && i < len(lists); i++ {
			num, err := strconv.Atoi(parts[i])
			if err != nil {
				fmt.Println("Failed to parse int", err)
				return
			}

			lists[i] = append(lists[i], num)
		}
	}

	// Sort lists
	for _, list := range lists {
		sort.Ints(list)
	}

	// Part 1: Calculate total distance
	distance := 0

	for i := 0; i < len(lists[0]); i++ {
		distance += abs(lists[0][i] - lists[1][i])
	}

	fmt.Println("distance:", distance)

	// Part 2: Similarity score
	// Count occurrences
	occurrences := make(map[int]int)

	for _, value := range lists[1] {
		count, exists := occurrences[value]

		if exists {
			occurrences[value] = count + 1
		} else {
			occurrences[value] = 1
		}
	}

	// Calculate similarity score
	similarity := 0
	for _, value := range lists[0] {
		count, exists := occurrences[value]
		if exists {
			similarity += value * count
		}
	}

	fmt.Println("similarity:", similarity)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}

	return x
}
