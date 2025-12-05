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
		maxTotal = maxTotal + battery.getMaxCharge(2)
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

		batteries = append(batteries, Battery{ cells: charges})
	}

	return batteries
}

type Battery struct {
	cells []int
	maxCharge int
}

func (b *Battery) getMaxCharge(cellCount int) int {
	// Find the largest number that not in the final index
	// find the largest number after that index
	firstDigit := 0
	secondDigit := 0

	for i := 0; i <= len(b.cells) - cellCount; i++ {
		lastIndex := i == len(b.cells)

		if !lastIndex {
			if b.cells[i] > firstDigit {
				firstDigit = b.cells[i]
				secondDigit = 0
			}

			if b.cells[i + 1] > secondDigit {
				secondDigit = b.cells[i + 1]
			}
		}
	}

	val, err := strconv.Atoi(strconv.Itoa(firstDigit) + strconv.Itoa(secondDigit))

	if err != nil {
		log.Fatal(err)
	}

	b.maxCharge = val

	return val
}