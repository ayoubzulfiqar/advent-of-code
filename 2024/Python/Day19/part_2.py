import functools
from typing import List, Tuple


def loadInput(filePath: str) -> Tuple[List[str], List[str]]:
    try:
        with open(filePath, "r") as f:
            available, targets = f.read().strip().split("\n\n")
        available_substrings = available.split(", ")
        target_strings = targets.splitlines()
        return available_substrings, target_strings
    except FileNotFoundError:
        raise FileNotFoundError(f"Input file {filePath} not found.")
    except ValueError:
        raise ValueError(
            "Input file format is incorrect. Ensure it contains two sections separated by a blank line."
        )


def differentWaysToMakeDesign(filePath: str) -> int:
    availableSubstrings, targetStrings = loadInput(filePath)

    @functools.lru_cache(None)
    def numWays(target: str) -> int:
        if not target:
            return 1
        total_ways = 0
        for start in availableSubstrings:
            if target.startswith(start):
                total_ways += numWays(target[len(start) :])
        return total_ways

    result = sum(numWays(target) for target in targetStrings)
    print(result)
    return result


# if __name__ == "__main__":
#     differentWaysToMakeDesign("input.txt")
