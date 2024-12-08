package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var total = 0
	regex, _ := regexp.Compile(`mul\(\d+,\d+\)`)
	r := bufio.NewReader(file)
	for {
		line, _, err := r.ReadLine()
		if err != nil {
			break
		}

		if len(line) > 0 {
			matches := regex.FindAllString(string(line), -1)
			for match := range matches {

			}
		}
	}

	// the answer
	fmt.Println(total)
}

func processMul(mul string) {

}
