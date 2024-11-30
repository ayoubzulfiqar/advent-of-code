import * as fs from 'fs';

function main() {
    const histories: number[][] = [];

    try {
        const file: Buffer = fs.readFileSync('./input.txt');
        const data: string[] = file.toString().split('\n');

        for (const line of data) {
            const values: number[] = parseInts(line);
            histories.push(values);
        }
    } catch (e) {
        console.log(`error reading input.txt: ${e}`);
        process.exit(1);
    }

    const extrapolatedValues: number[] = [];

    for (const history of histories) {
        const subLists: number[][] = [history];
        let allZeroes = false;

        while (!allZeroes) {
            const sublist: number[] = [];

            for (let i = 0; i < subLists[subLists.length - 1].length - 1; i++) {
                const difference = subLists[subLists.length - 1][i + 1] - subLists[subLists.length - 1][i];
                sublist.push(difference);
            }

            allZeroes = allZeroesInList(sublist);
            subLists.push(sublist);
        }

        subLists[subLists.length - 1].push(0);

        for (let i = subLists.length - 2; i >= 0; i--) {
            const extrapolatedValue = subLists[i][subLists[i].length - 1] + subLists[i + 1][subLists[i + 1].length - 1];
            subLists[i].push(extrapolatedValue);
        }

        extrapolatedValues.push(subLists[0][subLists[0].length - 1]);
    }

    const ans = sum(extrapolatedValues);

    console.log(ans);
}

function parseInts(input: string): number[] {
    const nums: number[] = [];
    const fields: string[] = input.split(' ');

    for (const field of fields) {
        const num: number = parseInt(field) || 0;
        nums.push(num);
    }

    return nums;
}

function allZeroesInList(nums: number[]): boolean {
    for (const num of nums) {
        if (num !== 0) {
            return false;
        }
    }
    return true;
}

function sum(nums: number[]): number {
    return nums.reduce((result, num) => result + num, 0);
}

main();
