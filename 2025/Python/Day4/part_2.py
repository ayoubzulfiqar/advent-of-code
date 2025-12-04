import time


def elevs_and_fork_lifts(input_str: str) -> tuple[str, float]:
    start_time = time.time()

    grid = [list(line) for line in input_str.strip().split("\n")]

    height = len(grid)
    total_removed = 0

    directions = [(-1, -1), (-1, 0), (-1, 1), (0, -1), (0, 1), (1, -1), (1, 0), (1, 1)]

    while True:
        cells_to_remove = []

        for row in range(height):
            current_row = grid[row]
            width = len(current_row)

            for col in range(width):
                if current_row[col] != "@":
                    continue

                adjacent_count = 0

                for dr, dc in directions:
                    neighbor_row = row + dr
                    neighbor_col = col + dc

                    if 0 <= neighbor_row < height and 0 <= neighbor_col < len(
                        grid[neighbor_row]
                    ):
                        if grid[neighbor_row][neighbor_col] == "@":
                            adjacent_count += 1

                if adjacent_count < 4:
                    cells_to_remove.append((row, col))

        if not cells_to_remove:
            break

        for row, col in cells_to_remove:
            grid[row][col] = "."

        total_removed += len(cells_to_remove)

    return str(total_removed), time.time() - start_time


def main():
    try:
        with open("input.txt", "r", encoding="utf-8") as file:
            input_str = file.read()
    except FileNotFoundError:
        print("Error: input.txt not found")
        return

    part2_result, part2_time = elevs_and_fork_lifts(input_str)
    print(f"Part 2 Result: {part2_result} (Time: {part2_time:.6f}s)")


if __name__ == "__main__":
    main()
