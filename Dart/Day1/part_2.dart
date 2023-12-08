import 'dart:io';

Map<String, int> wordToDigit = {
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

int sumCalibrationValues() {
  try {
    // Open the file named "input.txt" for reading.
    var file = File('input.txt');
    var lines = file.readAsLinesSync();

    // Variables to store the total sum of digit values.
    var total = 0;

    // Iterate over each line in the file.
    for (var line in lines) {
      // Variables to store the first and last digit values in the line.
      var firstDigit = 0;
      var lastDigit = 0;
      var firstSet = false;

      // Iterate over each character in the line.
      for (var i = 0; i < line.length; i++) {
        // If the character is a digit, update firstDigit and lastDigit.
        if (line[i].contains(RegExp(r'[0-9]'))) {
          var dig = int.parse(line[i]);
          if (!firstSet) {
            firstDigit = dig;
            firstSet = true;
          }
          lastDigit = dig;
        } else {
          // If the character is not a digit, check if it forms a word that corresponds to a digit.
          for (var entry in wordToDigit.entries) {
            var word = entry.key;
            var digit = entry.value;
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
    print('$total');

    // Return the total sum.
    return total;
  } catch (e) {
    // If there is an error reading the file, return 0.
    return 0;
  }
}

// checkWord checks if a word starting from a specific index in a line matches a given word.
bool checkWord(String line, int i, String word) {
  return line.startsWith(word, i);
}
