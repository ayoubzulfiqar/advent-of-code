import * as fs from 'fs';

class Point {
    R: number;
    C: number;
    S: number;

    constructor(R: number, C: number, S: number) {
        this.R = R;
        this.C = C;
        this.S = S;
    }
}

function sixtyFourElfSteps() {
    const fileContent = fs.readFileSync('./input.txt', 'utf8');
    const grid: string[] = fileContent.split('\n');

    let sr = 0, sc = 0;
    for (let r = 0; r < grid.length; r++) {
        for (let c = 0; c < grid[r].length; c++) {
            if (grid[r][c] == 'S') {
                sr = r;
                sc = c;
                break;
            }
        }
    }

    const ans: Map<Point, boolean> = new Map();
    const seen: Map<Point, boolean> = new Map();
    const q: Point[] = [new Point(sr, sc, 64)];

    while (q.length > 0) {
        const current: Point = q.shift()!;

        const r: number = current.R;
        const c: number = current.C;
        const s: number = current.S;

        if (s % 2 === 0) {
            ans.set(current, true);
        }
        if (s === 0) {
            continue;
        }

        const moves: number[][] = [
            [1, 0],
            [-1, 0],
            [0, 1],
            [0, -1],
        ];
        for (const move of moves) {
            const nr: number = r + move[0];
            const nc: number = c + move[1];

            if (
                nr < 0 ||
                nr >= grid.length ||
                nc < 0 ||
                nc >= grid[0].length ||
                grid[nr][nc] === '#' ||
                seen.get(new Point(nr, nc, 0))
            ) {
                continue;
            }

            seen.set(new Point(nr, nc, 0), true);
            q.push(new Point(nr, nc, s - 1));
        }
    }

    console.log(ans.size - 1);
}

sixtyFourElfSteps();
