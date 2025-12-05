package main

import (
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	batteries := readData("data.txt")

	maxTotal := 0

	for _, battery := range batteries {
		maxTotal = maxTotal + battery.getMaxCharge(12)
	}

	println(maxTotal)
}

func readData(filename string) []Battery {
	content, err := os.ReadFile(filename)

	if err != nil {
		log.Fatal(err)
	}

	batteries := []Battery{}

	for _, line := range strings.Split(string(content), "\n") {
		charges := []int{}
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		for _, c := range line {
			charge, err := strconv.Atoi(string(c))

			if err != nil {
				log.Fatal(err)
			}

			charges = append(charges, charge)
		}

		batteries = append(batteries, Battery{cells: charges})
	}

	return batteries
}

type Battery struct {
	cells     []int
	maxCharge int
}

func (b *Battery) getMaxCharge(cellCount int) int {
	result := 0
	startIdx := 0

	for pos := 0; pos < cellCount; pos++ {
		remaining := cellCount - pos - 1
		maxEndIdx := len(b.cells) - remaining

		maxDigit := -1
		maxDigitIdx := -1
		for i := startIdx; i < maxEndIdx; i++ {
			if b.cells[i] > maxDigit {
				maxDigit = b.cells[i]
				maxDigitIdx = i
			}
		}

		result = result*10 + maxDigit
		startIdx = maxDigitIdx + 1
	}

	b.maxCharge = result
	return result
}
