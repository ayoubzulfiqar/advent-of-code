import * as fs from 'fs';


function stepsToReachZ(): number {
  const input = fs.readFileSync("./input.txt", 'utf-8');
  const { moves, adj } = parse(input);
  let curr = "AAA";
  const target = "ZZZ";
  const n = moves.length;
  let i = 0;
  let res = 0;

  while (curr !== target) {
      if (moves[i % n] === 'L') {
          curr = adj[curr][0];
      } else {
          curr = adj[curr][1];
      }
      i += 1;
      res += 1;
  }
 console.log(res)
  return res;
}

function parse(input: string): { moves: string; adj: { [key: string]: string[] } } {
  const blocks = input.replace(/\r\n/g, '\n').split('\n\n');
  const adj: { [key: string]: string[] } = {};
  const moves = blocks[0];

  for (const line of blocks[1].split('\n')) {
      const contents = line.split(' = ');
      const source = contents[0];
      const dest = contents[1].split(', ');
      const left = dest[0].slice(1);
      const right = dest[1].slice(0, -1);

      adj[source] = [left, right];
  }

  return { moves, adj };
}

stepsToReachZ()