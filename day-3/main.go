package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	var result int

	part := os.Getenv("PART")
	if part == "two" {
		result = getByRegex("do\\(\\)mul\\(\\d+,\\d+\\)")
	} else {
		result = getByRegex("mul\\(\\d+,\\d+\\)")
	}

	fmt.Printf("result: %d\n", result)
}

func getByRegex(p string) int {
	input := loadInput()

	re := regexp.MustCompile(p)
	matches := re.FindAll([]byte(input), -1)

	count := 0
	for _, x := range matches {
		val := string(x)
		fmt.Printf("Found match: %s\n", val)
		count = count + compute(val)
	}

	return count
}

func compute(s string) int {
	var temp string
	temp = strings.ReplaceAll(s, "mul", "")
	temp = strings.ReplaceAll(temp, "(", "")
	temp = strings.ReplaceAll(temp, ")", "")

	arr := strings.Split(temp, ",")

	val1, err := strconv.Atoi(arr[0])
	if err != nil {
		panic(err)
	}

	val2, err := strconv.Atoi(arr[1])
	if err != nil {
		panic(err)
	}

	return val1 * val2
}

func loadInput() string {
	var input string

	f := os.Getenv("INPUT_FILE")
	if f == "" {
		f = "input.txt"
	}

	file, err := os.Open(f)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		input = input + scanner.Text()
	}

	return input
}
