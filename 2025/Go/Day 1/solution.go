package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	PositionCount := 0
	dial := Dial{
		CurrentPosition: 50,
		MaxPosition: 99,
		MinPosition: 0,
	}

	instructions := readFile("data.txt")

	for _, instruction := range instructions {
		count := dial.TurnWithZeroCount(instruction.Direction, instruction.Steps)
		PositionCount += count
	}

	fmt.Println(PositionCount)
}

func readFile(filename string) []Instruction {
	content, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	instructions := []Instruction{}

	for _, line := range strings.Split(string(content), "\n") {
		direction := line[0:1]
		steps, err := strconv.Atoi(line[1:])
		if err != nil {
			log.Fatal(err)
		}

		instructions = append(instructions, Instruction{
			Direction: direction,
			Steps: steps,
		})
	}

	return instructions
}

type Instruction struct {
	Direction string
	Steps int
}

type Dial struct {
	CurrentPosition int
	MaxPosition int
	MinPosition int
}

func (d *Dial) Turn(direction string, steps int) int {

	for i := 0; i < steps; i++ {
		if direction == "L" {
			d.CurrentPosition -= 1
		} else if direction == "R" {
			d.CurrentPosition += 1
		}

		if d.CurrentPosition > d.MaxPosition {
			d.CurrentPosition = d.MinPosition
		} else if d.CurrentPosition < d.MinPosition {
			d.CurrentPosition = d.MaxPosition
		}
	}

	return d.CurrentPosition
}

func (d *Dial) TurnWithZeroCount(direction string, steps int) int {
	zeroCount := 0

	for i := 0; i < steps; i++ {
		newPosition := d.Turn(direction, 1)
		if newPosition == 0 {
			zeroCount++
		}
	}

	return zeroCount
}