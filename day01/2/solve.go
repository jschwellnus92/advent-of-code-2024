package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const delimiter = "   "

func main() {
	// open input file
	file, err := os.Open("input.txt")
	defer file.Close()
	if err != nil {
		panic(err)
	}

	var listOne, listTwo []int

	// buffer file line by line
	// add segments split by delim to arrays
	r := bufio.NewReader(file)
	for {
		line, _, err := r.ReadLine()
		if len(line) > 0 {
			segments := strings.Split(string(line), delimiter)
			listOne = addSegmentToArray(segments[0], listOne)
			listTwo = addSegmentToArray(segments[1], listTwo)
		}

		if err != nil {
			break
		}
	}

	// calculate similarity score
	var totalSimilarityScore = 0
	for _, element := range listOne {
		count := count(listTwo, func(x int) bool { return x == element })
		if count == 0 {
			continue
		}

		totalSimilarityScore += element * count
	}

	// the answer
	fmt.Println(totalSimilarityScore)
}

func addSegmentToArray(segment string, array []int) []int {
	i, err := strconv.Atoi(segment)
	if err != nil {
		panic(err)
	}

	return append(array, i)
}

// return count of f func(T) bool returning true when passing in values from slice[]T to func
func count[T any](slice []T, f func(T) bool) int {
	count := 0
	for _, s := range slice {
		if f(s) {
			count++
		}
	}

	return count
}
