import * as fs from 'fs';

let _map: string[][] = [];

function TotalLoadOnNorthSupport(): number {
    const data: string[] = fs.readFileSync('./input.txt', 'utf-8').split('\n');
    _map = data.map(line => line.split(''));
    tilt();
    const result: number = countTotalLoad();
    console.log(result);
    return result;
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
        Array.from(_map.map(row => row[i]).reverse())
    );
    return _map;
}

function countTotalLoad(): number {
    const height: number = _map.length;
    return Array.from({ length: height }, (_, i) =>
        (height - i) * _map[i].filter(c => c === 'O').length
    ).reduce((acc, value) => acc + value, 0);
}

// Example usage
const result = TotalLoadOnNorthSupport();
console.log(result)