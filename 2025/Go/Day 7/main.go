package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	grid := readData("data.txt")

	grid.ProcessBeam()
	grid.Print()

	println(grid.TimesSplit)
	println(sumArray(grid.paths))
}

func sumArray(numbers []int) int {
	result := 0
	for _, num := range numbers {
		result += num
	}
	return result
}

func readData(filename string) Grid {
	grid := Grid{
		Cells:      [][]Cell{},
		TimesSplit: 0,
		paths:      []int{},
	}

	content, err := os.ReadFile(filename)

	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(content), "\n")
	grid.paths = make([]int, len(lines[0]))

	for _, line := range lines {
		line = strings.TrimSpace(line)

		row := make([]Cell, len(line))

		for i, val := range line {
			if val == 'S' {
				row[i] = Cell{HasEmitter: true}
				grid.paths[i] = 1
			}
			if val == '^' {
				row[i] = Cell{HasSplitter: true}
			}
			if val == '.' {
				row[i] = Cell{}
			}
		}

		grid.Cells = append(grid.Cells, row)
	}

	return grid
}

type Grid struct {
	Cells      [][]Cell
	TimesSplit int
	paths      []int
}

type Cell struct {
	HasEmitter  bool
	HasBeam     bool
	HasSplitter bool
	paths       int
}

func (g *Grid) ProcessBeam() {
	for i := 0; i < len(g.Cells)-1; i++ {
		row := g.Cells[i]
		newPaths := make([]int, len(row))

		for j := 0; j < len(row); j++ {
			cell := row[j]
			nextRow := g.Cells[i+1]

			if cell.HasEmitter {
				nextRow[j].HasBeam = true
				newPaths[j] += g.paths[j]
			}

			if cell.HasBeam {
				if cell.HasSplitter {
					row[j-1].HasBeam = true
					row[j+1].HasBeam = true

					nextRow[j-1].HasBeam = true
					nextRow[j+1].HasBeam = true

					cell.HasBeam = false
					g.TimesSplit++

					newPaths[j-1] += g.paths[j]
					newPaths[j+1] += g.paths[j]
				}

				if !cell.HasSplitter {
					nextRow[j].HasBeam = true
					newPaths[j] += g.paths[j]
				}
			}
		}

		g.paths = newPaths
	}
}

func (g *Grid) Print() {
	for _, row := range g.Cells {
		for _, cell := range row {
			if cell.HasEmitter {
				fmt.Print("S")
			} else if cell.HasSplitter {
				fmt.Print("^")
			} else if cell.HasBeam {
				fmt.Print("|")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}
