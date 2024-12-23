def combination(array, k):
    result: list = []

    def helper(_array, _k, _i, _current):
        if len(_current) == k:
            result.append(_current)
        if len(_current) == k or _i == len(_array):
            return

        helper(_array, _k, _i + 1, [_array[_i]] + _current)
        helper(_array, _k, _i + 1, _current)

    helper(array, k, 0, [])
    return result


def findSet(graph, path, size):
    current = path[-1]
    if len(path) == size + 1:
        if current == path[0]:
            return [path[:size]]
        else:
            return []

    if len(set(path)) != len(path):
        return []

    allSets = []
    for neighbor in graph[current]:
        path.append(neighbor)
        sets = findSet(graph, path, size)
        path.pop()

        if sets:
            allSets.extend(sets)

    return allSets


# Part_1
def computerNameT() -> int:
    with open("input.txt") as file:
        con = file.read()
    lines = con.strip().split("\n")
    graph: map = {}
    for line in lines:
        left, right = line.split("-")
        graph.setdefault(left, []).append(right)
        graph.setdefault(right, []).append(left)

    allSets = set()
    for node in graph.keys():
        sets = findSet(graph, [node], 3)
        allSets.update(",".join(sorted(s)) for s in sets)
        result: int = sum(
            1 for s in allSets if any(node.startswith("t") for node in s.split(","))
        )
    print(result)
    return result
