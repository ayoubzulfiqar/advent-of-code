import * as fs from 'fs';

function brickSafelyChosen(): number {
    const file = fs.readFileSync('./input.txt', 'utf8');
    const bricks: number[][] = [];

    file.split('\n').forEach((line) => {
        const values = line.replace('~', ',').split(',');
        const brick: number[] = [];
        values.forEach((v) => {
            const num = parseInt(v, 10);
            brick.push(num);
        });
        bricks.push(brick);
    });

    bricks.sort((a, b) => a[2] - b[2]);

    function overlaps(a: number[], b: number[]): boolean {
        return max(a[0], b[0]) <= min(a[3], b[3]) &&
            max(a[1], b[1]) <= min(a[4], b[4]);
    }

    for (let index = 0; index < bricks.length; index++) {
        let maxZ = 1;
        for (let i = 0; i < index; i++) {
            if (overlaps(bricks[index], bricks[i])) {
                maxZ = max(maxZ, bricks[i][5] + 1);
            }
        }
        bricks[index][5] -= bricks[index][2] - maxZ;
        bricks[index][2] = maxZ;
    }

    bricks.sort((a, b) => a[2] - b[2]);

    const kSupportsV: Record<number, Set<number>> = {};
    const vSupportsK: Record<number, Set<number>> = {};

    for (let i = 0; i < bricks.length; i++) {
        kSupportsV[i] = new Set();
        vSupportsK[i] = new Set();
    }

    for (let j = 0; j < bricks.length; j++) {
        for (let i = 0; i < j; i++) {
            if (overlaps(bricks[i], bricks[j]) && bricks[j][2] === bricks[i][5] + 1) {
                kSupportsV[i].add(j);
                vSupportsK[j].add(i);
            }
        }
    }

    let total = 0;

    for (let i = 0; i < bricks.length; i++) {
        let satisfies = true;
        for (let j of kSupportsV[i]) {
            if (vSupportsK[j].size < 2) {
                satisfies = false;
                break;
            }
        }
        if (satisfies) {
            total++;
        }
    }

    console.log(total);
    return total;
}

function max(a: number, b: number): number {
    return a > b ? a : b;
}

function min(a: number, b: number): number {
    return a < b ? a : b;
}

brickSafelyChosen();
