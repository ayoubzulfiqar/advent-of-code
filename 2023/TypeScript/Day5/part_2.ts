import * as fs from 'fs';

class SubMapping {
  readonly source: number;
  size!: number;
  offset!: number;

  constructor({ source, size, offset }: { source: number; size: number; offset: number }) {
    this.source = source;
    this.size = size;
    this.offset = offset;
  }
}

class Mapping {
  mappings: SubMapping[] = [];

  convert(input: number): number {
    for (const c of this.mappings) {
      if (input >= c.source && input <= c.source + c.size) {
        return input + c.offset;
      }
    }
    return input;
  }
}

function splitRangesAt(ranges: number[][], n: number): number[][] {
  for (let i = 0; i < ranges.length; i++) {
    const ss = ranges[i];
    if (n > ss[0] && n <= ss[1]) {
      ranges[i][1] = n - 1;
      ranges.push([n, ss[1]]);
      return ranges;
    }
  }
  return ranges;
}

function findMin(a: number, b: number): number {
  return (a < b) ? a : b;
}

function lowestInitialSeedNumber(): number {
  try {
    const fileContent = fs.readFileSync('./input.txt', 'utf-8');
    const lines = fileContent.split('\n');
    const seedPairs: number[][] = [];
    const mappings: Mapping[] = [];

    let currentMapping = new Mapping();
    for (const line of lines) {
      if (line.trim().length === 0) {
        continue;
      }

      if (line.includes('seeds: ')) {
        const seedsString = line.split('seeds: ');
        const seedList = seedsString[1].split(' ');

        for (let i = 0; i < seedList.length; i += 2) {
          seedPairs.push([
            parseInt(seedList[i], 10),
            parseInt(seedList[i], 10) + parseInt(seedList[i + 1], 10),
          ]);
        }
        continue;
      }

      if (line.includes('-')) {
        if (currentMapping.mappings.length > 0) {
          mappings.push(currentMapping);
        }
        currentMapping = new Mapping();
        continue;
      }

      const values = line.split(' ');

      currentMapping.mappings.push(new SubMapping({
        source: parseInt(values[1], 10),
        size: parseInt(values[2], 10),
        offset: parseInt(values[0], 10) - parseInt(values[1], 10),
      }));
    }
    if (currentMapping.mappings.length > 0) {
      mappings.push(currentMapping);
    }

    let lowest = -1;

    for (const pair of seedPairs) {
      let ranges = [pair];

      for (const mapping of mappings) {
        for (const subMapping of mapping.mappings) {
          ranges = splitRangesAt(ranges, subMapping.source);
        }

        for (let i = 0; i < ranges.length; i++) {
          ranges[i][0] = mapping.convert(ranges[i][0]);
          ranges[i][1] = mapping.convert(ranges[i][1]);
        }
      }

      for (let i = 0; i < ranges.length; i++) {
        if (lowest === -1) {
          lowest = ranges[i][0];
        }

        lowest = findMin(lowest, ranges[i][0]);
      }
    }

    console.log(lowest);
    return lowest;
  } catch (e) {
    console.log(`Error opening file: ${e}`);
    return 0;
  }
}

lowestInitialSeedNumber();
