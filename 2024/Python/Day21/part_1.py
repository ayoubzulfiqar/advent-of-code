from itertools import permutations

import numpy as np

content = []

# Read input lines and filter lines matching the pattern
with open("input.txt", "r") as f:
    for line in f:
        if "A" in line and any(char.isdigit() for char in line[:3]):
            content.append(line.strip())

# Positions and directions as numpy vectors
POSITIONS = {
    "7": np.array([0, 0]),
    "8": np.array([0, 1]),
    "9": np.array([0, 2]),
    "4": np.array([1, 0]),
    "5": np.array([1, 1]),
    "6": np.array([1, 2]),
    "1": np.array([2, 0]),
    "2": np.array([2, 1]),
    "3": np.array([2, 2]),
    "0": np.array([3, 1]),
    "A": np.array([3, 2]),
    "^": np.array([0, 1]),
    "a": np.array([0, 2]),
    "<": np.array([1, 0]),
    "v": np.array([1, 1]),
    ">": np.array([1, 2]),
}

DIRECTIONS = {
    "^": np.array([-1, 0]),
    "v": np.array([1, 0]),
    "<": np.array([0, -1]),
    ">": np.array([0, 1]),
}


def seeToMoveSet(start, fin, avoid=np.array([0, 0])):
    delta = fin - start
    string = ""
    dx, dy = delta

    # Generate moves
    if dx < 0:
        string += "^" * abs(dx)
    else:
        string += "v" * dx
    if dy < 0:
        string += "<" * abs(dy)
    else:
        string += ">" * dy

    # Generate unique permutations of moves
    rv = [
        "".join(p) + "a"
        for p in set(permutations(string))
        if not any(
            (sum(DIRECTIONS[move] for move in p[:i]) + start == avoid).all()
            for i in range(len(p))
        )
    ]

    return rv if rv else ["a"]


memoization = {}


def minLength(s, lim=2, depth=0):
    key = (s, depth, lim)
    if key in memoization:
        return memoization[key]

    avoid = np.array([3, 0]) if depth == 0 else np.array([0, 0])
    cur = POSITIONS["A"] if depth == 0 else POSITIONS["a"]
    length = 0

    for char in s:
        nextCurrent = POSITIONS[char]
        moveSet = seeToMoveSet(cur, nextCurrent, avoid)
        if depth == lim:
            length += len(moveSet[0])
        else:
            length += min(minLength(move, lim, depth + 1) for move in moveSet)
        cur = nextCurrent

    memoization[key] = length
    return length


def sumOfFiveCodeComplexities() -> int:
    complexityA = 0
    for code in content:
        lengthA = minLength(code)
        numeric = int(code[:3])
        complexityA += lengthA * numeric
    return complexityA
