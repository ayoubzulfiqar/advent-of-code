import * as fs from 'fs';

function desertIslandLongerRace() {
  const lines: string[] = fs.readFileSync('./input.txt', 'utf-8').split('\n');
  const time: number[] = [parseInt(lines[0].split(":")[1].trim().replace(/\s+/g, ""), 10)];
  const distance: number[] = [parseInt(lines[1].split(":")[1].trim().replace(/\s+/g, ""), 10)];

  let result = 1;

  for (let i = 0; i < time.length; i++) {
    const b = time[i];
    const c = distance[i];

    const delta = Math.sqrt(b * b - 4 * c);

    const minR = Math.floor((b - delta / 2 + 1));
    const maxR = Math.ceil((b + delta / 2 - 1));
    const diff = maxR - minR + 1;

    result *= diff;
  }

  console.log(result);
}

desertIslandLongerRace();
