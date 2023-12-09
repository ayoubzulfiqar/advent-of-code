import * as fs from 'fs';



function onlyNodesEndWithZ(): number {
  const input = fs.readFileSync("./input.txt", 'utf-8');

  const { moves, adj } = parse(input);
  const starts: string[] = [];

  for (const k in adj) {
    if (k[k.length - 1] === 'A') {
      starts.push(k);
    }
  }

  const numMoves: number[] = [];
  for (const start of starts) {
    const i = firstEndsWithZ(start, moves, adj);
    numMoves.push(i);
  }
  console.log(lcm(...numMoves))

  return lcm(...numMoves);
}

function firstEndsWithZ(curr: string, moves: string, adj: { [key: string]: string[] }): number {
  interface Entry {
    node: string;
    moveIdx: number;
  }

  const seen: { [key: string]: boolean } = {};
  const n = moves.length;
  let i = 0;

  while (!seen[`${curr}-${i % n}`]) {
    seen[`${curr}-${i % n}`] = true;

    if (curr[curr.length - 1] === 'Z') {
      return i;
    }

    if (moves[i % n] === 'L') {
      curr = adj[curr][0];
    } else {
      curr = adj[curr][1];
    }
    i += 1;
  }

  throw new Error("No valid path found");
}

function gcd(x: number, y: number): number {
  while (y !== 0) {
    [x, y] = [y, x % y];
  }

  return x;
}

function lcm(...nums: number[]): number {
  if (nums.length === 0) {
    return 0;
  }

  let res = nums[0];
  for (let i = 1; i < nums.length; i++) {
    res = (res * nums[i]) / gcd(res, nums[i]);
  }

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

onlyNodesEndWithZ()
