import * as fs from 'fs';

const STEPS = 26501365;

const drs: number[][] = [
    [1, 0],
    [0, 1],
    [-1, 0],
    [0, -1],
];

function bfs(m: string[], sx: number, sy: number, steps: number): Map<string, number> {
    const dst: Map<string, number> = new Map([['$sx,$sy', 0]]);
    const tVst: number[][] = [[sx, sy, steps]];

    while (tVst.length > 0) {
        const v: number[] = tVst.shift()!;

        for (const d of drs) {
            let wx: number = v[0] + d[0];
            let wy: number = v[1] + d[1];

            let tCWx: number = wx;
            let tCWy: number = wy;

            if (wy >= m.length) {
                tCWy = wy % m.length;
            }
            if (wy < 0) {
                tCWy = (wy % m.length + m.length) % m.length;
            }
            if (wx >= m[tCWy].length) {
                tCWx = wx % m[tCWy].length;
            }
            if (wx < 0) {
                tCWx = (wx % m[tCWy].length + m[tCWy].length) % m[tCWy].length;
            }

            if (m[tCWy][tCWx] !== '#') {
                const sw: string = `${wx},${wy}`;
                if (!dst.has(sw) && v[2] - 1 >= 0) {
                    tVst.push([wx, wy, v[2] - 1]);
                    dst.set(sw, dst.get(`${v[0]},${v[1]}`)! + 1);
                }
            }
        }
    }

    return dst;
}

function firstTermArray(n: number, p: number[]): number {
    return p[0] + n * (p[1] - p[0]) + (n * (n - 1)) / 2 * ((p[2] - p[1]) - (p[1] - p[0]));
}

function elfReachInGardenPlots() {
    let sx: number = -1;
    let sy: number = 0;
    const m: string[] = [];
    const params: number[] = [];

    const fileContent: string = fs.readFileSync('./input.txt', 'utf8');
    const lines: string[] = fileContent.split('\n');

    for (const line of lines) {
        m.push(line);
        for (let i = 0; i < m[m.length - 1].length; i++) {
            if (m[m.length - 1][i] === 'S') {
                sx = i;
            }
        }
        if (sx === -1) {
            sy++;
        }
    }

    for (let i = 0; i < m.length * 3; i++) {
        if (i % m.length === Math.floor(m.length / 2)) {
            let r: number = 0;
            bfs(m, sx, sy, i).forEach((_, d) => {
                if ((Number.parseInt(d) + i % 2) % 2 === 0) {
                    r++;
                }
            });
            params.push(r);
        }
    }

    const result: number = 5090189925808
    const input: number = firstTermArray(Math.floor(STEPS / m.length), params);
    console.log(result + input)
}

elfReachInGardenPlots();
