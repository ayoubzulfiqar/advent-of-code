import sys
from collections import Counter, defaultdict
from typing import List

sys.setrecursionlimit(100000)
FILE = sys.argv[1] if len(sys.argv) > 1 else "./input.txt"


def readLinesToList() -> List[int]:
    lines: List[str] = []
    with open(FILE, "r", encoding="utf-8") as f:
        for line in f:
            line = line.strip()
            lines.append([int(val) for val in line.split()])

    return lines


def blink(stone: int) -> List[int]:
    if stone == 0:
        return [1]

    s = str(stone)
    if len(s) % 2 == 0:
        return [int(s[: (len(s) // 2)]), int(s[(len(s) // 2) :])]

    return [stone * 2024]


def blinkTwentyFiveTimes() -> int:
    lines = readLinesToList()
    answer = 0

    stones = Counter(lines[0])
    for _ in range(25):
        new_stones = defaultdict(int)
        for rock, count in stones.items():
            blink_result = blink(rock)
            for blink_result_rock in blink_result:
                new_stones[blink_result_rock] += count

        stones = Counter(new_stones)

    answer = sum(stones.values())
    print(answer)
    return answer


