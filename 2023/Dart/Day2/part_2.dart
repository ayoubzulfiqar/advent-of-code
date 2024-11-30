import 'dart:io';

int max(int a, int b) {
  return a > b ? a : b;
}

int cubeConundrumPower() {
  try {
    // Open the file named "input.txt".
    File file = File("text.txt");
    List<String> lines = file.readAsLinesSync();

    int power = 0;
    int bufferCount = 0;

    for (String line in lines) {
      int redCount = 0, greenCount = 0, blueCount = 0;

      for (int j = 0; j < line.length; j++) {
        // If the character is a digit, update bufferCount accordingly.
        if (line.codeUnitAt(j) >= '0'.codeUnitAt(0) &&
            line.codeUnitAt(j) <= '9'.codeUnitAt(0)) {
          bufferCount *= 10;
          bufferCount += int.parse(line[j]);
          continue;
        }
        // If the character is a space, skip to the next character.
        if (line[j] == ' ') {
          continue;
        }
        // Check the color of the current character and compare bufferCount.
        switch (line[j]) {
          case 'r':
            redCount = max(redCount, bufferCount);
            break;
          case 'g':
            greenCount = max(greenCount, bufferCount);
            break;
          case 'b':
            blueCount = max(blueCount, bufferCount);
            break;
        }

        // Reset bufferCount for the next color.
        bufferCount = 0;
      }

      // Update the power based on the product of color counts.
      power += redCount * greenCount * blueCount;
    }

    print(power);
    return power;
  } catch (error) {
    print('Error opening file: $error');
    return 0;
  }
}
