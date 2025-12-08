package main

import (
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	problems := readDataVertical("data.txt")
	total := 0

	for _, problem := range problems {
		solution := problem.SolveProblem()
		total += solution
	}

	println(total)
}

func readData(filename string) []MathsProblem {
	content, err := os.ReadFile(filename)

	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(content), "\n")

	nonEmptyLines := []string{}
	for _, line := range lines {
		line = strings.TrimRight(line, "\r")
		if strings.TrimSpace(line) != "" {
			nonEmptyLines = append(nonEmptyLines, line)
		}
	}

	if len(nonEmptyLines) == 0 {
		return []MathsProblem{}
	}

	operationLine := nonEmptyLines[len(nonEmptyLines)-1]
	numberLines := nonEmptyLines[:len(nonEmptyLines)-1]

	maxLen := 0
	for _, line := range nonEmptyLines {
		if len(line) > maxLen {
			maxLen = len(line)
		}
	}

	type numWithPos struct {
		num        int
		spaceCount int
		colIndex   int
	}

	allNumbersByLine := [][]numWithPos{}
	for _, line := range numberLines {
		lineNums := []numWithPos{}

		leadingSpaces := 0
		for leadingSpaces < len(line) && line[leadingSpaces] == ' ' {
			leadingSpaces++
		}

		fields := strings.Fields(line)
		currentPos := 0
		fieldIndex := 0

		for _, field := range fields {
			pos := strings.Index(line[currentPos:], field)
			if pos == -1 {
				continue
			}
			startPos := currentPos + pos

			spacesBefore := 0
			for i := 0; i < startPos; i++ {
				if line[i] == ' ' {
					spacesBefore++
				}
			}

			number, err := strconv.Atoi(field)
			if err == nil {
				lineNums = append(lineNums, numWithPos{
					num:        number,
					spaceCount: spacesBefore,
					colIndex:   fieldIndex,
				})
			}

			currentPos = startPos + len(field)
			fieldIndex++
		}
		allNumbersByLine = append(allNumbersByLine, lineNums)
	}

	numColumns := 0
	for _, lineNums := range allNumbersByLine {
		for _, n := range lineNums {
			if n.colIndex >= numColumns {
				numColumns = n.colIndex + 1
			}
		}
	}

	opFields := strings.Fields(operationLine)
	operations := make([]Operation, numColumns)
	for i := 0; i < numColumns; i++ {
		if i < len(opFields) {
			if opFields[i] == "+" {
				operations[i] = Addition
			} else if opFields[i] == "*" {
				operations[i] = Multiplication
			} else {
				operations[i] = Addition
			}
		} else {
			operations[i] = Addition
		}
	}

	problems := []MathsProblem{}
	for col := 0; col < numColumns; col++ {
		nums := []Num{}

		for _, lineNums := range allNumbersByLine {
			for _, n := range lineNums {
				if n.colIndex == col {
					nums = append(nums, Num{
						SpaceCount: n.spaceCount,
						Number:     n.num,
					})
				}
			}
		}

		if len(nums) > 0 {
			problems = append(problems, MathsProblem{
				Operation: operations[col],
				Nums:      nums,
			})
		}
	}

	return problems
}

func readDataVertical(filename string) []MathsProblem {
	content, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(content), "\n")
	chars := [][]byte{}
	for _, line := range lines {
		line = strings.TrimRight(line, "\r")
		if strings.TrimSpace(line) != "" {
			chars = append(chars, []byte(line))
		}
	}

	if len(chars) == 0 {
		return []MathsProblem{}
	}

	NN := len(chars)
	MM := len(chars[NN-1])

	for i := 0; i < NN-1; i++ {
		for len(chars[i]) < MM {
			chars[i] = append(chars[i], ' ')
		}
	}

	problems := []MathsProblem{}
	j := 0

	for j < MM {
		left := j

		for (j+1 < MM && chars[NN-1][j+1] == ' ') || j+1 == MM {
			j++
		}

		right := j

		nums := []Num{}
		for col := left; col < right; col++ {
			numStr := ""
			for row := 0; row < NN-1; row++ {
				if col < len(chars[row]) && chars[row][col] != ' ' {
					numStr += string(chars[row][col])
				}
			}
			if numStr != "" {
				number, _ := strconv.Atoi(numStr)
				spaces := 0
				for i := 0; i < left && i < len(chars[0]) && chars[0][i] == ' '; i++ {
					spaces++
				}
				nums = append(nums, Num{SpaceCount: spaces, Number: number})
			}
		}

		if len(nums) > 0 && left < len(chars[NN-1]) {
			opChar := chars[NN-1][left]
			var op Operation
			if opChar == '+' {
				op = Addition
			} else if opChar == '*' {
				op = Multiplication
			} else {
				op = Addition
			}
			problems = append(problems, MathsProblem{Operation: op, Nums: nums})
		}

		j = right + 1
	}

	for i, j := 0, len(problems)-1; i < j; i, j = i+1, j-1 {
		problems[i], problems[j] = problems[j], problems[i]
	}

	return problems
}

type Num struct {
	SpaceCount int
	Number     int
}

type MathsProblem struct {
	Operation Operation
	Nums      []Num
	Solution  int
}

func (M *MathsProblem) SolveProblem() int {
	solution := 0

	if len(M.Nums) == 0 {
		M.Solution = 0
		return 0
	}

	if M.Operation == Addition {
		for _, num := range M.Nums {
			solution += num.Number
		}
	} else if M.Operation == Multiplication {
		solution = 1
		for _, num := range M.Nums {
			solution *= num.Number
		}
	}

	M.Solution = solution
	return solution
}

type Operation int

const (
	Addition Operation = iota
	Multiplication
)

func (o Operation) String() string {
	switch o {
	case Addition:
		return "+"
	case Multiplication:
		return "*"
	default:
		return "unknown"
	}
}
