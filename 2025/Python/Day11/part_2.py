from collections import deque
from functools import lru_cache
from typing import Dict, List, Set


class UltraOptimizedSolver:
    def __init__(self) -> None:
        self.adjacencyLIST: Dict[str, List[str]] = {}
        self.nodes_reachingFFT: Set[str] = set()
        self.nodes_reachingDAC: Set[str] = set()

    def loadGraph(self, filename: str) -> None:
        adjacencyLIST: Dict[str, List[str]] = {}

        with open(filename, "r", encoding="utf-8") as file:
            for line in file:
                parts = line.strip().split()
                if len(parts) < 2:
                    continue

                parent_node = parts[0].rstrip(":")
                child_nodes = parts[1:]
                adjacencyLIST[parent_node] = child_nodes

        self.adjacencyLIST = adjacencyLIST

        self._precomputeReachabilityToSpecialNodes()

    def _precomputeReachabilityToSpecialNodes(self) -> None:
        def findNodesReachingTarget(target_node: str) -> Set[str]:
            reachable_nodes: Set[str] = set()

            reverse_adjacency: Dict[str, List[str]] = {}
            for parent in self.adjacencyLIST:
                for child in self.adjacencyLIST[parent]:
                    if child not in reverse_adjacency:
                        reverse_adjacency[child] = []
                    reverse_adjacency[child].append(parent)

            queue = deque([target_node])
            visited_nodes = {target_node}

            while queue:
                current_node = queue.popleft()
                reachable_nodes.add(current_node)

                for predecessor in reverse_adjacency.get(current_node, []):
                    if predecessor not in visited_nodes:
                        visited_nodes.add(predecessor)
                        queue.append(predecessor)

            return reachable_nodes

        self.nodes_reachingFFT = findNodesReachingTarget("fft")
        self.nodes_reachingDAC = findNodesReachingTarget("dac")

    @lru_cache(maxsize=None)
    def _countPathsFromNode(
        self, current_node: str, visited_special_mask: int
    ) -> int:
        if current_node == "out":
            return 1 if visited_special_mask == 0b11 else 0

        if (
            not (visited_special_mask & 0b01)
            and current_node not in self.nodes_reachingFFT
        ):
            return 0

        if (
            not (visited_special_mask & 0b10)
            and current_node not in self.nodes_reachingDAC
        ):
            return 0

        total_paths = 0

        for neighbor_node in self.adjacencyLIST.get(current_node, []):
            new_mask = visited_special_mask

            if neighbor_node == "fft":
                new_mask |= 0b01
            elif neighbor_node == "dac":
                new_mask |= 0b10

            if not (new_mask & 0b01) and neighbor_node not in self.nodes_reachingFFT:
                continue
            if not (new_mask & 0b10) and neighbor_node not in self.nodes_reachingDAC:
                continue

            total_paths += self._countPathsFromNode(neighbor_node, new_mask)

        return total_paths

    def countPathsFromSvr(self) -> int:
        if "svr" not in self.adjacencyLIST:
            return 0

        return self._countPathsFromNode("svr", 0)


def countPathsWithBothSpecialNodes() -> None:
    path_counter = UltraOptimizedSolver()

    try:
        path_counter.loadGraph("input.txt")
    except FileNotFoundError as file_error:
        print(f"Error: Could not open file 'input.txt': {file_error}")
        return

    path_count = path_counter.countPathsFromSvr()
    print(f"Result: {path_count}")


if __name__ == "__main__":
    countPathsWithBothSpecialNodes()
