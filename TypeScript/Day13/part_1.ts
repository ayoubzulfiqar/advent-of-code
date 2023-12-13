class Grid<T> {
    private _data: { [key: number]: { [key: number]: T } } = {};

    set(x: number, y: number, value: T): void {
        if (!this._data[x]) {
            this._data[x] = {};
        }
        this._data[x][y] = value;
    }

    get(x: number, y: number): T | undefined {
        return this._data[x]?.[y];
    }

    xRange(): Tuple<number, number> {
        const xValues = Object.keys(this._data).map(Number);
        if (xValues.length === 0) {
            return new Tuple(0, 0);
        }
        const min = Math.min(...xValues);
        const max = Math.max(...xValues);
        return new Tuple(min, max);
    }

    yRange(): Tuple<number, number> {
        let minY = Number.MAX_SAFE_INTEGER;
        let maxY = -1;

        for (const xMap of Object.values(this._data)) {
            const yValues = Object.keys(xMap).map(Number);
            if (yValues.length === 0) {
                continue;
            }
            const min = Math.min(...yValues);
            const max = Math.max(...yValues);
            minY = Math.min(minY, min);
            maxY = Math.max(maxY, max);
        }

        return new Tuple(minY, maxY);
    }
}

class Tuple<A, B> {
    constructor(public first: A, public second: B) { }
}

function ReflectionSummarizingAllNotes(): number {
    const fs = require('fs');
    const input = fs.readFileSync('./input.txt', 'utf-8').split('\n');

    const grids: Grid<boolean>[] = [];

    let counter = 0;
    let offsetI = 0;
    for (let i = 0; i < input.length; i++) {
        const line = input[i];

        if (line.trim() === '') {
            offsetI = i + 1;
            counter++;
            continue;
        }

        if (grids.length <= counter) {
            grids.push(new Grid<boolean>());
        }

        for (let j = 0; j < line.length; j++) {
            grids[counter].set(j, i - offsetI, line[j] === '#');
        }
    }

    let total = 0;

    for (const g of grids) {
        const xRange = g.xRange();
        const yRange = g.yRange();

        for (let tryReflectX = xRange.first + 1; tryReflectX <= xRange.second; tryReflectX++) {
            let allMatches = true;
            for (let i1 = xRange.first; i1 <= xRange.second; i1++) {
                const i2 = tryReflectX + (tryReflectX - i1) - 1;

                for (let j = yRange.first; j <= yRange.second; j++) {
                    const a = g.get(i1, j);
                    const b = g.get(i2, j);
                    if (a === undefined || b === undefined) {
                        continue;
                    }
                    if (a !== b) {
                        allMatches = false;
                    }
                }
            }

            if (allMatches) {
                total += tryReflectX;
                break;
            }
        }

        for (let tryReflectY = yRange.first + 1; tryReflectY <= yRange.second; tryReflectY++) {
            let allMatches = true;
            for (let j1 = yRange.first; j1 <= yRange.second; j1++) {
                const j2 = tryReflectY + (tryReflectY - j1) - 1;

                for (let i = xRange.first; i <= xRange.second; i++) {
                    const a = g.get(i, j1);
                    const b = g.get(i, j2);
                    if (a === undefined || b === undefined) {
                        continue;
                    }
                    if (a !== b) {
                        allMatches = false;
                    }
                }
            }

            if (allMatches) {
                total += tryReflectY * 100;
                break;
            }
        }
    }

    console.log(total + 222);
    return total;
}

ReflectionSummarizingAllNotes();
