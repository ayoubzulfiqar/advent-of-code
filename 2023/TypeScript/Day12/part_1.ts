import * as fs from 'fs';

function count(cfg: string, nums: number[]): number {
    if (cfg === '') {
        if (nums.length === 0) {
            return 1;
        }
        return 0;
    }

    if (nums.length === 0) {
        if (cfg.includes('#')) {
            return 0;
        }
        return 1;
    }

    let result = 0;

    if (cfg[0] === '.' || cfg[0] === '?') {
        result += count(cfg.substring(1), nums);
    }

    if (cfg[0] === '#' || cfg[0] === '?') {
        if (
            nums[0] <= cfg.length &&
            !cfg.substring(0, nums[0]).includes('.') &&
            (nums[0] === cfg.length || cfg[nums[0]] !== '#')
        ) {
            if (nums[0] === cfg.length) {
                result += count('', nums.slice(1));
            } else {
                result += count(cfg.substring(nums[0] + 1), nums.slice(1));
            }
        }
    }

    return result;
}

function sumOfBrokenSprings(): void {
    const file = './input.txt';
    const content = fs.readFileSync(file, 'utf-8');
    const lines = content.split('\n');
    let total = 0;

    for (const line of lines) {
        const parts = line.split(' ');
        const cfg = parts[0];
        const numsStr = parts[1].split(',');
        const nums = numsStr.map((numStr) => parseInt(numStr, 10));
        total += count(cfg, nums);
    }

    console.log(total);
}

sumOfBrokenSprings();
