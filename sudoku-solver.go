package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"slices"
	"strings"
)

type Grid [9][9]int

func main() {
	filePath := flag.String("path", "", "The file path.")
	flag.Parse()

	if *filePath == "" {
		log.Fatal("Please, provide the file path")
	}

	Grid := readGridFromFile(*filePath)
	fmt.Println("Input grid:")
	printGrid(Grid)

	SolvedGrid, solved := solve(Grid)
	if solved {
		fmt.Println("Solved grid:")
		printGrid(SolvedGrid)
	} else {
		log.Fatal("I couldn't solve the grid :(")
	}
}

func solve(grid Grid) (Grid, bool) {
	row, col := firstEmptyCell(grid)
	if row == -1 && col == -1 { // Grid completely filled
		return grid, true
	}

	for i := 1; i <= 9; i++ {
		if isValid(grid, row, col, i) {
			grid[row][col] = i
			solvedGrid, solved := solve(grid)
			if solved {
				return solvedGrid, true
			}
		}
	}
	return grid, false
}

func isValid(grid Grid, row int, col int, i int) bool {
	grid[row][col] = i
	return validRow(grid, row) && validColumn(grid, col) && validSquare(grid, row, col)
}

func validRow(grid Grid, row int) bool {
	seenValues := []int{}
	for i := 0; i < 9; i++ {
		if grid[row][i] != 0 {
			if slices.Contains(seenValues, grid[row][i]) {
				return false
			}
			seenValues = append(seenValues, grid[row][i])
		}
	}
	return true
}

func validColumn(grid Grid, col int) bool {
	seenValues := []int{}
	for i := 0; i < 9; i++ {
		if grid[i][col] != 0 {
			if slices.Contains(seenValues, grid[i][col]) {
				return false
			}
			seenValues = append(seenValues, grid[i][col])
		}
	}
	return true
}

func validSquare(grid Grid, row int, col int) bool {
	seenValues := []int{}
	rowSpan := getSpan(row)
	columnSpan := getSpan(col)
	for i := rowSpan[0]; i <= rowSpan[2]; i++ {
		for j := columnSpan[0]; j <= columnSpan[2]; j++ {
			if grid[i][j] != 0 {
				if slices.Contains(seenValues, grid[i][j]) {
					return false
				}
				seenValues = append(seenValues, grid[i][j])
			}
		}
	}

	return true
}

func getSpan(index int) []int {
	if index < 3 {
		return []int{0, 1, 2}
	} else if index < 6 {
		return []int{3, 4, 5}
	}
	return []int{6, 7, 8}
}

func firstEmptyCell(grid Grid) (int, int) {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if grid[i][j] == 0 {
				return i, j
			}
		}
	}
	return -1, -1
}

func readGridFromFile(filePath string) Grid {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var grid Grid

	scanner := bufio.NewScanner(file)

	for i := 0; i < 9; i++ {
		scanner.Scan()
		line := scanner.Text()
		for j := 0; j < 9; j++ {
			grid[i][j] = int(line[j]) - '0'
		}
	}

	return grid
}

func printGrid(grid Grid) {
	var result string
	for i := 0; i < 9; i++ {
		if i > 0 && i%3 == 0 {
			result += "- - - + - - - + - - -\n"
		}
		result += rowToString(grid[i])
	}
	result += "\n"

	fmt.Print(result)
}

func rowToString(row [9]int) string {
	var builder strings.Builder

	for i, value := range row {
		if i > 0 && i%3 == 0 {
			builder.WriteString("| ")
		}
		builder.WriteString(fmt.Sprintf("%v ", value))
	}
	builder.WriteString("\n")

	return builder.String()
}
