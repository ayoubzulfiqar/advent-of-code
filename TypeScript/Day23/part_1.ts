type Point = [number, number];

function parse(taskInput: string[]): Map<Point, string> {
  const hikingMap = new Map<Point, string>();
  taskInput.forEach((line, rowIndex) => {
    line.split('').forEach((character, columnIndex) => {
      hikingMap.set([rowIndex, columnIndex], character);
    });
  });
  return hikingMap;
}

function* getNeighbors(hikingMap: Map<Point, string>, point: Point): Generator<Point> {
  if (hikingMap.get(point) === 'v') {
    yield [point[0] + 1, point[1]];
    return;
  }

  if (hikingMap.get(point) === '^') {
    yield [point[0] - 1, point[1]];
    return;
  }

  if (hikingMap.get(point) === '>') {
    yield [point[0], point[1] + 1];
    return;
  }

  if (hikingMap.get(point) === '<') {
    yield [point[0], point[1] - 1];
    return;
  }

  const directions: Point[] = [[0, -1], [1, 0], [-1, 0], [0, 1]];
  for (const direction of directions) {
    const newPoint: Point = [point[0] + direction[0], point[1] + direction[1]];
    if (!hikingMap.has(newPoint) || hikingMap.get(newPoint) === '#') {
      continue;
    }
    yield newPoint;
  }
}

function findTheLongestPath(hikingMap: Map<Point, string>, start: Point, end: Point): number {
  const toCheck: [Point, Map<Point, boolean>, number][] = [[start, new Map(), 0]];

  const costSoFar = new Map<Point, number>();
  costSoFar.set(start, 0);

  while (toCheck.length > 0) {
    const [currentPoint, path, currentCost] = toCheck.pop()!;
    const [rowIndex, columnIndex] = currentPoint;

    if (rowIndex === end[0] && columnIndex === end[1]) {
      continue;
    }

    for (const newPoint of getNeighbors(hikingMap, currentPoint)) {
      const newCost = costSoFar.get(currentPoint)! + 1;

      if (path.get(newPoint)) {
        continue;
      }

      if (!costSoFar.has(newPoint) || newCost > costSoFar.get(newPoint)!) {
        costSoFar.set(newPoint, newCost);

        const newPath = new Map(path);
        newPath.set(newPoint, true);

        toCheck.push([newPoint, newPath, newCost]);
      }
    }
  }

  return costSoFar.get(end)!;
}

function solution(taskInput: string[]): number {
  const hikingMap = parse(taskInput);
  let maxRows = 0;
  for (const [rowIndex] of hikingMap.keys()) {
    if (rowIndex > maxRows) {
      maxRows = rowIndex;
    }
  }

  let start: Point = [0, 0];
  let end: Point = [0, 0];
  for (const [point, tile] of hikingMap) {
    const [row, column] = point;
    if (row === 0 && tile === '.') {
      start = point;
    }
    if (row === maxRows && tile === '.') {
      end = point;
    }
  }

  return findTheLongestPath(hikingMap, start, end);
}

// Usage: LongISTheLongestHike();

// Note: File operations (reading from "input.txt") and console outputs are not included in the TypeScript code.
