package main

import (
	"bufio"
	"os"
	"strconv"
	"unicode"
)

func SumCalibration() (int, error) {
	// First read the input file

	file, err := os.Open("input.txt")
	if err != nil {
		return 0, err
	}
	defer file.Close()

	// if the file open successfully then we will scan the things inside the file
	scanner := bufio.NewScanner(file)
	var sumCalibration int = 0
	for scanner.Scan() {
		var lines string = scanner.Text()
		// Now lines hold the each and every line so let's select the line from each and every one
		firstDigit, foundFirst := findFirstDigit(lines)
		lastDigit, foundLast := findLastDigit(lines)

		// Means Both we find and return true then
		if foundFirst && foundLast {
			calibrationValue := (firstDigit * 10) + lastDigit
			sumCalibration += calibrationValue
		}

	}

	// Now We Return the final Sum
	// Now the Answer should be 142 as shown in example
	return sumCalibration, nil
}

/*


Now Lets break the code like first lets find the first element digit
*/

func findFirstDigit(s string) (int, bool) {
	// we will iterate though the string to scan for any available digits
	for i := 0; i < len(s); i++ {
		var char rune = rune(s[i])
		if unicode.IsDigit(char) {
			digit, err := strconv.Atoi(string(char))
			if err != nil {
				return 0, false
			}
			return digit, true

		}
	}
	return 0, false
}

// Lets Find the last digit inside the String

func findLastDigit(s string) (int, bool) {

	// Running the for loop from back of the string line
	// so we will select the last element easily
	for i := len(s) - 1; i >= 0; i-- {
		var char rune = rune(s[i])
		if unicode.IsDigit(char) {
			digit, err := strconv.Atoi(string(char))
			if err != nil {
				return 0, false
			}
			return digit, true
		}
	}
	return 0, false
}
