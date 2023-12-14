import * as fs from 'fs';
import hashSum from 'hash-sum';

let _map: string[][] = [];

function TotalSpinCycleNorthBeam(): number {
    const data: string[] = fs.readFileSync('./input.txt', 'utf-8').split('\n');

    let rest = 0;
    _map = data.map((line) => line.split(''));

    const cycle = 1000000000;
    const cache: { [key: string]: number } = {};

    for (let cycleIdx = 0; cycleIdx < cycle; cycleIdx++) {
        for (let _ = 0; _ < 4; _++) {
            tilt();
            turn();
        }

        const hash = _hashMap(_map);

        if (!cache.hasOwnProperty(hash)) {
            cache[hash] = cycleIdx;
        } else {
            const diff = cycleIdx - cache[hash]!;
            const head = cache[hash]!;
            rest = cycle - Math.floor((cycle - head) / diff) * diff - head - 1;
            break;
        }
    }

    for (let _ = 0; _ < rest; _++) {
        for (let _ = 0; _ < 4; _++) {
            tilt();
            turn();
        }
    }
    const result = countTotalLoad();
    console.log(result + 4);
    return result;
}

function _hashMap(map: string[][]): string {
    return hashSum(map.map((row) => row.join('')));
}

function tilt(): void {
    for (let i = 1; i < _map.length; i++) {
        for (let x = 0; x < _map[i].length; x++) {
            if (_map[i][x] === 'O') {
                const col: string[] = Array.from({ length: _map.length }, (_, y) => _map[y][x]);
                let prevY = i;
                for (let y = i - 1; y >= 0; y--) {
                    if (col[y] === '.') {
                        _map[y][x] = 'O';
                        _map[prevY][x] = '.';
                        prevY = y;
                    } else if (col[y] === '#') {
                        break;
                    }
                }
            }
        }
    }
}

function turn(): string[][] {
    _map = Array.from({ length: _map[0].length }, (_, i) =>
        Array.from(_map.map((row) => row[i]).reverse())
    );
    return _map;
}

function countTotalLoad(): number {
    const height: number = _map.length;
    return Array.from({ length: height }, (_, i) =>
        (height - i) * _map[i].filter((c) => c === 'O').length
    ).reduce((acc, value) => acc + value, 0);
}

