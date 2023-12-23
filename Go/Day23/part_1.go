package main

import (
	"bufio"
	"container/list"
	"fmt"
	"os"
)

func parse(taskInput []string) map[[2]int]rune {
	hikingMap := make(map[[2]int]rune)
	for rowIndex, line := range taskInput {
		for columnIndex, character := range line {
			hikingMap[[2]int{rowIndex, columnIndex}] = character
		}
	}
	return hikingMap
}

func getNeighbors(hikingMap map[[2]int]rune, point [2]int) <-chan [2]int {
	neighbors := make(chan [2]int)
	go func() {
		defer close(neighbors)

		if hikingMap[point] == 'v' {
			neighbors <- [2]int{point[0] + 1, point[1]}
			return
		}

		if hikingMap[point] == '^' {
			neighbors <- [2]int{point[0] - 1, point[1]}
			return
		}

		if hikingMap[point] == '>' {
			neighbors <- [2]int{point[0], point[1] + 1}
			return
		}

		if hikingMap[point] == 'v' {
			neighbors <- [2]int{point[0], point[1] - 1}
			return
		}

		directions := [][2]int{{0, -1}, {1, 0}, {-1, 0}, {0, 1}}
		for _, direction := range directions {
			newPoint := [2]int{point[0] + direction[0], point[1] + direction[1]}
			if _, ok := hikingMap[newPoint]; !ok || hikingMap[newPoint] == '#' {
				continue
			}
			neighbors <- newPoint
		}
	}()
	return neighbors
}

func findTheLongestPath(hikingMap map[[2]int]rune, start, end [2]int) int {
	toCheck := list.New()
	toCheck.PushBack([3]interface{}{start, make(map[[2]int]struct{}), 0})

	costSoFar := make(map[[2]int]int)
	costSoFar[start] = 0

	for toCheck.Len() > 0 {
		element := toCheck.Remove(toCheck.Back()).([3]interface{})
		rowIndex, columnIndex, path := element[0].([2]int)[0], element[0].([2]int)[1], element[1].(map[[2]int]struct{})

		if [2]int{rowIndex, columnIndex} == end {
			continue
		}

		for newPoint := range getNeighbors(hikingMap, [2]int{rowIndex, columnIndex}) {
			newCost := costSoFar[[2]int{rowIndex, columnIndex}] + 1

			if _, exists := path[newPoint]; exists {
				continue
			}

			if _, exists := costSoFar[newPoint]; !exists || newCost > costSoFar[newPoint] {
				costSoFar[newPoint] = newCost

				newPath := make(map[[2]int]struct{})
				for k := range path {
					newPath[k] = struct{}{}
				}
				newPath[newPoint] = struct{}{}

				toCheck.PushFront([3]interface{}{newPoint, newPath, newCost})
			}
		}
	}

	return costSoFar[end]
}

func solution(taskInput []string) int {
	hikingMap := parse(taskInput)
	maxRows := 0
	for key := range hikingMap {
		if key[0] > maxRows {
			maxRows = key[0]
		}
	}

	start := [2]int{}
	end := [2]int{}
	for point, tile := range hikingMap {
		if point[0] == 0 && tile == '.' {
			start = point
		}
		if point[0] == maxRows && tile == '.' {
			end = point
		}
	}

	return findTheLongestPath(hikingMap, start, end)
}

func LongISTheLongestHike() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	var taskInput []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		taskInput = append(taskInput, scanner.Text())
	}

	result := solution(taskInput)
	fmt.Println(result)
}
