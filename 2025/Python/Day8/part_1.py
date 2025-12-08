from collections import defaultdict

MAX_DIST = 188_000_000


def distance(a, b):
    return abs(a[0] - b[0]) ** 2 + abs(a[1] - b[1]) ** 2 + abs(a[2] - b[2]) ** 2


class DisjointSet:
    def __init__(self, n):
        self.parent = list(range(n))
        self.rank = [0] * n

    def find(self, x):
        if self.parent[x] != x:
            self.parent[x] = self.find(self.parent[x])
        return self.parent[x]

    def union(self, x, y):
        root_x = self.find(x)
        root_y = self.find(y)

        if root_x == root_y:
            return

        if self.rank[root_x] < self.rank[root_y]:
            root_x, root_y = root_y, root_x

        self.parent[root_y] = root_x
        if self.rank[root_x] == self.rank[root_y]:
            self.rank[root_x] += 1


def multiplicationThreeLargestCircuits():
    coords = []
    with open("input.txt", "r", encoding="utf-8") as f:
        for line in f:
            line = line.strip()
            if not line:
                continue
            parts = line.split(",")
            if len(parts) != 3:
                continue
            coord = tuple(map(int, parts))
            coords.append(coord)

    edges = []
    n = len(coords)

    for i in range(n - 1):
        for j in range(i + 1, n):
            dist_val = distance(coords[i], coords[j])
            if dist_val < MAX_DIST:
                edges.append((i, j, dist_val))

    edges.sort(key=lambda x: x[2])

    dsu = DisjointSet(n)
    for i in range(min(1000, len(edges))):
        a, b, _ = edges[i]
        dsu.union(a, b)

    circuit_counts = defaultdict(int)
    for i in range(n):
        root = dsu.find(i)
        circuit_counts[root] += 1

    circuits = list(circuit_counts.values())
    circuits.sort(reverse=True)

    result = 1
    for i in range(min(3, len(circuits))):
        result *= circuits[i]

    print(result)


if __name__ == "__main__":
    multiplicationThreeLargestCircuits()
