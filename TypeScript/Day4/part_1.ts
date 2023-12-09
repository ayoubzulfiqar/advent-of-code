import * as fs from 'fs';

function scratchcardsWorth(): number {
  try {
    const fileContent = fs.readFileSync('./input.txt', 'utf-8');
    const lines = fileContent.split('\n');
    let t = 0;

    for (const line of lines) {
      const parts = line.split(":");
      if (parts.length < 2) {
        continue;
      }

      let x = parts[1].trim();
      const arrays = x.split(" | ");
      if (arrays.length !== 2) {
        continue;
      }

      const a = parseIntArray(arrays[0]);
      const b = parseIntArray(arrays[1]);

      let j = 0;
      for (const q of b) {
        for (const value of a) {
          if (q === value) {
            j++;
            break;
          }
        }
      }

      if (j > 0) {
        t += 1 << (j - 1);
      }
    }
    console.log(t);
    return t;
  } catch (e) {
    console.log(`Error: ${e}`);
    return 0;
  }
}

function parseIntArray(s: string): number[] {
  const result: number[] = [];
  const parts = s.split(' ');
  for (const part of parts) {
    const num = parseInt(part, 10);
    if (!isNaN(num)) {
      result.push(num);
    }
  }
  return result;
}

scratchcardsWorth();
