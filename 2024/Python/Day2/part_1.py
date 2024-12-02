# Read and process the input file
def readInput(filename: str) -> list[list[int]]:
    with open(filename, "r", encoding="utf-8") as file:
        # Read lines, split them, and convert to integers
        data = [list(map(int, line.split())) for line in file if line.strip()]
    return data


# Check if a line is "safe" based on its pattern
def isSafe(line: list[int]) -> bool:
    # Determine direction based on the first two numbers
    if len(line) < 2:
        return False  # A single number cannot form a pattern

    if line[1] < line[0]:  # Decreasing
        allowed_diffs = {-1, -2, -3}
    elif line[1] > line[0]:  # Increasing
        allowed_diffs = {1, 2, 3}
    else:
        return False  # First two numbers are the same, not safe

    # Check if all consecutive differences are allowed
    for i in range(1, len(line)):
        if (line[i] - line[i - 1]) not in allowed_diffs:
            return False

    return True


def safeReports() -> int:
    filename = "D:/Projects/advent-of-code/2024/Python/Day2/input.txt"
    data = readInput(filename)
    total_safe = sum(1 for line in data if isSafe(line))
    print("Part 1:", total_safe)
    return total_safe
