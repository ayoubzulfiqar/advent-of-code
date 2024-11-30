import * as fs from 'fs';

function sumOfPossibleArrangementCount(): number {
    try {
        const input = fs.readFileSync('./input.txt', 'utf-8');
        const recordsAndGroups = parse(input);
        const records = recordsAndGroups[0];
        const groups = recordsAndGroups[1];
        let res = 0;

        for (let i = 0; i < records.length; i++) {
            res += solve(unfoldRecord(records[i]), unfoldGroup(groups[i]));
        }

        console.log(res);
        return res;
    } catch (e) {
        console.error(e);
        return 0;
    }
}

function unfoldRecord(record: string): string {
    const res: string[] = [];
    for (let i = 0; i < record.length * 5; i++) {
        if (i !== 0 && i % record.length === 0) {
            res.push('?');
        }
        res.push(record[i % record.length]);
    }

    return res.join('');
}

function unfoldGroup(group: number[]): number[] {
    const res: number[] = [];
    for (let i = 0; i < group.length * 5; i++) {
        res.push(group[i % group.length]);
    }

    return res;
}

function solve(record: string, group: number[]): number {
    const cache = Array.from({ length: record.length }, () =>
        Array(group.length + 1).fill(-1)
    );

    return dp(0, 0, record, group, cache);
}

function dp(
    i: number,
    j: number,
    record: string,
    group: number[],
    cache: number[][]
): number {
    if (i >= record.length) {
        if (j < group.length) {
            return 0;
        }
        return 1;
    }

    if (cache[i][j] !== -1) {
        return cache[i][j];
    }

    let res = 0;
    if (record[i] === '.') {
        res = dp(i + 1, j, record, group, cache);
    } else {
        if (record[i] === '?') {
            res += dp(i + 1, j, record, group, cache);
        }
        if (j < group.length) {
            let count = 0;
            for (let k = i; k < record.length; k++) {
                if (
                    count > group[j] ||
                    record[k] === '.' ||
                    (count === group[j] && record[k] === '?')
                ) {
                    break;
                }
                count += 1;
            }

            if (count === group[j]) {
                if (i + count < record.length && record[i + count] !== '#') {
                    res += dp(i + count + 1, j + 1, record, group, cache);
                } else {
                    res += dp(i + count, j + 1, record, group, cache);
                }
            }
        }
    }

    cache[i][j] = res;
    return res;
}

function parse(input: string): [string[], number[][]] {
    const records: string[] = [];
    const groups: number[][] = [];

    for (const line of input.split('\n')) {
        const parts = line.split(' ');
        records.push(parts[0]);
        const group = parts[1].split(',').map((num) => parseInt(num, 10));
        groups.push(group);
    }

    return [records, groups];
}

sumOfPossibleArrangementCount();
