package main

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	ranges, ingredients := readFile("data.txt")

	freshCount := 0

	for _, id := range ingredients.IngredientIds {
		if isInAnyRange(id, ranges) {
			freshCount++
		}
	}

	fmt.Println(getFreshIdCount(ConsolidateRanges(sortRanges(ranges))))
}

func isInAnyRange(id int, ranges []Range) bool {
	for _, r := range ranges {
		if id >= r.Start && id <= r.End {
			return true
		}
	}
	return false
}

func sortRanges(ranges []Range) []Range {
	sort.Slice(ranges, func(i, j int) bool {
		return ranges[i].Start < ranges[j].Start
	})
	return ranges
}

func getFreshIdCount(ranges []Range) int {
	freshCount := 0
	for _, r := range ranges {
		count := r.End - r.Start + 1
		freshCount += count
	}
	return freshCount
}

func ConsolidateRanges(ranges []Range) []Range {
	if len(ranges) == 0 {
		return ranges
	}

	result := []Range{ranges[0]}

	for i := 1; i < len(ranges); i++ {
		last := &result[len(result)-1]
		current := ranges[i]

		if current.Start <= last.End {
			if current.End > last.End {
				last.End = current.End
			}
		} else {
			result = append(result, current)
		}
	}

	return result
}

func readFile(filename string) ([]Range, IngredientList) {
	content, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(content), "\n")

	var ranges []Range
	var ingredients IngredientList
	parsingRanges := true

	for _, line := range lines {
		line = strings.TrimSpace(line)

		if line == "" {
			parsingRanges = false
			continue
		}

		if parsingRanges {
			parts := strings.Split(line, "-")
			if len(parts) == 2 {
				start, err1 := strconv.Atoi(parts[0])
				end, err2 := strconv.Atoi(parts[1])
				if err1 == nil && err2 == nil {
					ranges = append(ranges, Range{Start: start, End: end})
				}
			}
		} else {
			id, err := strconv.Atoi(line)
			if err == nil {
				ingredients.IngredientIds = append(ingredients.IngredientIds, id)
			}
		}
	}

	return ranges, ingredients
}

type FreshIdMap struct {
	IdMap map[int]bool
}

func (f *FreshIdMap) AddId(id int) {
	f.IdMap[id] = true
}

func (f *FreshIdMap) RemoveId(id int) {
	f.IdMap[id] = false
}

type Range struct {
	Start int
	End   int
}

type IngredientList struct {
	IngredientIds []int
}
