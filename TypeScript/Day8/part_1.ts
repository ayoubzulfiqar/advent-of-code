import * as fs from 'fs/promises';

interface Node {
  Left: string;
  Right: string;
}

async function process(contents: string): Promise<number> {
  const nodesMap: { [key: string]: Node } = {};
  const lines = contents.split('\n');
  const turns = lines[0];

  const nodeRe = /^([A-Z][A-Z][A-Z]) = \(([A-Z][A-Z][A-Z]), ([A-Z][A-Z][A-Z])\)$/;

  for (let i = 2; i < lines.length; i++) {
    const line = lines[i];
    const nodeCaps = line.match(nodeRe);
    if (nodeCaps) {
      const source = nodeCaps[1];
      const lDest = nodeCaps[2];
      const rDest = nodeCaps[3];
      nodesMap[source] = { Left: lDest, Right: rDest };
    }
  }

  let currentNode = "AAA";
  let steps = 0;
  let turnIdx = 0;

  while (currentNode !== "ZZZ") {
    const turn = turns[turnIdx];

    switch (turn) {
      case 'L':
        currentNode = nodesMap[currentNode].Left;
        break;
      case 'R':
        currentNode = nodesMap[currentNode].Right;
        break;
      default:
        throw new Error("unexpected char in turns");
    }

    steps++;
    turnIdx = (turnIdx + 1) % turns.length;
  }

  return steps;
}

async function stepsToReachZ() {
  const filename = "input.txt";

  try {
    const contents = await fs.readFile(filename, 'utf-8');
    const result = await process(contents);
    console.log(`result = ${result}`);
  } catch (err: unknown) {
    if (err instanceof Error) {
      console.error(`Error opening or reading file: ${err.message}`);
    } else {
      console.error(`Unknown error: ${err}`);
    }
    // process.exit(1);
  }
}

stepsToReachZ();
