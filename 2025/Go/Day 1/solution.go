package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	dial := Dial{
		CurrentPosition: 50,
		MaxPosition:     99,
		MinPosition:     0,
	}

	instructions := readFile("data.txt")

	for _, instruction := range instructions {
		dial.Turn(instruction.Direction, instruction.Steps)
	}

	fmt.Println(dial.FinishedOnZeroCount)
	fmt.Println(dial.PassedZeroCount)
}

func readFile(filename string) []Instruction {
	content, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	instructions := []Instruction{}

	for _, line := range strings.Split(string(content), "\n") {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		direction := line[0:1]
		steps, err := strconv.Atoi(line[1:])
		if err != nil {
			log.Fatal(err)
		}

		instructions = append(instructions, Instruction{
			Direction: direction,
			Steps:     steps,
		})
	}

	return instructions
}

type Instruction struct {
	Direction string
	Steps     int
}

type Dial struct {
	CurrentPosition     int
	MaxPosition         int
	MinPosition         int
	FinishedOnZeroCount int
	PassedZeroCount     int
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

		if d.CurrentPosition == 0 {
			d.PassedZeroCount++
		}
	}

	if d.CurrentPosition == 0 {
		d.FinishedOnZeroCount++
	}

	return d.CurrentPosition
}
