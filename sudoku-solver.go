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

const (
	sepSize = 8
)

type Grid [9][9]int

func main() {
	filePath := flag.String("path", "", "The file path.")
	flag.Parse()

	if *filePath == "" {
		log.Fatal("Please, provide the file path")
	}

	Grids, err := readGridsFromFile(*filePath)
	if err != nil {
		log.Fatal("Error on reading grids from file:", err)
	}

	log.Printf("Found %v grid(s) in %v\n", len(Grids), *filePath)

	for i, grid := range Grids {
		log.Printf("Solving grid %v...", i+1)
		SolvedGrid, solved := solve(grid)
		if solved {
			log.Printf("Solved:")
			printGrids(grid, SolvedGrid)
		} else {
			log.Printf("I couldn't solve the grid :(")
		}
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

func readGridsFromFile(filePath string) ([]Grid, error) {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	grids := []Grid{}
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		if len(line) == 0 {
			continue
		}

		var grid Grid
		for i := 0; i < 9; i++ {
			if i > 0 {
				scanner.Scan()
				if len(scanner.Text()) == 0 {
					break
				}
				line = scanner.Text()
			}

			for j := 0; j < 9; j++ {
				grid[i][j] = int(line[j]) - '0'
			}
		}
		grids = append(grids, grid)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return grids, nil
}

func printGrids(sourceGrid Grid, solvedGrid Grid) {
	var result string
	for i := 0; i < 9; i++ {
		if i > 0 && i%3 == 0 {
			result += "- - - + - - - + - - - " + strings.Repeat(" ", sepSize) + "- - - + - - - + - - -\n"
		}
		result += rowToString(sourceGrid[i], solvedGrid[i], (i+1)%5 == 0)
	}
	result += "\n"

	fmt.Print(result)
}

func rowToString(firstRow [9]int, secondRow [9]int, showArrow bool) string {
	var builder strings.Builder

	for i, value := range firstRow {
		if i > 0 && i%3 == 0 {
			builder.WriteString("| ")
		}
		builder.WriteString(fmt.Sprintf("%v ", value))
	}

	if showArrow {
		builder.WriteString(
			strings.Repeat(" ", (sepSize-4)/2) +
				"--->" +
				strings.Repeat(" ", (sepSize-4)/2))
	} else {
		builder.WriteString(strings.Repeat(" ", sepSize))
	}

	for i, value := range secondRow {
		if i > 0 && i%3 == 0 {
			builder.WriteString("| ")
		}
		builder.WriteString(fmt.Sprintf("%v ", value))
	}

	builder.WriteString("\n")

	return builder.String()
}
