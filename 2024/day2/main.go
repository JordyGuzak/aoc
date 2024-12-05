package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	safeCount := 0

	file, err := os.Open("input.txt")

	if err != nil {
		fmt.Println("Failed to read input file", err)
		return
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		levels, err := mapToInts(strings.Fields(line))

		if err != nil {
			fmt.Println("Failed to convert integer:", err)
			return
		}

		if isSafe(levels) {
			safeCount++
		} else {
			for i := range levels {
				if isSafe(remove(levels, i)) {
					safeCount++
					break
				}
			}
		}
	}

	fmt.Println("Safe records: ", safeCount)
}

func mapToInts(stringArray []string) ([]int, error) {
	result := make([]int, len(stringArray))

	for i, str := range stringArray {
		num, err := strconv.Atoi(str)

		if err != nil {
			return nil, errors.New("Failed to convert to int: " + str)
		}

		result[i] = num
	}

	return result, nil
}

func isSafe(levels []int) bool {
	const min = 1
	const max = 3

	dir := 0
	for i := 1; i < len(levels); i++ {
		diff, newDir := getDifferenceAndDirection(levels[i], levels[i-1])

		if dir == 0 {
			dir = newDir
		} else if dir != newDir {
			return false
		}

		if diff < min {
			return false
		}

		if diff > max {
			return false
		}
	}

	return true
}

func getDifferenceAndDirection(a int, b int) (int, int) {
	diff := a - b
	direction := 1
	if diff < 0 {
		direction = -1
	}

	return abs(diff), direction
}

func abs(x int) int {
	if x < 0 {
		return -x
	}

	return x
}

func remove(slice []int, index int) []int {
	var newSlice []int
	if index > 0 {
		newSlice = append(newSlice, slice[:index]...)
	}
	newSlice = append(newSlice, slice[index+1:]...)
	return newSlice
}
