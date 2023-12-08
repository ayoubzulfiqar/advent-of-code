import * as fs from 'fs';

function max(a: number, b: number): number {
  return a > b ? a : b;
}

function cubeConundrum(): number {
  try {
    // Read the contents of the file named "input.txt".
    const fileContent: string = fs.readFileSync('input.txt', 'utf-8');

    let total: number = 0;
    let bufferCount: number = 0;

    // Split the content into lines.
    const lines: string[] = fileContent.split('\n');

    for (let i = 0; i < lines.length; i++) {
      const line: string = lines[i];
      for (let j = 0; j < line.length; j++) {
        // If the character is a digit, update bufferCount accordingly.
        if (line[j] >= '0' && line[j] <= '9') {
          bufferCount *= 10;
          bufferCount += parseInt(line[j], 10);
          continue;
        }
        // If the character is a space, skip to the next character.
        if (line[j] === ' ') {
          continue;
        }
        // Check the color of the current character and compare bufferCount.
        switch (line[j]) {
          case 'r':
            if (bufferCount > 12) {
              // If bufferCount exceeds the limit for red, exit the loop.
              bufferCount = 0;
              break;
            }
            break;
          case 'g':
            if (bufferCount > 13) {
              // If bufferCount exceeds the limit for green, exit the loop.
              bufferCount = 0;
              break;
            }
            break;
          case 'b':
            if (bufferCount > 14) {
              // If bufferCount exceeds the limit for blue, exit the loop.
              bufferCount = 0;
              break;
            }
            break;
        }

        // Reset bufferCount for the next character.
        bufferCount = 0;
      }
      // If the loop completes without exceeding any limits, update the total.
      total += i + 1;
    }

    console.log(total);
    return total;
  } catch (error) {
    console.log('Error opening file:', error);
    return 0;
  }
}

// Call the function to execute the TypeScript code.
cubeConundrum();
