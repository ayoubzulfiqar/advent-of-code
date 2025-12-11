from collections import deque
from typing import Dict, List


def countPaths(adj: Dict[str, List[str]], start: str) -> int:
    all_nodes = set(adj.keys())
    for node in adj:
        all_nodes.update(adj[node])

    in_degree = {node: 0 for node in all_nodes}
    for node in adj:
        for child in adj[node]:
            in_degree[child] += 1

    dp = {node: 0 for node in all_nodes}
    dp["out"] = 1

    queue = deque([node for node in in_degree if in_degree[node] == 0])
    topo_order = []

    while queue:
        node = queue.popleft()
        topo_order.append(node)
        for child in adj.get(node, []):
            in_degree[child] -= 1
            if in_degree[child] == 0:
                queue.append(child)

    if len(topo_order) != len(all_nodes):
        print("Warning: Graph has cycles, DP approach may not work correctly")
        return -1

    for node in reversed(topo_order):
        if node == "out":
            continue
        total = 0
        for child in adj.get(node, []):
            total += dp[child]
        dp[node] = total

    return dp.get(start, 0)


def youPathOut() -> None:
    adj = {}

    try:
        with open("input.txt", "r", encoding="utf-8") as f:
            for line in f:
                parts = line.strip().split()
                if len(parts) < 2:
                    continue
                key = parts[0].rstrip(":")
                adj[key] = parts[1:]
    except FileNotFoundError as fileError:
        print(f"Error: File 'input.txt' not found:\n{fileError}")
        return

    result = countPaths(adj, "you")
    if result is not None:
        print(result)


if __name__ == "__main__":
    youPathOut()
