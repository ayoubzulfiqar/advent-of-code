import * as fs from 'fs';

async function engineGearRatioSum() {
  try {
    const file = fs.readFileSync('./input.txt', 'utf-8');
    const lines = file.split('\n');

    let total = 0;

    for (let r = 0; r < lines.length; r++) {
      const row = lines[r];
      for (let c = 0; c < row.length; c++) {
        const ch = row[c];

        if (ch !== '*') {
          continue;
        }

        const cs: { [key: string]: boolean } = {};

        for (const cr of [r - 1, r, r + 1]) {
          for (let cc of [c - 1, c, c + 1]) {
            if (
              cr < 0 ||
              cr >= lines.length ||
              cc < 0 ||
              cc >= lines[cr].length ||
              lines[cr][cc].charCodeAt(0) < '0'.charCodeAt(0) ||
              lines[cr][cc].charCodeAt(0) > '9'.charCodeAt(0)
            ) {
              continue;
            }

            while (
              cc > 0 &&
              lines[cr][cc - 1].charCodeAt(0) >= '0'.charCodeAt(0) &&
              lines[cr][cc - 1].charCodeAt(0) <= '9'.charCodeAt(0)
            ) {
              cc--;
            }

            cs[`(${cr},${cc})`] = true;
          }
        }

        if (Object.keys(cs).length !== 2) {
          continue;
        }

        const ns: number[] = [];

        for (const key in cs) {
          const match = /\((\d+),(\d+)\)/.exec(key);

          if (match !== null) {
            const cr = parseInt(match[1]!);
            let cc = parseInt(match[2]!);
            let s = '';

            while (
              cc < lines[cr].length &&
              lines[cr][cc].charCodeAt(0) >= '0'.charCodeAt(0) &&
              lines[cr][cc].charCodeAt(0) <= '9'.charCodeAt(0)
            ) {
              s += lines[cr][cc];
              cc++;
            }

            const n = parseInt(s);

            if (!isNaN(n)) {
              ns.push(n);
            }
          }
        }

        total += ns[0] * ns[1];
      }
    }

    console.log(total);
  } catch (e) {
    console.log(`Error: ${e}`);
  }
}

engineGearRatioSum()