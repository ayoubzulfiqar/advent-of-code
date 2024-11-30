import * as fs from 'fs';

class SubMapping {
  source: number;
  size: number;
  offset: number;

  constructor({ source, size, offset }: { source: number; size: number; offset: number }) {
    this.source = source;
    this.size = size;
    this.offset = offset;
  }
}

class Mapping implements Iterable<SubMapping> {
  private _list: SubMapping[] = [];

  get length(): number {
    return this._list.length;
  }

  [Symbol.iterator](): Iterator<SubMapping> {
    return this._list[Symbol.iterator]();
  }

  operator(index: number): SubMapping {
    return this._list[index];
  }

  set(index: number, value: SubMapping): void {
    this._list[index] = value;
  }

  add(value: SubMapping): void {
    this._list.push(value);
  }
}

function lowestLocationSeedNumber(): number {
  try {
    const fileContent = fs.readFileSync('./input.txt', 'utf-8');
    const lines = fileContent.split('\n');
    const seeds: number[] = [];
    const mappings: Mapping[] = [];
    let currentMapping = new Mapping();

    for (const line of lines) {
      if (line.trim().length === 0) {
        continue;
      }

      if (line.includes("seeds: ")) {
        const seedsString = line.split("seeds: ");
        const seedList = seedsString[1].split(" ");
        for (const seedItem of seedList) {
          const seed = parseInt(seedItem, 10);
          if (!isNaN(seed)) {
            seeds.push(seed);
          }
        }
        continue;
      }

      if (line.includes("-")) {
        if (currentMapping.length > 0) {
          mappings.push(currentMapping);
        }
        currentMapping = new Mapping();
        continue;
      }

      const values = line.split(' ');

      const source = parseInt(values[1], 10);
      const size = parseInt(values[2], 10);

      if (isNaN(source) || isNaN(size)) {
        console.log("Error converting source or size to integer");
        return 0;
      }

      const offset = sti(values[0]) - sti(values[1]);

      currentMapping.add(new SubMapping({
        source: source,
        size: size,
        offset: offset,
      }));
    }

    if (currentMapping.length > 0) {
      mappings.push(currentMapping);
    }

    let lowest = -1;

    for (const seed of seeds) {
      let val = seed;

      for (const mapping of mappings) {
        for (const subMapping of mapping) {
          if (val >= subMapping.source && val <= subMapping.source + subMapping.size) {
            val += subMapping.offset;
            break;
          }
        }
      }

      if (lowest === -1 || val < lowest) {
        lowest = val;
      }
    }

    console.log(lowest);
    return lowest;
  } catch (e) {
    console.log(`Error: ${e}`);
    return 0;
  }
}

function sti(s: string): number {
  const i = parseInt(s, 10);
  if (isNaN(i)) {
    console.log("Error converting string to integer");
  }
  return isNaN(i) ? 0 : i;
}

lowestLocationSeedNumber();
