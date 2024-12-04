package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var safeCount = 0
	r := bufio.NewReader(file)
	for {
		line, _, err := r.ReadLine()
		if len(line) > 0 {
			levels := strings.Fields(string(line))

			if isReportSafe(levels) || isDampenedReportSafe(levels) {
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

func isReportSafe(levels []string) bool {
	var incrementing = false

	for i := range levels {
		nextLevelIdx := i + 1
		if nextLevelIdx == len(levels) {
			break
		}

		currLevel, _ := strconv.Atoi(levels[i])
		nextLevel, _ := strconv.Atoi(levels[nextLevelIdx])

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

func isDampenedReportSafe(levels []string) bool {
	for i := range levels {
		var dampenedReport []string
		dampenedReport = append(dampenedReport, levels[:i]...)
		dampenedReport = append(dampenedReport, levels[i+1:]...)

		if isReportSafe(dampenedReport) {
			return true
		}
	}

	return false
}
