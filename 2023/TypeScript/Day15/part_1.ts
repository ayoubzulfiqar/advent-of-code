import * as fs from 'fs';

function hash(s: string): number {
  let v = 0;

  for (let i = 0; i < s.length; i++) {
    v += s.charCodeAt(i);
    v *= 17;
    v %= 256;
  }

  return v;
}

function hashSumOfResults(): void {
  const input: string = fs.readFileSync('./input.txt', 'utf-8');
  const inputList: string[] = input.split(',');

  let result: number = 0;
  for (const str of inputList) {
    result += hash(str);
  }

  console.log(result);
}

hashSumOfResults();
