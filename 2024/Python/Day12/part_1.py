from collections import defaultdict

import numpy as np

with open("./input.txt") as f:
    grid = np.array([list(line) for line in f.read().splitlines()])

shapeX, shapeY = np.shape(grid)

neighborDict = defaultdict(lambda: set())

for i in range(shapeX):
    for j in range(shapeY):
        for di, dj in [
            [-1, 0],
            [
                0,
                -1,
            ],
            [0, 1],
            [1, 0],
        ]:
            if i + di in range(shapeX) and j + dj in range(shapeY):
                if grid[i][j] == grid[i + di][j + dj]:
                    neighborDict[(i, j)].add((i + di, j + dj))



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


def perimeter(region):
    return sum(4 - len(neighborDict[point]) for point in region)


def area(region):
    return len(region)

answer = sum(perimeter(region) * area(region) for region in regions)
print(answer)
