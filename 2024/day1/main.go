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
	file, err := os.Open("input.txt")

	if err != nil {
		fmt.Println("Failed to read input file", err)
		return
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	var lists [listCount][]int

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

	for _, list := range lists {
		sort.Ints(list)
	}

	distance := 0

	for i := 0; i < len(lists[0]); i++ {
		distance += abs(lists[0][i] - lists[1][i])
	}

	fmt.Println("distance:", distance)

	smap := make(map[int]int)

	for _, v := range lists[1] {
		count, exists := smap[v]

		if exists {
			smap[v] = count + 1
		} else {
			smap[v] = 1
		}
	}

	similarity := 0
	for _, v := range lists[0] {
		count, exists := smap[v]
		if exists {
			similarity += v * count
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
