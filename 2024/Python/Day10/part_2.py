from collections import deque
from typing import Dict, List


def findTrails(grid: List[List[int]], starting: Dict[str, int], part1: bool) -> int:
    width, height = len(grid[0]), len(grid)
    queue = deque([starting.copy()])
    visited = set()  # Only used in part 1

    paths = 0
    while queue:
        current = queue.popleft()

        if grid[current["y"]][current["x"]] == 9:
            paths += 1
            continue

        # Try all directions
        for direction in [
            {"x": 0, "y": -1},
            {"x": 1, "y": 0},
            {"x": 0, "y": 1},
            {"x": -1, "y": 0},
        ]:
            position = {
                "x": current["x"] + direction["x"],
                "y": current["y"] + direction["y"],
            }
            if (
                position["x"] < 0
                or position["x"] >= width
                or position["y"] < 0
                or position["y"] >= height
                or (f"{position['x']},{position['y']}" in visited and part1)
                or grid[position["y"]][position["x"]] - grid[current["y"]][current["x"]]
                != 1
            ):
                continue

            queue.append(position)
            if part1:
                visited.add(f"{position['x']},{position['y']}")

    return paths


def ratingOfAllTrailHeads() -> int:
    with open("./input.txt") as file:
        con = file.read()
    grid = [[int(num) for num in line] for line in con.strip().split("\n")]
    width, height = len(grid[0]), len(grid)

    # Add all trailHead ratings
    total = 0
    for y in range(height):
        for x in range(width):
            if grid[y][x] == 0:
                total += findTrails(grid, {"x": x, "y": y}, False)
    print(total)
    return total
