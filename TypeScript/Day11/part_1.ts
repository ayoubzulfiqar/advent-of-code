import * as fs from 'fs';

function addLines(mat: string[][]): string[][] {
    const size = mat.length;
    const nm: string[][] = [];

    for (let i = 0; i < size; i++) {
        let r = true;
        for (let j = 0; j < size; j++) {
            r = r && mat[i][j] !== '#';
        }
        nm.push([...mat[i]]);
        if (r) {
            nm.push(Array(size).fill('.'));
        }
    }

    return nm;
}

function addColumn(mat: string[][]): string[][] {
    const size = mat.length;
    const nm: string[][] = [];

    for (let i = 0; i < size; i++) {
        nm.push([]);
    }

    for (let i = 0; i < mat[0].length; i++) {
        let c = true;
        for (let j = 0; j < size; j++) {
            c = c && mat[j][i] !== '#';
            nm[j].push(mat[j][i]);
        }
        if (c) {
            for (let j = 0; j < size; j++) {
                nm[j].push('.');
            }
        }
    }

    return nm;
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

function shortestPathGalaxyLength() {
    const file = fs.readFileSync('./input.txt', 'utf8');
    const lines = file.split('\n');

    let mat: string[][] = lines.map((line) => line.split(''));
    mat = addLines(mat);
    mat = addColumn(mat);
    const points = findPoints(mat);

    const dists: number[] = [];

    for (let i = 0; i < points.length; i++) {
        const p1 = points[i];
        for (let j = i + 1; j < points.length; j++) {
            const p2 = points[j];
            const x = Math.abs(p1.x - p2.x) + Math.abs(p1.y - p2.y);
            dists.push(x);
        }
    }

    const total = dists.reduce((sum, dist) => sum + dist, 0);
    console.log(total);
}

shortestPathGalaxyLength();
