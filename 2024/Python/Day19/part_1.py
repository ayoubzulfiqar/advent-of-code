import functools
from typing import List, Tuple


def loadInput(filePath: str) -> Tuple[List[str], List[str]]:
    """
    Load the input data from a file.
    Args:
        file_path (str): Path to the input file.
    Returns:
        Tuple[List[str], List[str]]: Available substrings and target strings.
    """
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


def countPossibleDesigns(file_path: str) -> int:
    """
    Counts how many target strings can be constructed using available substrings.
    Args:
        file_path (str): Path to the input file.
    Returns:
        int: Number of target strings that can be constructed.
    """
    available_substrings, target_strings = loadInput(file_path)

    @functools.lru_cache(None)
    def is_possible(target: str) -> bool:
        if not target:
            return True
        for start in available_substrings:
            if target.startswith(start):
                if is_possible(target[len(start) :]):
                    return True
        return False

    result = sum(is_possible(target) for target in target_strings)
    print(result)
    return result


# if __name__ == "__main__":
#     countPossibleDesigns("input.txt")
