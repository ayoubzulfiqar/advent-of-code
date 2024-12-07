package main

import (
	"os"
	"strconv"
	"strings"
)

// ParseInputFile reads and parses the input file into a slice of result-term pairs.
func ParseInputFile() ([][2]interface{}, error) {
	content, err := os.ReadFile("./input.txt")
	if err != nil {
		return nil, err
	}

	lines := strings.Split(strings.TrimSpace(string(content)), "\n")
	var calibrationEquations [][2]interface{}

	for _, line := range lines {
		parts := strings.Split(line, ":")
		result, _ := strconv.Atoi(strings.TrimSpace(parts[0]))
		termsStr := strings.Fields(strings.TrimSpace(parts[1]))
		var terms []int
		for _, term := range termsStr {
			termInt, _ := strconv.Atoi(term)
			terms = append(terms, termInt)
		}
		calibrationEquations = append(calibrationEquations, [2]interface{}{result, terms})
	}
	return calibrationEquations, nil
}

// Add operator: sums the accumulator and term.
func Add(acc *int, term int) int {
	if acc == nil {
		return term
	}
	return *acc + term
}

// Multiply operator: multiplies the accumulator and term.
func Multiply(acc *int, term int) int {
	if acc == nil {
		return term
	}
	return *acc * term
}

// Concatenate operator: appends the current term to the accumulator.
func Concatenate(acc *int, term int) int {
	if acc == nil {
		return term
	}
	accStr := strconv.Itoa(*acc)
	termStr := strconv.Itoa(term)
	concatenated, _ := strconv.Atoi(accStr + termStr)
	return concatenated
}

// ValidateEquation recursively validates if the result can be achieved using the given terms and operators.
func ValidateEquation(result int, terms []int, acc *int, operators []func(*int, int) int) bool {
	// Base case
	if len(terms) == 0 {
		return acc != nil && *acc == result
	}

	// Recursive case
	for _, op := range operators {
		newAcc := op(acc, terms[0])
		if ValidateEquation(result, terms[1:], &newAcc, operators) {
			return true
		}
	}
	return false
}

// Part1 solves part 1 using Add and Multiply operators.
func totalEquationCalibration() (int, error) {
	calibrationEquations, err := ParseInputFile()
	if err != nil {
		return 0, err
	}

	total := 0
	for _, eq := range calibrationEquations {
		result := eq[0].(int)
		terms := eq[1].([]int)
		if ValidateEquation(result, terms, nil, []func(*int, int) int{Add, Multiply}) {
			total += result
		}
	}
	return total, nil
}

