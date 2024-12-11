package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	matrix := loadinput()
	safe_reports := getSafeReports(matrix)
	fmt.Printf("Safe Reports: %d\n", safe_reports)
}

func getSafeReports(m [][]int) int {
	safe := 0

	for i := 0; i < len(m); i++ {
		if isSafe(m[i]) {
			safe++
		} else if os.Getenv("PART") == "two" {
			if isCleanedSafe(m[i]) {
				safe++
			}
		}
	}

	return safe
}

func isCleanedSafe(arr []int) bool {
	safe := false

	copiedArr := make([]int, len(arr))

	for i := 0; i < len(copiedArr); i++ {
		copy(copiedArr, arr)
		newArr := remove(copiedArr, i)
		if isSafe(newArr) {
			safe = true
		}
	}

	return safe
}

func remove(slice []int, s int) []int {
	return append(slice[:s], slice[s+1:]...)
}

func isSafe(arr []int) bool {
	isIncreasing, isDecresing := false, false

	for j := 1; j < len(arr); j++ {
		diff := arr[j] - arr[j-1]

		if math.Abs(float64(diff)) > 3 {
			return false
		}

		if diff > 0 {
			isIncreasing = true
		} else if diff < 0 {
			isDecresing = true
		} else {
			return false
		}

		if isIncreasing && isDecresing {
			return false
		}
	}

	return true
}

func loadinput() [][]int {
	f := os.Getenv("INPUT_FILE")
	if f == "" {
		f = "input.txt"
	}

	var matrix [][]int

	file, err := os.Open(f)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	lineIndex := 0
	for scanner.Scan() {
		var nums []int

		values := strings.Split(scanner.Text(), " ")
		for _, val := range values {
			num, err := strconv.Atoi(val)
			if err != nil {
				panic(err)
			}
			nums = append(nums, num)
		}
		matrix = append(matrix, nums)
		lineIndex++
	}

	return matrix
}
