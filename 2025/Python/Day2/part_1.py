import math
from dataclasses import dataclass
from typing import List


@dataclass
class IDRange:
    start_id: int
    end_id: int


def parse_ranges(input_str: str) -> List[IDRange]:
    ranges = []
    line_pairs = input_str.strip().split(",")

    for pair in line_pairs:
        ids = pair.split("-")
        if len(ids) != 2:
            continue

        try:
            start_id = int(ids[0].strip())
            end_id = int(ids[1].strip())
            ranges.append(IDRange(start_id=start_id, end_id=end_id))
        except ValueError:
            continue

    return ranges


def count_digits(num: int) -> int:
    if num == 0:
        return 1
    return int(math.log10(num)) + 1


def is_divisible_by(dividend: int, divisor: int) -> bool:
    return dividend % divisor == 0


def find_invalid_ids(start_id: int, end_id: int) -> List[int]:
    invalid_ids = []

    start_digits = count_digits(start_id)
    end_digits = count_digits(end_id)

    # Determine minimum half digits
    if is_divisible_by(start_digits, 2):
        min_half_digits = max(start_digits // 2, 1)
    else:
        min_half_digits = start_digits // 2 + 1

    max_half_digits = end_digits // 2

    for half_length in range(min_half_digits, max_half_digits + 1):
        first_half_lower_bound = 10 ** (half_length - 1)
        if first_half_lower_bound < start_id // (10**half_length):
            first_half_lower_bound = start_id // (10**half_length)

        first_half_upper_bound = 10**half_length - 1
        if first_half_upper_bound > end_id // (10**half_length):
            first_half_upper_bound = end_id // (10**half_length)

        if first_half_lower_bound > first_half_upper_bound:
            continue

        # Generate invalid IDs
        for first_half in range(first_half_lower_bound, first_half_upper_bound + 1):
            repeating_id = first_half * (10**half_length) + first_half

            if start_id <= repeating_id <= end_id:
                invalid_ids.append(repeating_id)

    return invalid_ids


def calculate_invalid_id_sum(id_ranges: List[IDRange]) -> int:
    total_sum = 0

    for id_range in id_ranges:
        invalid_ids = find_invalid_ids(id_range.start_id, id_range.end_id)
        for invalid_id in invalid_ids:
            total_sum += invalid_id

    return total_sum


def solve_part1() -> None:
    try:
        with open("input.txt", "r") as file:
            data = file.read()
    except FileNotFoundError:
        print("Error: input.txt not found")
        return
    except Exception as e:
        print(f"Error reading file: {e}")
        return

    id_ranges = parse_ranges(data)
    solution = calculate_invalid_id_sum(id_ranges)
    print(f"Part 1 Result: {solution}")


if __name__ == "__main__":
    solve_part1()
