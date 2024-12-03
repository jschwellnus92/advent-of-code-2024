package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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

	// sort the two arrays
	sort.Ints(listOne[:])
	sort.Ints(listTwo[:])

	// calculate difference between values in sorted arrays
	var totalDifference = 0
	for index, element := range listOne {
		difference := element - listTwo[index]
		if difference < 0 {
			difference = -difference
		}

		totalDifference += difference
	}

	// the answer
	fmt.Println(totalDifference)
}

func addSegmentToArray(segment string, array []int) []int {
	i, err := strconv.Atoi(segment)
	if err != nil {
		panic(err)
	}

	return append(array, i)
}
