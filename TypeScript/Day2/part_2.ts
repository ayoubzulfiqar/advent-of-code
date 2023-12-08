import * as fs from 'fs';

// Function to find the maximum of two numbers
function max(a: number, b: number): number {
  return a > b ? a : b;
}

// Function to calculate power based on color counts and buffer values
function cubeConundrumPower(): number {
  try {
    // Read the contents of the file named "input.txt".
    const fileContent: string = fs.readFileSync('input.txt', 'utf-8');

    let power: number = 0;
    let bufferCount: number = 0;

    // Split the content into lines.
    const lines: string[] = fileContent.split('\n');

    for (let i = 0; i < lines.length; i++) {
      let redCount: number = 0, greenCount: number = 0, blueCount: number = 0;
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
        // Check the color of the current character and update the corresponding color count.
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

        // Reset bufferCount for the next character.
        bufferCount = 0;
      }

      // Update the power based on the product of color counts.
      power += redCount * greenCount * blueCount;
    }

    console.log(power);
    return power;
  } catch (error) {
    console.log('Error opening file:', error);
    return 0;
  }
}

// Call the function to execute the TypeScript code.
cubeConundrumPower();
