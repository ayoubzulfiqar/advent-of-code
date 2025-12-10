package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
)

type Puzzle struct {
	target  string
	joltage []int
	buttons [][]int
}

func parseInput(filename string) ([]Puzzle, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var puzzles []Puzzle
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}

		parts := strings.Fields(line)
		if len(parts) < 2 {
			continue
		}

		target := ""
		if len(parts[0]) >= 2 && parts[0][0] == '[' && parts[0][len(parts[0])-1] == ']' {
			target = parts[0][1 : len(parts[0])-1]
		} else {
			continue
		}

		joltageStr := parts[len(parts)-1]
		var joltage []int
		if len(joltageStr) >= 2 && joltageStr[0] == '{' && joltageStr[len(joltageStr)-1] == '}' {
			values := strings.Split(joltageStr[1:len(joltageStr)-1], ",")
			for _, v := range values {
				if num, err := strconv.Atoi(strings.TrimSpace(v)); err == nil {
					joltage = append(joltage, num)
				}
			}
		} else {
			continue
		}

		var buttons [][]int
		for i := 1; i < len(parts)-1; i++ {
			btn := parts[i]
			if len(btn) >= 2 && btn[0] == '(' && btn[len(btn)-1] == ')' {
				positions := strings.Split(btn[1:len(btn)-1], ",")
				var button []int
				for _, pos := range positions {
					if pos = strings.TrimSpace(pos); pos != "" {
						if num, err := strconv.Atoi(pos); err == nil {
							button = append(button, num)
						}
					}
				}
				buttons = append(buttons, button)
			}
		}

		puzzles = append(puzzles, Puzzle{
			target:  target,
			joltage: joltage,
			buttons: buttons,
		})
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return puzzles, nil
}

func gaussianElimination(matrix [][]int) ([]int, [][]int) {
	m := len(matrix)
	if m == 0 {
		return nil, nil
	}
	n := len(matrix[0]) - 1 

	pivotCols := []int{}
	currentRow := 0

	mat := make([][]int, m)
	for i := range matrix {
		mat[i] = make([]int, n+1)
		copy(mat[i], matrix[i])
	}

	for col := 0; col < n && currentRow < m; col++ {
		pivotRow := -1
		for row := currentRow; row < m; row++ {
			if mat[row][col] != 0 {
				pivotRow = row
				break
			}
		}

		if pivotRow == -1 {
			continue
		}

		mat[currentRow], mat[pivotRow] = mat[pivotRow], mat[currentRow]
		pivotCols = append(pivotCols, col)

		for row := currentRow + 1; row < m; row++ {
			if mat[row][col] != 0 {
				factor := mat[row][col]
				pivotVal := mat[currentRow][col]

				for j := col; j <= n; j++ {
					mat[row][j] = mat[row][j]*pivotVal - mat[currentRow][j]*factor
				}
			}
		}

		currentRow++
	}

	return pivotCols, mat
}

func solveSystem(buttons [][]int, joltages []int) []int {
	n := len(buttons)
	m := len(joltages)

	matrix := make([][]int, m)
	for i := range matrix {
		matrix[i] = make([]int, n+1)
		for j := 0; j < n; j++ {
			affects := false
			for _, pos := range buttons[j] {
				if pos == i {
					affects = true
					break
				}
			}
			if affects {
				matrix[i][j] = 1
			}
		}
		matrix[i][n] = joltages[i]
	}

	pivotCols, reducedMatrix := gaussianElimination(matrix)
	if reducedMatrix == nil {
		return nil
	}

	pivotSet := make(map[int]bool)
	for _, col := range pivotCols {
		pivotSet[col] = true
	}

	freeVars := []int{}
	for i := 0; i < n; i++ {
		if !pivotSet[i] {
			freeVars = append(freeVars, i)
		}
	}

	bestSolution := make([]int, n)
	bestSum := -1

	var trySolution func(freeValues []int)
	trySolution = func(freeValues []int) {
		solution := make([]int, n)
		for i, varIdx := range freeVars {
			if i < len(freeValues) {
				solution[varIdx] = freeValues[i]
			}
		}

		for i := len(pivotCols) - 1; i >= 0; i-- {
			row := i
			col := pivotCols[i]
			total := reducedMatrix[row][n] // Constant term

			for j := col + 1; j < n; j++ {
				total -= reducedMatrix[row][j] * solution[j]
			}

			if reducedMatrix[row][col] == 0 {
				return 
			}

			if total%reducedMatrix[row][col] != 0 {
				return
			}

			val := total / reducedMatrix[row][col]
			if val < 0 {
				return
			}

			solution[col] = val
		}

		for i := 0; i < m; i++ {
			total := 0
			for j := 0; j < n; j++ {
				if solution[j] > 0 {
					for _, pos := range buttons[j] {
						if pos == i {
							total += solution[j]
							break
						}
					}
				}
			}
			if total != joltages[i] {
				return
			}
		}

		totalPresses := 0
		for _, val := range solution {
			totalPresses += val
		}

		if bestSum == -1 || totalPresses < bestSum {
			copy(bestSolution, solution)
			bestSum = totalPresses
		}
	}

	if len(freeVars) == 0 {
		trySolution([]int{})
	} else if len(freeVars) == 1 {
		maxVal := 0
		for _, j := range joltages {
			if j > maxVal {
				maxVal = j
			}
		}
		maxVal *= 3
		for val := 0; val <= maxVal; val++ {
			if bestSum != -1 && val > bestSum {
				break
			}
			trySolution([]int{val})
		}
	} else if len(freeVars) == 2 {
		maxVal := 0
		for _, j := range joltages {
			if j > maxVal {
				maxVal = j
			}
		}
		if maxVal < 200 {
			maxVal = 200
		}
		for v1 := 0; v1 <= maxVal; v1++ {
			for v2 := 0; v2 <= maxVal; v2++ {
				if bestSum != -1 && v1+v2 > bestSum {
					continue
				}
				trySolution([]int{v1, v2})
			}
		}
	} else if len(freeVars) == 3 {
		for v1 := 0; v1 < 250; v1++ {
			for v2 := 0; v2 < 250; v2++ {
				for v3 := 0; v3 < 250; v3++ {
					if bestSum != -1 && v1+v2+v3 > bestSum {
						continue
					}
					trySolution([]int{v1, v2, v3})
				}
			}
		}
	} else if len(freeVars) == 4 {
		for v1 := 0; v1 < 30; v1++ {
			for v2 := 0; v2 < 30; v2++ {
				for v3 := 0; v3 < 30; v3++ {
					for v4 := 0; v4 < 30; v4++ {
						if bestSum != -1 && v1+v2+v3+v4 > bestSum {
							continue
						}
						trySolution([]int{v1, v2, v3, v4})
					}
				}
			}
		}
	} else {
		trySolution(make([]int, len(freeVars)))
	}

	if bestSum == -1 {
		return make([]int, n)
	}

	return bestSolution
}

func ConfigureJoltLevel(puzzles []Puzzle) int {
	total := 0
	var wg sync.WaitGroup
	results := make(chan int, len(puzzles))

	for _, puzzle := range puzzles {
		wg.Add(1)
		go func(p Puzzle) {
			defer wg.Done()
			solution := solveSystem(p.buttons, p.joltage)
			sum := 0
			for _, val := range solution {
				sum += val
			}
			results <- sum
		}(puzzle)
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	for res := range results {
		total += res
	}

	return total
}

func main() {
	puzzles, err := parseInput("input.txt")
	if err != nil {
		fmt.Printf("Error reading input: %v\n", err)
		return
	}

	result2 := ConfigureJoltLevel(puzzles)
	fmt.Println("Part-2:", result2)
}
