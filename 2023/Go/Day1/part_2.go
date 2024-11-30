package main

import (
	"bufio"
	"os"
	"strings"
)

/*


The Idea for part two is that in the input lines we have
like english letter of digits like one two three etc

so that why we have to parse them  and scan them and set them to digit so we can work with them





*/

// here I created the map for those letter so we will iterate over them and set them to corresponding digits
var wordToDigit = map[string]int{
	"zero":  0,
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
}

func SumCalibrationValues() int {

	// Lets read the input file

	file, err := os.Open("input.txt")
	if err != nil {
		return -1
	}
	defer file.Close()
	var sum int = 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		// Now we got the each and every line
		var lines string = scanner.Text()
		var firstDigit int
		var lastDigit int
		// first set mean if we set the first letter to the digit
		var firstSet bool
		// Let's iterate over each adn every line

		for i := 0; i < len(lines); i++ {
			if lines[i] >= '0' && lines[i] <= '9' {
				var digit int = int(lines[i] - '0')
				// we have not set the first digit then
				if !firstSet {
					firstDigit = digit
					firstSet = true
				}
				lastDigit = digit
			} else {
				for word, dig := range wordToDigit {
					if checkWord(lines, i, word) {
						if !firstSet {
							firstDigit = dig
							firstSet = true
						}
						lastDigit = dig
						break
					}
				}
			}
		}
		sum += (firstDigit * 10) + lastDigit
	}

	// storing the final value
	return sum
}

// Now Let's run the code

// Let's write a function to check the word lines

func checkWord(line string, idx int, word string) bool {
	// basically we are scanning the hole like idx: mean till end of the line
	// and we are looking for the word
	return strings.HasPrefix(line[idx:], word)
}
