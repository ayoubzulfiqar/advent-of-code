import * as fs from 'fs';

const wordToDigit: { [key: string]: number } = {
  "zero": 0,
  "one": 1,
  "two": 2,
  "three": 3,
  "four": 4,
  "five": 5,
  "six": 6,
  "seven": 7,
  "eight": 8,
  "nine": 9,
};

function sumCalibrationValues(): number {
  try {
    // Open the file named "input.txt" for reading.
    const fileContent = fs.readFileSync('input.txt', 'utf8');
    const lines = fileContent.split('\n');

    // Variables to store the total sum of digit values.
    let total = 0;

    // Iterate over each line in the file.
    for (const line of lines) {
      // Variables to store the first and last digit values in the line.
      let firstDigit = 0;
      let lastDigit = 0;
      let firstSet = false;

      // Iterate over each character in the line.
      for (let i = 0; i < line.length; i++) {
        // If the character is a digit, update firstDigit and lastDigit.
        if (/[0-9]/.test(line[i])) {
          const dig = parseInt(line[i], 10);
          if (!firstSet) {
            firstDigit = dig;
            firstSet = true;
          }
          lastDigit = dig;
        } else {
          // If the character is not a digit, check if it forms a word that corresponds to a digit.
          for (const [word, digit] of Object.entries(wordToDigit)) {
            if (checkWord(line, i, word)) {
              if (!firstSet) {
                firstDigit = digit;
                firstSet = true;
              }
              lastDigit = digit;
              break;
            }
          }
        }
      }

      // Calculate the total sum based on the first and last digits of the line.
      total += (firstDigit * 10) + lastDigit;
    }

    // Print the total sum.
    console.log(`${total}`);

    // Return the total sum.
    return total;
  } catch (e) {
    // If there is an error reading the file, return 0.
    return 0;
  }
}

// checkWord checks if a word starting from a specific index in a line matches a given word.
function checkWord(line: string, i: number, word: string): boolean {
  return line.startsWith(word, i);
}

// Call the sumCalibrationValues function to perform the calculation.
sumCalibrationValues();