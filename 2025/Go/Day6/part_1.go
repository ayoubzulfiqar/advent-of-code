package main

import (
	"fmt"
	"os"
	"strings"
)

type Data struct {
	operators []byte
	positions []int
	lines     []string
}

func IndividualGrandTotals() int {
	fileContent, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Printf("Error reading input.txt: %v\n", err)
		os.Exit(1)
	}
	content := string(fileContent)

	lines := strings.Split(content, "\n")

	if len(lines) > 0 && lines[len(lines)-1] == "" {
		lines = lines[:len(lines)-1]
	}

	if len(lines) == 0 {
		return 0
	}

	operationLine := lines[len(lines)-1]
	operators := make([]byte, 0)
	operatorColumns := make([]int, 0)

	for column := 0; column < len(operationLine); column++ {
		if operationLine[column] != ' ' {
			operators = append(operators, operationLine[column])
			operatorColumns = append(operatorColumns, column)
		}
	}
	operatorColumns = append(operatorColumns, len(operationLine))

	dataRowCount := len(lines) - 1
	if dataRowCount <= 0 {
		return 0
	}

	grandTotal := 0

	for operatorIndex := 0; operatorIndex < len(operators); operatorIndex++ {
		currentOperator := operators[operatorIndex]
		columnResult := 0

		if currentOperator == '*' {
			columnResult = 1
		} else {
			columnResult = 0
		}

		columnStart := operatorColumns[operatorIndex]
		columnEnd := operatorColumns[operatorIndex+1]

		if operatorIndex < len(operators)-1 {
			columnEnd -= 1
		}

		for row := 0; row < dataRowCount; row++ {
			parsedNumber := 0

			for col := columnStart; col < columnEnd; col++ {
				digitChar := lines[row][col]
				if digitChar != ' ' {
					parsedNumber = parsedNumber*10 + int(digitChar-'0')
				}
			}

			if currentOperator == '*' {
				columnResult *= parsedNumber
			} else {
				columnResult += parsedNumber
			}
		}

		grandTotal += columnResult
	}

	return grandTotal
}
