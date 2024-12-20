from collections import deque
from itertools import combinations


def parseInput():
    with open("input.txt", "r") as file:
        content = file.read()
    grid = [list(line) for line in content.strip().split("\n")]
    start, end = (0, 0), (0, 0)
    for r, row in enumerate(grid):
        for c, cell in enumerate(row):
            if cell == "S":
                start = (r, c)
            elif cell == "E":
                end = (r, c)
    return grid, start, end


def bfs(grid, start, end):
    queue = deque([(start[0], start[1], 0)])
    dists = {}
    while queue:
        r, c, n = queue.popleft()
        if (r, c) in dists:
            continue
        dists[(r, c)] = n
        if (r, c) == end:
            continue
        for dr, dc in [(-1, 0), (1, 0), (0, -1), (0, 1)]:
            rr, cc = r + dr, c + dc
            if 0 <= rr < len(grid) and 0 <= cc < len(grid[0]) and grid[rr][cc] != "#":
                queue.append((rr, cc, n + 1))
    return dists


# Part -1
def picoSecondsCheats() -> int:
    grid, start, end = parseInput()
    dists = bfs(grid, start, end)
    p1 = 0
    for ((r1, c1), n1), ((r2, c2), n2) in combinations(dists.items(), 2):
        d = abs(r1 - r2) + abs(c1 - c2)
        if d <= 2 and abs(n2 - n1) >= d + 100:
            p1 += 1
    print(p1)
    return p1
