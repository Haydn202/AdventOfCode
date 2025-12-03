package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	ranges := readData("data.txt")
	result := sumInvalid(ranges)
	fmt.Println(result)
}

func sumInvalid(ranges []Range) int {
	sum := 0

	for _, r := range ranges {
		for i := r.bottom; i <= r.top; i++ {
			if isInvalidPart2(i) {
				sum += i
			}
		}
	}

	return sum
}

func isInvalidPart1(num int) bool {
	s := strconv.Itoa(num)
	n := len(s)

	if n%2 != 0 {
		return false
	}

	half := n / 2
	firstHalf := s[:half]
	secondHalf := s[half:]

	return firstHalf == secondHalf
}

func isInvalidPart2(num int) bool {
	s := strconv.Itoa(num)
	n := len(s)

	for i := 0; i < n/2; i++ {
		patternLen := i + 1

		if n%patternLen != 0 {
			continue
		}

		section1 := s[0:patternLen]
		section2 := s

		if checkForDivisor(section1, section2) {
			return true
		}
	}

	return false
}

func checkForDivisor(string1 string, string2 string) bool {
	return string1+string2 == string2+string1
}

func readData(filename string) []Range {
	content, err := os.ReadFile(filename)

	if err != nil {
		log.Fatal(err)
	}

	ranges := formatData(string(content))

	return ranges
}

func formatData(content string) []Range {
	idRanges := []Range{}

	content = strings.TrimSpace(content)
	rangeStrings := strings.Split(content, ",")

	for _, r := range rangeStrings {
		r = strings.TrimSpace(r)
		if r == "" {
			continue
		}
		values := strings.Split(r, "-")

		bottom, err := strconv.Atoi(strings.TrimSpace(values[0]))
		if err != nil {
			log.Fatal(err)
		}
		top, err := strconv.Atoi(strings.TrimSpace(values[1]))
		if err != nil {
			log.Fatal(err)
		}

		idRanges = append(idRanges, Range{
			bottom: bottom,
			top:    top,
		})
	}

	return idRanges
}

type Range struct {
	bottom int
	top    int
}
