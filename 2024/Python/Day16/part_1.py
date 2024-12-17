import heapq
from typing import Dict, List, Tuple


class MinHeap:
    """
    Min heap implementation using heapq library
    """

    def __init__(self):
        self.heap = []

    def insert(self, element: Tuple[int, str]):
        heapq.heappush(self.heap, element)

    def extract_min(self) -> Tuple[int, str]:
        return heapq.heappop(self.heap)

    def size(self) -> int:
        return len(self.heap)


DIRECTIONS = [
    {"x": 1, "y": 0},
    {"x": 0, "y": 1},
    {"x": -1, "y": 0},
    {"x": 0, "y": -1},
]


def dijkstra(
    graph: Dict[str, Dict[str, int]], start: Dict[str, int], directionless: bool
) -> Dict[str, int]:
    queue = MinHeap()
    distances = {}

    starting_key = (
        f"{start['x']},{start['y']},0"
        if not directionless
        else f"{start['x']},{start['y']}"
    )
    queue.insert((0, starting_key))
    distances[starting_key] = 0

    while queue.size() > 0:
        current_score, current_node = queue.extract_min()

        if distances[current_node] < current_score:
            continue

        if current_node not in graph:
            continue

        for next_node, weight in graph[current_node].items():
            new_score = current_score + weight
            if next_node not in distances or distances[next_node] > new_score:
                distances[next_node] = new_score
                queue.insert((new_score, next_node))

    return distances


def parse_grid(grid: List[str]):
    width, height = len(grid[0]), len(grid)

    start = {"x": 0, "y": 0}
    end = {"x": 0, "y": 0}
    forward = {}
    reverse = {}

    for y in range(height):
        for x in range(width):
            if grid[y][x] == "S":
                start = {"x": x, "y": y}
            if grid[y][x] == "E":
                end = {"x": x, "y": y}

            if grid[y][x] != "#":
                for i, direction in enumerate(DIRECTIONS):
                    position = {"x": x + direction["x"], "y": y + direction["y"]}

                    key = f"{x},{y},{i}"
                    move_key = f"{position['x']},{position['y']},{i}"

                    if (
                        0 <= position["x"] < width
                        and 0 <= position["y"] < height
                        and grid[position["y"]][position["x"]] != "#"
                    ):
                        forward.setdefault(key, {})[move_key] = 1
                        reverse.setdefault(move_key, {})[key] = 1

                    for rotate_key in [
                        f"{x},{y},{(i + 3) % 4}",
                        f"{x},{y},{(i + 1) % 4}",
                    ]:
                        forward.setdefault(key, {})[rotate_key] = 1000
                        reverse.setdefault(rotate_key, {})[key] = 1000

    for i in range(len(DIRECTIONS)):
        key = f"{end['x']},{end['y']}"
        rotate_key = f"{end['x']},{end['y']},{i}"

        forward.setdefault(rotate_key, {})[key] = 0
        reverse.setdefault(key, {})[rotate_key] = 0

    return {"start": start, "end": end, "forward": forward, "reverse": reverse}


def lowestScoreOfReindeer() -> int:
    with open("./input.txt") as file:
        con = file.read()
    grid = con.strip().split("\n")
    parsed = parse_grid(grid)
    distances = dijkstra(parsed["forward"], parsed["start"], False)
    return distances.get(f"{parsed['end']['x']},{parsed['end']['y']}", float("inf"))
