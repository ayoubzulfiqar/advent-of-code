import * as fs from 'fs';

function hashed(s: string): number {
    let v = 0;

    for (const ch of s) {
        v += ch.codePointAt(0)!;
        v *= 17;
        v %= 256;
    }

    return v;
}

function powerOfResultLens(): void {
    const boxes: string[][] = Array.from({ length: 256 }, () => []);
    const focalLengths: { [key: string]: number } = {};

    const instructions: string[] = fs
        .readFileSync('./input.txt', 'utf-8')
        .split(',');

    for (const instruction of instructions) {
        if (instruction.includes('-')) {
            const label = instruction.slice(0, -1);
            const index = hashed(label);

            boxes[index] = boxes[index].filter((l) => l !== label);
        } else {
            const [label, length] = instruction.split('=');
            const lengthValue = parseInt(length, 10);

            const index = hashed(label);
            if (!boxes[index].includes(label)) {
                boxes[index].push(label);
            }

            focalLengths[label] = lengthValue;
        }
    }

    let total = 0;

    for (let i = 0; i < boxes.length; i++) {
        for (let j = 0; j < boxes[i].length; j++) {
            const label = boxes[i][j];
            total += (i + 1) * (j + 1) * focalLengths[label]!;
        }
    }

    console.log(total);
}

powerOfResultLens();
