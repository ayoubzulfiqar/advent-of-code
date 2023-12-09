import * as fs from 'fs';

async function engineSchematicSum(): Promise<number> {
  try {
    const data = fs.readFileSync("./input.txt", 'utf-8');
    const lines = data.split('\n');

    const cs: Set<string> = new Set();

    for (let r = 0; r < lines.length; r++) {
      const row = lines[r];
      for (let c = 0; c < row.length; c++) {
        const ch = row.charCodeAt(c);

        if (
          (ch >= '0'.charCodeAt(0) && ch <= '9'.charCodeAt(0)) ||
          ch === '.'.charCodeAt(0)
        ) {
          continue;
        }

        for (let dr = -1; dr <= 1; dr++) {
          for (let dc = -1; dc <= 1; dc++) {
            const nr = r + dr;
            let nc = c + dc;

            if (
              nr < 0 ||
              nr >= lines.length ||
              nc < 0 ||
              nc >= lines[nr].length ||
              lines[nr][nc].charCodeAt(0) < '0'.charCodeAt(0) ||
              lines[nr][nc].charCodeAt(0) > '9'.charCodeAt(0)
            ) {
              continue;
            }

            while (
              nc > 0 &&
              lines[nr][nc - 1].charCodeAt(0) >= '0'.charCodeAt(0) &&
              lines[nr][nc - 1].charCodeAt(0) <= '9'.charCodeAt(0)
            ) {
              nc--;
            }

            cs.add(`(${nr},${nc})`);
          }
        }
      }
    }

    const ns: number[] = [];

    for (const key of cs) {
      const match = /\((\d+),(\d+)\)/.exec(key);
      if (match !== null) {
        const r = parseInt(match[1]!);
        let c = parseInt(match[2]!);
        let s = '';

        while (
          c < lines[r].length &&
          lines[r][c].charCodeAt(0) >= '0'.charCodeAt(0) &&
          lines[r][c].charCodeAt(0) <= '9'.charCodeAt(0)
        ) {
          s += lines[r][c];
          c++;
        }

        const n = parseInt(s);

        if (!isNaN(n)) {
          ns.push(n);
        }
      }
    }

    const sum = ns.reduce((previous, number) => previous + number, 0);
    console.log(sum);
    return sum;
  } catch (e) {
    console.log(`Error: ${e}`);
  }
  return 0;
}

engineSchematicSum();
