import multiprocessing as mp
import sys
from typing import List, Tuple


class Puzzle:
    def __init__(self, target: str, joltage: List[int], buttons: List[List[int]]):
        self.target = target
        self.joltage = joltage
        self.buttons = buttons


def parse_input(filename: str) -> List[Puzzle]:
    puzzles = []

    try:
        with open(filename, "r", encoding="utf-8") as file:
            for line in file:
                line = line.strip()
                if not line:
                    continue

                parts = line.split()
                if len(parts) < 2:
                    continue

                # Parse target configuration
                target = ""
                if len(parts[0]) >= 2 and parts[0][0] == "[" and parts[0][-1] == "]":
                    target = parts[0][1:-1]
                else:
                    continue

                # Parse joltage values
                joltage_str = parts[-1]
                joltage = []
                if (
                    len(joltage_str) >= 2
                    and joltage_str[0] == "{"
                    and joltage_str[-1] == "}"
                ):
                    values = joltage_str[1:-1].split(",")
                    for v in values:
                        v = v.strip()
                        if v:
                            try:
                                num = int(v)
                                joltage.append(num)
                            except ValueError:
                                pass
                else:
                    continue

                # Parse buttons
                buttons = []
                for i in range(1, len(parts) - 1):
                    btn = parts[i]
                    if len(btn) >= 2 and btn[0] == "(" and btn[-1] == ")":
                        positions = btn[1:-1].split(",")
                        button = []
                        for pos in positions:
                            pos = pos.strip()
                            if pos:
                                try:
                                    num = int(pos)
                                    button.append(num)
                                except ValueError:
                                    pass
                        buttons.append(button)

                puzzles.append(Puzzle(target, joltage, buttons))

    except FileNotFoundError:
        print(f"Error: File '{filename}' not found")
        sys.exit(1)

    return puzzles


def gaussian_elimination(matrix: List[List[int]]) -> Tuple[List[int], List[List[int]]]:
    if not matrix:
        return [], []

    m = len(matrix)
    n = len(matrix[0]) - 1

    pivot_cols = []
    current_row = 0

    # Make a deep copy
    mat = [row[:] for row in matrix]

    for col in range(n):
        if current_row >= m:
            break

        # Find pivot row
        pivot_row = -1
        for row in range(current_row, m):
            if mat[row][col] != 0:
                pivot_row = row
                break

        if pivot_row == -1:
            continue

        # Swap rows
        mat[current_row], mat[pivot_row] = mat[pivot_row], mat[current_row]
        pivot_cols.append(col)

        # Eliminate below
        for row in range(current_row + 1, m):
            if mat[row][col] != 0:
                factor = mat[row][col]
                pivot_val = mat[current_row][col]

                for j in range(col, n + 1):
                    mat[row][j] = mat[row][j] * pivot_val - mat[current_row][j] * factor

        current_row += 1

    return pivot_cols, mat


def solve_system_exact(buttons: List[List[int]], joltages: List[int]) -> int:
    """Exact translation of Go solveSystem function"""
    n = len(buttons)
    m = len(joltages)

    # Create augmented matrix
    matrix = [[0] * (n + 1) for _ in range(m)]
    for i in range(m):
        for j in range(n):
            # Check if button j affects position i
            affects = False
            for pos in buttons[j]:
                if pos == i:
                    affects = True
                    break
            if affects:
                matrix[i][j] = 1
        matrix[i][n] = joltages[i]

    # Perform Gaussian elimination
    pivot_cols, reduced_matrix = gaussian_elimination(matrix)

    # Identify free variables
    pivot_set = set(pivot_cols)
    free_vars = [i for i in range(n) if i not in pivot_set]

    best_solution = [0] * n
    best_sum = -1

    def try_solution(free_values: List[int]) -> None:
        nonlocal best_solution, best_sum

        solution = [0] * n
        for i, var_idx in enumerate(free_vars):
            solution[var_idx] = free_values[i] if i < len(free_values) else 0

        # Back-substitute
        for idx in range(len(pivot_cols) - 1, -1, -1):
            row = idx
            col = pivot_cols[idx]
            total = reduced_matrix[row][n]

            for j in range(col + 1, n):
                total -= reduced_matrix[row][j] * solution[j]

            if reduced_matrix[row][col] == 0:
                return

            if total % reduced_matrix[row][col] != 0:
                return

            val = total // reduced_matrix[row][col]
            if val < 0:
                return

            solution[col] = val

        # Verify solution
        for i in range(m):
            total = 0
            for j in range(n):
                if solution[j] > 0:
                    for pos in buttons[j]:
                        if pos == i:
                            total += solution[j]
                            break
            if total != joltages[i]:
                return

        # Calculate total presses
        total_presses = sum(solution)

        if best_sum == -1 or total_presses < best_sum:
            best_solution = solution[:]
            best_sum = total_presses

    # EXACT search bounds from Go code
    if len(free_vars) == 0:
        try_solution([])
    elif len(free_vars) == 1:
        max_val = 0
        for j in joltages:
            if j > max_val:
                max_val = j
        max_val *= 3
        for val in range(max_val + 1):
            if best_sum != -1 and val > best_sum:
                break
            try_solution([val])
    elif len(free_vars) == 2:
        max_val = 0
        for j in joltages:
            if j > max_val:
                max_val = j
        if max_val < 200:
            max_val = 200
        for v1 in range(max_val + 1):
            for v2 in range(max_val + 1):
                if best_sum != -1 and v1 + v2 > best_sum:
                    continue
                try_solution([v1, v2])
    elif len(free_vars) == 3:
        for v1 in range(250):
            for v2 in range(250):
                for v3 in range(250):
                    if best_sum != -1 and v1 + v2 + v3 > best_sum:
                        continue
                    try_solution([v1, v2, v3])
    elif len(free_vars) == 4:
        for v1 in range(30):
            for v2 in range(30):
                for v3 in range(30):
                    for v4 in range(30):
                        if best_sum != -1 and v1 + v2 + v3 + v4 > best_sum:
                            continue
                        try_solution([v1, v2, v3, v4])
    else:
        # Fallback
        try_solution([0] * len(free_vars))

    return 0 if best_sum == -1 else best_sum


def process_puzzle(puzzle: Puzzle) -> int:
    return solve_system_exact(puzzle.buttons, puzzle.joltage)


def configure_jolt_level(puzzles: List[Puzzle]) -> int:
    if not puzzles:
        return 0

    # Use all CPU cores
    num_workers = min(mp.cpu_count(), len(puzzles))

    with mp.Pool(processes=num_workers) as pool:
        results = pool.map(process_puzzle, puzzles)

    return sum(results)


def main():
    puzzles = parse_input("input.txt")
    if not puzzles:
        print("No puzzles to solve")
        return

    result2 = configure_jolt_level(puzzles)
    print("Part-2:", result2)


if __name__ == "__main__":
    mp.freeze_support()
    main()
