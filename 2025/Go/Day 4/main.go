package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	grid := readData("data.txt")

	count := grid.getCountOfItemsWithLessThanXNeighbors(4)
	println(count)
}

type Grid struct {
	items [][]bool
}

func (g *Grid) getCountOfItemsWithLessThanXNeighbors(count int) int {
	hasEnough := 0
	for i := 0; i < len(g.items); i++ {
		for j := 0; j < len(g.items[i]); j++ {
			if g.items[i][j] {
				if g.getNeighborCount(i, j) < count {
					hasEnough++
				}
			}
		}
	}

	return hasEnough
}

func (g *Grid) getNeighborCount(x int, y int) int {
	// count the item itself
	count := -1

	for i := x - 1; i <= x+1; i++ {
		for j := y - 1; j <= y+1; j++ {
			if i >= 0 && j >= 0 && i < len(g.items) && j < len(g.items[i]) {
				if g.items[i][j] {
					count++
				}
			}
		}
	}

	return count
}

func (g *Grid) Print() {
	for _, row := range g.items {
		for _, cell := range row {
			if cell {
				fmt.Print("█")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}

func readData(filename string) Grid {
	grid := Grid{
		items: [][]bool{},
	}

	content, err := os.ReadFile(filename)

	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(content), "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		row := make([]bool, len(line))
		for i, val := range line {
			if val == '@' {
				row[i] = true
			}
		}
		grid.items = append(grid.items, row)
	}

	return grid
}
