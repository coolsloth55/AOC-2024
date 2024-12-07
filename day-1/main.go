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

	arr1, arr2 := loadinput(f)

	sortMethod := os.Getenv("SORT_METHOD")
	distance := getDistance(arr1, arr2, sortMethod)

	fmt.Printf("Distance: %d\n", distance)
}

func getDistance(arr1 []int, arr2 []int, method string) int {
	if method == "counting" {

	} else {
		bubblesort(arr1)
		bubblesort(arr2)
	}

	distance := 0
	for idx, val1 := range arr1 {
		distance += int(math.Abs(float64(val1 - arr2[idx])))
	}
	return distance
}

func bubblesort(arr []int) {
	for i := 1; i < len(arr); i++ {
		for j := 0; j < len(arr)-i; j++ {
			if arr[j+1] < arr[j] {
				arr[j+1], arr[j] = arr[j], arr[j+1]
			}
		}
	}
}

func loadinput(f string) ([]int, []int) {
	var arr1 []int
	var arr2 []int

	file, err := os.Open(f)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		values := strings.Split(scanner.Text(), "   ")
		first, _ := strconv.Atoi(values[0])
		sceond, _ := strconv.Atoi(values[1])
		arr1 = append(arr1, first)
		arr2 = append(arr2, sceond)
	}

	return arr1, arr2
}
