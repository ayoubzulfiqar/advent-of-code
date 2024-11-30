import 'dart:io';

int cubeConundrum() {
  try {
    // Open the file named "input.txt".
    File file = File("test.txt");
    // Check if the file exists.
    if (!file.existsSync()) {
      print("Error: File 'test.txt' not found.");
      return 0;
    }
    List<String> lines = file.readAsLinesSync();

    int total = 0;
    int bufferCount = 0;
    bool shouldExit = false;

    for (int i = 0; i < lines.length && !shouldExit; i++) {
      String line = lines[i];
      for (int j = 0; j < line.length; j++) {
        // If the character is a digit, update bufferCount accordingly.
        if (line[j].codeUnitAt(0) >= '0'.codeUnitAt(0) &&
            line[j].codeUnitAt(0) <= '9'.codeUnitAt(0)) {
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
            if (bufferCount > 12) {
              shouldExit = true;
            }
            break;
          case 'g':
            if (bufferCount > 13) {
              shouldExit = true;
            }
            break;
          case 'b':
            if (bufferCount > 14) {
              shouldExit = true;
            }
            break;
        }

        // Reset bufferCount for the next color.
        bufferCount = 0;
      }
      // If the loop completes without exceeding any limits, update the total.
      if (!shouldExit) {
        total += i + 1;
      }
    }

    print(total);
    return total;
  } catch (error) {
    print('Error opening file: $error');
    return 0;
  }
}
