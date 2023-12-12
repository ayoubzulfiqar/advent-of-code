import * as fs from 'fs';

function empty(mat: string[][]): { er: { [key: number]: {} }, ec: { [key: number]: {} } } {
    const size = mat.length;
    const er: { [key: number]: {} } = {};
    const ec: { [key: number]: {} } = {};

    for (let i = 0; i < size; i++) {
        let r = true;
        let c = true;
        for (let j = 0; j < size; j++) {
            r = r && mat[i][j] !== '#';
            c = c && mat[j][i] !== '#';
        }
        if (r) {
            er[i] = {};
        }
        if (c) {
            ec[i] = {};
        }
    }
    return { er, ec };
}

interface Point {
    x: number;
    y: number;
}

function findPoints(mat: string[][]): Point[] {
    const p: Point[] = [];

    for (let i = 0; i < mat.length; i++) {
        for (let j = 0; j < mat[i].length; j++) {
            if (mat[i][j] === '#') {
                p.push({ x: i, y: j });
            }
        }
    }

    return p;
}

function shortestPathLengthInPairGalaxies() {
    const file = fs.readFileSync('./input.txt', 'utf8');
    const lines = file.split('\n');

    let mat: string[][] = lines.map((line) => line.split(''));
    const result = empty(mat);
    const el: { [key: number]: {} } = result.er;
    const ec: { [key: number]: {} } = result.ec;
    const points = findPoints(mat);
    const distS: number[] = [];
    const weight = 1000000;

    for (let i = 0; i < points.length; i++) {
        const p1 = points[i];
        for (let j = i + 1; j < points.length; j++) {
            const p2 = points[j];
            let x = 0;

            for (let k = Math.min(p1.x, p2.x); k < Math.max(p1.x, p2.x); k++) {
                if (el[k] !== undefined) {
                    x += weight;
                } else {
                    x++;
                }
            }

            for (let k = Math.min(p1.y, p2.y); k < Math.max(p1.y, p2.y); k++) {
                if (ec[k] !== undefined) {
                    x += weight;
                } else {
                    x++;
                }
            }

            distS.push(x);
        }
    }

    const total = distS.reduce((sum, dist) => sum + dist, 0);
    console.log(total);
}

shortestPathLengthInPairGalaxies();
