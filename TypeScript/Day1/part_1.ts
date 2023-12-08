import * as fs from 'fs';

async function sumCalibration(): Promise<number> {
  try {
    const fileContent: string = await fs.promises.readFile('test.txt', 'utf-8');
    const lines: string[] = fileContent.split('\n');

    let sum: number = 0;
    for (const line of lines) {
      const firstDigit: number | null = findFirstDigit(line);
      const lastDigit: number | null = findLastDigit(line);

      if (firstDigit !== null && lastDigit !== null) {
        const calibrationValue: number = firstDigit * 10 + lastDigit;
        sum += calibrationValue;
      }
    }

    return sum;
  } catch (e) {
    console.error('Error:', e);
    return 0;
  }
}

function findFirstDigit(s: string): number | null {
  for (let i = 0; i < s.length; i++) {
    const char: string = s[i];
    if (isDigit(char)) {
      const digit: number | null = toInt(char);
      if (digit !== null) {
        return digit;
      }
    }
  }
  return null;
}

function findLastDigit(s: string): number | null {
  for (let i = s.length - 1; i >= 0; i--) {
    const char: string = s[i];
    if (isDigit(char)) {
      const digit: number | null = toInt(char);
      if (digit !== null) {
        return digit;
      }
    }
  }
  return null;
}

function isDigit(char: string): boolean {
  return /\d/.test(char);
}

function toInt(char: string): number | null {
  if (/\d/.test(char)) {
    return parseInt(char, 10);
  }
  return null;
}


