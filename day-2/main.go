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
	f := os.Getenv("INPUT_FILE")
	if f == "" {
		f = "input.txt"
	}

	matrix := loadinput(f)
	safe_reports := getSafeReports(matrix)
	fmt.Printf("Safe Reports: %d\n", safe_reports)
}

func getSafeReports(m [][]int) int {
	safe := 0

	for i := 0; i < len(m); i++ {
		gradual := true
		change := map[string]int{
			"increasing": 0,
			"decresing":  0,
			"noChange":   0,
		}

		for j := 0; j < len(m[i])-1; j++ {
			diff := m[i][j+1] - m[i][j]

			if diff > 0 {
				change["increasing"]++
			} else if diff == 0 {
				change["noChange"]++
			} else {
				change["decresing"]++
			}

			if math.Abs(float64(diff)) > 3 {
				gradual = false
			}
		}

		noSwitch := change["increasing"] > 0 && change["decresing"] == 0 || change["increasing"] == 0 && change["decresing"] > 0
		if gradual && noSwitch && change["noChange"] == 0 {
			safe++
		}
	}

	return safe
}

func isIncreasing(a int) bool {
	if a > 0 {
		return true
	}
	return false
}

func loadinput(f string) [][]int {
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
