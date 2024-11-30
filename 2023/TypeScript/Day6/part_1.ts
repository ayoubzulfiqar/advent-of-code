import * as fs from 'fs';

function desertIslandNumbers() {
  const input: string = fs.readFileSync('./input.txt', 'utf-8');
  const lines: string[] = input.split('\n');

  const times: number[] = lines[0]
    .split(/\s+/)
    .slice(1)
    .filter((s) => s.trim().length > 0)
    .map((s) => parseInt(s, 10))
    .filter((n) => !isNaN(n));

  const records: number[] = lines[1]
    .split(/\s+/)
    .slice(1)
    .filter((s) => s.trim().length > 0)
    .map((s) => parseInt(s, 10))
    .filter((n) => !isNaN(n));

  const races: Tuple2<number, number>[] = times.map((time, index) => new Tuple2(time, records[index]));

  let totalWays = 1;

  for (const race of races) {
    let waysToWin = 0;

    for (let holdTime = 0; holdTime < race.item1; holdTime++) {
      const distance = (race.item1 - holdTime) * holdTime;

      if (distance > race.item2) {
        waysToWin += 1;
      }
    }

    totalWays *= waysToWin;
  }

  console.log(`${totalWays}`);
}

class Tuple2<T1, T2> {
  readonly item1: T1;
  readonly item2: T2;

  constructor(item1: T1, item2: T2) {
    this.item1 = item1;
    this.item2 = item2;
  }
}

desertIslandNumbers();
