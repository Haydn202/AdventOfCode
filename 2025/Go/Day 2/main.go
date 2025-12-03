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
			if hasCommonDivisor(stringIndex) {
			// if hasARepeat(stringIndex) {
				invalidIds = append(invalidIds, i)
			}
		}
	} 

	return invalidIds
}

func hasARepeat(value string) bool {
	if (len(value) % 2 == 0) {
		num1 := value[0:(len(value) / 2)]
		num2 := value[(len(value) / 2):]

		if num1 == num2 {
			return true
		}
	}

	return false
}

func hasCommonDivisor(value string) bool {
	hasDivisor := false

	for i := 0; i < len(value) - 1; i++ {
		section1 := value[0: i + 1]
		section2 := value

		hasDivisor = checkForDivisor(section1, section2)

		if hasDivisor {
			break
		}
	}

	return hasDivisor
}

func checkForDivisor (string1 string, string2 string) bool {
	if string1 + string2 != string2 + string1 {
		return false
	}

	return true
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
