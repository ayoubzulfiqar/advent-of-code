import time


def fork_list_rolls(input_str: str) -> tuple[str, float]:
    start_time = time.time()

    lines = [list(line) for line in input_str.strip().split("\n")]

    height = len(lines)
    accessible_count = 0

    for row in range(height):
        current_line = lines[row]
        width = len(current_line)

        for col in range(width):
            if current_line[col] != "@":
                continue

            adjacent_count = 0

            for dr, dc in [
                (-1, -1),
                (-1, 0),
                (-1, 1),
                (0, -1),
                (0, 1),
                (1, -1),
                (1, 0),
                (1, 1),
            ]:
                neighbor_row = row + dr
                neighbor_col = col + dc

                if 0 <= neighbor_row < height and 0 <= neighbor_col < len(
                    lines[neighbor_row]
                ):
                    if lines[neighbor_row][neighbor_col] == "@":
                        adjacent_count += 1

            if adjacent_count < 4:
                accessible_count += 1

    return str(accessible_count), time.time() - start_time


def main():
    try:
        with open("input.txt", "r", encoding="utf-8") as file:
            input_str = file.read()
    except FileNotFoundError:
        print("Error: input.txt not found")
        return

    part1_result, part1_time = fork_list_rolls(input_str)
    print(f"Part 1 Result: {part1_result} (Time: {part1_time:.6f}s)")


if __name__ == "__main__":
    main()
