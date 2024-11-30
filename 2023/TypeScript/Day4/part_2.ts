import * as fs from 'fs';

function totalScratchcards(): number {
  try {
    const fileContent = fs.readFileSync('./input.txt', 'utf-8');
    const lines = fileContent.split('\n');
    const s: string[] = [];

    for (const line of lines) {
      s.push(line);
    }

    let total = 0;
    const multiString: { [key: number]: number } = {};

    for (let i = 0; i < s.length; i++) {
      multiString[i] = 1;
    }

    for (let i = 0; i < s.length; i++) {
      const game = s[i].split(': ');

      const games = game[1].split('|');

      const left = games[0].split(' ');
      const right = games[1].split(' ');

      const leftNumbers = parseNumbers(left);
      const rightNumbers = parseNumbers(right);

      const matches = intersection(leftNumbers, rightNumbers);

      const multiplier = multiString[i];

      total += multiplier;

      for (let j = 0; j < matches.length; j++) {
        multiString[i + j + 1] = (multiString[i + j + 1] || 0) + multiplier;
      }
    }

    console.log(total);
    return total;
  } catch (e) {
    console.log(`Error: ${e}`);
    return 0;
  }
}

function parseNumbers(nums: string[]): number[] {
  const result: number[] = [];
  for (const n of nums) {
    if (n.trim().length === 0) {
      continue;
    }
    const val = parseInt(n.trim(), 10);
    if (!isNaN(val)) {
      result.push(val);
    }
  }
  return result;
}

function intersection(a: number[], b: number[]): number[] {
  const result: number[] = [];
  const seen: { [key: number]: boolean } = {};
  for (const v of a) {
    seen[v] = true;
  }
  for (const v of b) {
    if (seen[v] === true) {
      result.push(v);
    }
  }
  return result;
}

totalScratchcards();
