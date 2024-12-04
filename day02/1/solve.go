package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// open input file
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// buffer file line by line
	var safeCount = 0
	r := bufio.NewReader(file)
	for {
		line, _, err := r.ReadLine()
		if len(line) > 0 {
			if isReportSafe(string(line)) {
				safeCount++
			}
		}

		if err != nil {
			break
		}
	}

	// the answer
	fmt.Println(safeCount)
}

func isReportSafe(levelStr string) bool {
	levels := strings.Fields(levelStr)
	var incrementing = false

	for i := 0; i < len(levels)-1; i++ {
		currLevel, _ := strconv.Atoi(levels[i])
		nextLevel, _ := strconv.Atoi(levels[i+1])
		if i == 0 {
			incrementing = nextLevel > currLevel
		}

		if incrementing && nextLevel < currLevel {
			return false
		}
		if !incrementing && nextLevel > currLevel {
			return false
		}

		levelDiff := currLevel - nextLevel
		if levelDiff < 0 {
			levelDiff = -levelDiff
		}

		if levelDiff == 0 || levelDiff > 3 {
			return false
		}
	}

	return true
}
