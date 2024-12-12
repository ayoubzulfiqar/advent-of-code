from collections import defaultdict

import numpy as np
from more_itertools import split_when

with open("./input.txt") as f:
    grid = np.array([list(line) for line in f.read().splitlines()])

shapeX, shapeY = np.shape(grid)

neighborDict = defaultdict(lambda: set())
edge_dict = defaultdict(lambda: set())

for i in range(shapeX):
    for j in range(shapeY):
        for k, (di, dj) in enumerate(
            [
                [-1, 0],
                [
                    0,
                    -1,
                ],
                [0, 1],
                [1, 0],
            ]
        ):
            if (
                i + di in range(shapeX)
                and j + dj in range(shapeY)
                and grid[i][j] == grid[i + di][j + dj]
            ):
                neighborDict[(i, j)].add((i + di, j + dj))
            else:
                edge_dict[(i, j)].add(
                    (i + max(di, 0), j + max(dj, 0), k, "vh"[abs(di)])
                )


def getRegion(point):
    region = set()
    remaining = {point}
    while remaining:
        cur_point = remaining.pop()
        region.add(cur_point)
        remaining |= neighborDict[cur_point] - region
    return region


regions = []
remaining_points = {(i, j) for i in range(shapeX) for j in range(shapeY)}
while remaining_points:
    region = getRegion(remaining_points.pop())
    regions.append(region)
    remaining_points = set(remaining_points) - region


def numSides(region):
    edges = set()
    for point in region:
        edges |= edge_dict[point]
    horizontals = sorted([line for line in edges if line[-1] == "h"])
    verticals = sorted(
        [line for line in edges if line[-1] == "v"], key=lambda x: [x[1], x[0]]
    )
    horizontals = list(
        split_when(
            horizontals,
            lambda x, y: not (x[0] == y[0] and y[1] - x[1] == 1 and y[2] == x[2]),
        )
    )
    verticals = list(
        split_when(
            verticals,
            lambda x, y: not (x[1] == y[1] and y[0] - x[0] == 1 and y[2] == x[2]),
        )
    )
    return len(horizontals) + len(verticals)


def area(region):
    return len(region)


answer = sum(numSides(region) * area(region) for region in regions)
print(answer)
