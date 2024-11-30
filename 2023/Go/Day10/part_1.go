package main

import (
	"bufio"
	"fmt"
	"os"
)

func FurthestDistance() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var grid [][]byte

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := []byte(scanner.Text())
		grid = append(grid, line)
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	var startingRow, startingColumn int

OuterLoop:
	for row := range grid {
		for column := range grid[row] {
			if grid[row][column] == 'S' {
				startingRow = row
				startingColumn = column
				break OuterLoop
			}
		}
	}

	checkPipes := []struct{ row, column int }{{startingRow, startingColumn}}
	seenPipes := map[struct{ row, column int }]struct{}{
		{startingRow, startingColumn}: {},
	}

	for len(checkPipes) > 0 {
		row, column := checkPipes[0].row, checkPipes[0].column
		checkPipes = checkPipes[1:]

		currentPipe := grid[row][column]

		if row > 0 && contains("S|LJ", currentPipe) && contains("|7F", grid[row-1][column]) && struct{ row, column int }{row - 1, column} != (struct{ row, column int }{}) {
			neighbor := struct{ row, column int }{row - 1, column}
			if _, exists := seenPipes[neighbor]; !exists {
				seenPipes[neighbor] = struct{}{}
				checkPipes = append(checkPipes, neighbor)
			}
		}

		if row < len(grid)-1 && contains("S|7F", currentPipe) && contains("|LJ", grid[row+1][column]) && struct{ row, column int }{row + 1, column} != (struct{ row, column int }{}) {
			neighbor := struct{ row, column int }{row + 1, column}
			if _, exists := seenPipes[neighbor]; !exists {
				seenPipes[neighbor] = struct{}{}
				checkPipes = append(checkPipes, neighbor)
			}
		}

		if column > 0 && contains("S-7J", currentPipe) && contains("-LF", grid[row][column-1]) && struct{ row, column int }{row, column - 1} != (struct{ row, column int }{}) {
			neighbor := struct{ row, column int }{row, column - 1}
			if _, exists := seenPipes[neighbor]; !exists {
				seenPipes[neighbor] = struct{}{}
				checkPipes = append(checkPipes, neighbor)
			}
		}

		if column < len(grid[row])-1 && contains("S-LF", currentPipe) && contains("-J7", grid[row][column+1]) && struct{ row, column int }{row, column + 1} != (struct{ row, column int }{}) {
			neighbor := struct{ row, column int }{row, column + 1}
			if _, exists := seenPipes[neighbor]; !exists {
				seenPipes[neighbor] = struct{}{}
				checkPipes = append(checkPipes, neighbor)
			}
		}
	}

	furthestDistance := len(seenPipes) / 2
	fmt.Println(furthestDistance)
}

func contains(s string, c byte) bool {
	for i := 0; i < len(s); i++ {
		if s[i] == c {
			return true
		}
	}
	return false
}
