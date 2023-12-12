const PREV = 0;
const CUR = 1;

const input = "./input.txt";

type PipeMap = string[][];

interface Point {
    row: number;
    col: number;
    sym?: MapSymbol;
    dir?: boolean;
}

const pipeMap: PipeMap = input.split('\n').map(x => x.split(''));
const onMap = (pipeMap: PipeMap) => (p: Point) =>
    !(
        p.row < 0 ||
        p.col < 0 ||
        p.row >= pipeMap.length ||
        p.col >= pipeMap[0].length
    )

type MapSymbol = '|' | '-' | 'L' | 'J' | 'F' | '7';
const advanceRow = (point: Point, offset: number, sym?: MapSymbol, dir?: boolean): Point => ({ row: point.row + offset, col: point.col, sym, dir })
const advanceCol = (point: Point, offset: number, sym?: MapSymbol, dir?: boolean): Point => ({ row: point.row, col: point.col + offset, sym, dir })
const advanceBoth = (point: Point, oRow: number, oCol: number, sym?: MapSymbol, dir?: boolean): Point => advanceCol(advanceRow(point, oRow), oCol, sym, dir);

const inner = (point: Point, sym: MapSymbol, reverse: boolean): Point[] => {
    if (sym === "|") return [advanceCol(point, reverse ? -1 : 1, "|", reverse)] // from bottom
    if (sym === "-") return [advanceRow(point, reverse ? -1 : 1, "-", reverse)] // from left
    if (sym === "L") return reverse ? [advanceCol(point, -1, "|", reverse), advanceBoth(point, 1, -1, "L", reverse), advanceRow(point, 1, "-", !reverse)] : [] // from right
    if (sym === "J") return reverse ? [] : [advanceCol(point, 1, "|", reverse), advanceBoth(point, 1, 1, "J", reverse), advanceRow(point, 1, "-", reverse)]  // from left
    if (sym === "F") return reverse ? [advanceCol(point, -1, "|", reverse), advanceBoth(point, -1, -1, "F", reverse), advanceRow(point, -1, "-", reverse)] : [] // from bottom
    if (sym === "7") return reverse ? [] : [advanceCol(point, 1, "|", reverse), advanceBoth(point, -1, 1, "7", reverse), advanceRow(point, -1, "-", !reverse)] // from bottom
    throw new Error("Invalid symbol");
}

enum Going {
    UP,
    DOWN,
    LEFT, RIGHT
}

const direction = (from: Point, to: Point): Going => {
    const hOffset = to.col - from.col;
    if (hOffset !== 0) return hOffset > 0 ? Going.RIGHT : Going.LEFT;
    const vOffset = to.row - from.row;
    return vOffset > 0 ? Going.DOWN : Going.UP;
}

const isReverse = (sym: MapSymbol, d: Going) => {
    if (sym === "|") return d === Going.UP ? false : true;
    if (sym === "-") return d === Going.RIGHT ? false : true;
    if (sym === "L") return d === Going.LEFT ? false : true;
    if (sym === "J") return d === Going.RIGHT ? false : true;
    if (sym === "F") return d === Going.UP ? false : true;
    if (sym === "7") return d === Going.UP ? false : true;
    throw new Error("invalid symbol");
}

const next = (cur: Point, sym: string): Point[] => {
    if (sym === '|') return [advanceRow(cur, -1), advanceRow(cur, 1)];
    if (sym === '-') return [advanceCol(cur, -1), advanceCol(cur, 1)];
    if (sym === 'L') return [advanceRow(cur, -1), advanceCol(cur, 1)];
    if (sym === 'J') return [advanceRow(cur, -1), advanceCol(cur, -1)];
    if (sym === 'F') return [advanceRow(cur, 1), advanceCol(cur, 1)];
    if (sym === '7') return [advanceRow(cur, 1), advanceCol(cur, -1)];
    if (sym === '.') return []

    throw new Error(`invalid symbol ${sym}`);
}

const areEqual = (a: Point, b: Point) => a.row === b.row && a.col === b.col;

const pointsTo = (pipeMap: PipeMap, from: Point, to: Point) =>
    next(from, pipeMap[from.row][from.col]).some(p => areEqual(p, to));


const surrounding = (pipeMap: PipeMap, p: Point): Point[] => [advanceCol(p, 1), advanceCol(p, -1), advanceRow(p, 1), advanceRow(p, -1)].filter(onMap(pipeMap))

const step = (pipeMap: PipeMap, previous: Point, current: Point) => {
    const pipe = pipeMap[current.row][current.col];
    const nextPoint = next(current, pipe).filter(p => !areEqual(p, previous))[0];
    return nextPoint;
}



const findStart = (pipeMap: PipeMap): Point => {
    for (let rowI = 0; rowI < pipeMap.length; rowI++) {
        for (let colI = 0; colI < pipeMap[rowI].length; colI++) {
            if (pipeMap[rowI][colI] === "S") {
                return {
                    row: rowI,
                    col: colI
                }
            }
        }
    }

    throw new Error("No start");
}

const part1 = (pipeMap: PipeMap) => {
    const start = findStart(pipeMap);
    let startPoints = surrounding(pipeMap, start).filter(p => pointsTo(pipeMap, p, start)).map(p => [start, p]);
    let steps = 1;

    while (!startPoints.every(points => areEqual(points[CUR], startPoints[0][CUR]))) {
        startPoints = startPoints.map(points => [points[CUR], step(pipeMap, points[PREV], points[CUR])])
        steps++;
    }

    return steps;
}

const part2 = (pipeMap: PipeMap, r: boolean) => {
    const marks = pipeMap.map(row => row.map<boolean | string>(_ => false));
    const mark = (p: Point, s: string) => marks[p.row][p.col] = s;
    const isMarked = (p: Point) => marks[p.row][p.col] !== false;
    const start = findStart(pipeMap);
    mark(start, 'U');
    let prevPoint = start;
    let surroundingPoints = surrounding(pipeMap, start).filter(p => pointsTo(pipeMap, p, start));
    let currentPoint = r ? surroundingPoints.reverse()[0] : surroundingPoints[0];
    const innerPoints: Point[] = [];

    let turnDirection = 0;
    // Add inner points
    while (pipeMap[currentPoint.row][currentPoint.col] !== 'S') {
        const sym = pipeMap[currentPoint.row][currentPoint.col] as MapSymbol;
        innerPoints.push(...inner(currentPoint, sym, isReverse(sym, direction(prevPoint, currentPoint))).filter(onMap(pipeMap)))

        mark(currentPoint, 'U');
        const nextPoint = step(pipeMap, prevPoint, currentPoint);
        prevPoint = currentPoint;
        currentPoint = nextPoint;
    }

    let innerCount = 0;
    while (innerPoints.length > 0) {
        const cur = innerPoints.shift();
        if (cur === undefined) throw new Error('Impossible');
        if (!isMarked(cur)) {
            innerCount++;
            innerPoints.push(...inner(cur, cur.sym!, cur.dir!).filter(onMap(pipeMap)))
            mark(cur, 'X')
        }
    }
    return innerCount;
}

console.log(part1(pipeMap));
console.log(part2(pipeMap, true), "or", part2(pipeMap, false));
