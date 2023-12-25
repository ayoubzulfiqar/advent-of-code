import * as fs from "fs";
import * as readline from "readline";

interface Hailstone {
    px: number;
    py: number;
    pz: number;
    vx: number;
    vy: number;
    vz: number;
}

const hailstones: Hailstone[] = [];

const velocitiesX: Record<number, number[]> = {};
const velocitiesY: Record<number, number[]> = {};
const velocitiesZ: Record<number, number[]> = {};

const getRockVelocity = (velocities: Record<number, number[]>): number => {
    let possibleV: number[] = [];
    for (let x = -1000; x <= 1000; x++) {
        possibleV.push(x);
    }

    Object.keys(velocities).forEach((velocity) => {
        const vel = parseInt(velocity, 10);
        if (velocities[vel].length < 2) {
            return;
        }

        let newPossibleV: number[] = [];
        possibleV.forEach((possible) => {
            if ((velocities[vel][0] - velocities[vel][1]) % (possible - vel) === 0) {
                newPossibleV.push(possible);
            }
        });

        possibleV = newPossibleV;
    });

    return possibleV[0];
};

async function coordinatesOfInitialPosition() {
    const fileStream = fs.createReadStream("./input.txt");

    const rl = readline.createInterface({
        input: fileStream,
        crlfDelay: Infinity,
    });

    for await (const line of rl) {
        const [positions, velocity] = line.split(' @ ');
        const [px, py, pz] = positions.split(', ').map((n) => Number(n));
        const [vx, vy, vz] = velocity.split(', ').map((n) => Number(n));

        if (!velocitiesX[vx]) {
            velocitiesX[vx] = [px];
        } else {
            velocitiesX[vx].push(px);
        }

        if (!velocitiesY[vy]) {
            velocitiesY[vy] = [py];
        } else {
            velocitiesY[vy].push(py);
        }

        if (!velocitiesZ[vz]) {
            velocitiesZ[vz] = [pz];
        } else {
            velocitiesZ[vz].push(pz);
        }

        hailstones.push({ px, py, pz, vx, vy, vz });
    }

    let possibleV: number[] = [];
    for (let x = -1000; x <= 1000; x++) {
        possibleV.push(x);
    }

    const rvx = getRockVelocity(velocitiesX);
    const rvy = getRockVelocity(velocitiesY);
    const rvz = getRockVelocity(velocitiesZ);

    const results: Record<number, number> = {};
    for (let i = 0; i < hailstones.length; i++) {
        for (let j = i + 1; j < hailstones.length; j++) {
            const stoneA = hailstones[i];
            const stoneB = hailstones[j];

            const ma = (stoneA.vy - rvy) / (stoneA.vx - rvx);
            const mb = (stoneB.vy - rvy) / (stoneB.vx - rvx);

            const ca = stoneA.py - ma * stoneA.px;
            const cb = stoneB.py - mb * stoneB.px;

            const rpx = Number.parseInt(((cb - ca) / (ma - mb)).toString(), 10);
            const rpy = Number.parseInt((ma * rpx + ca).toString(), 10);

            const time = Math.round((rpx - stoneA.px) / (stoneA.vx - rvx));
            const rpz = stoneA.pz + (stoneA.vz - rvz) * time;

            const result = rpx + rpy + rpz;
            if (!results[result]) {
                results[result] = 1;
            } else {
                results[result]++;
            }
        }
    }

    const result = Object.keys(results).sort((a, b) => results[Number(b)] - results[Number(a)])[0];

    console.log(result);
}

coordinatesOfInitialPosition();
