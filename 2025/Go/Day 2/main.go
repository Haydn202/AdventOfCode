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
	invalid := getInvalid(ranges)
	result := sumArray(invalid)
	
	fmt.Println(result)
}

func sumArray(numbers []int) int {
    result := 0
    for _, num := range numbers {
        result += num
    }
    return result
}

func getInvalid(ranges []Range) []int {
	invalidIds := []int{}

	for _, values := range ranges {
		for i := values.bottom; i <= values.top; i++  {
			stringIndex := strconv.Itoa(i)
			if (len(stringIndex) % 2 == 0) {
				num1 := stringIndex[0:(len(stringIndex) / 2)]
				num2 := stringIndex[(len(stringIndex) / 2):]

				if num1 == num2 {
					invalidIds = append(invalidIds, i)
				}
			}
		}
	} 

	return invalidIds
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

	rangeStrings := strings.Split(content, ",")

	for _, r := range rangeStrings {
		values := strings.Split(r, "-")

		bottom, err := strconv.Atoi(values[0])
		if err != nil {
			log.Fatal(err)
		}
		top, err := strconv.Atoi(values[1])
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
