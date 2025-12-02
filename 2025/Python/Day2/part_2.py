import math
from dataclasses import dataclass
from typing import List, Set


@dataclass
class MultiRepeatRange:
    lower_bound: int
    upper_bound: int


def parse_multi_repeat_ranges(input_data: str) -> List[MultiRepeatRange]:
    """Parse input string into list of multi-repeat ranges."""
    range_collection = []
    pair_strings = input_data.strip().split(",")

    for range_str in pair_strings:
        bounds = range_str.split("-")
        if len(bounds) != 2:
            continue

        try:
            lower_value = int(bounds[0].strip())
            upper_value = int(bounds[1].strip())
            range_collection.append(
                MultiRepeatRange(lower_bound=lower_value, upper_bound=upper_value)
            )
        except ValueError:
            continue

    return range_collection


def digit_count(value: int) -> int:
    """Count the number of digits in a number."""
    if value == 0:
        return 1
    return int(math.log10(value)) + 1


def is_evenly_divisible(dividend: int, divisor: int) -> bool:
    """Check if dividend is evenly divisible by divisor."""
    return dividend % divisor == 0


def discover_multi_repeating_ids(min_id: int, max_id: int) -> Set[int]:
    """Discover IDs that repeat segments multiple times within given bounds."""
    found_ids = set()

    min_digits = digit_count(min_id)
    max_digits = digit_count(max_id)

    for repeat_times in range(2, max_digits + 1):
        # Determine segment length bounds
        if is_evenly_divisible(min_digits, repeat_times):
            segment_min_length = max(min_digits // repeat_times, 1)
        else:
            segment_min_length = min_digits // repeat_times + 1

        segment_max_length = max_digits // repeat_times

        for segment_length in range(segment_min_length, segment_max_length + 1):
            # Calculate bounds for segment values
            min_segment_value = 10 ** (segment_length - 1)
            divisor = 10 ** (segment_length * (repeat_times - 1))

            if min_segment_value < min_id // divisor:
                min_segment_value = min_id // divisor

            max_segment_value = 10**segment_length - 1
            if max_segment_value > max_id // divisor:
                max_segment_value = max_id // divisor

            if min_segment_value > max_segment_value:
                continue

            # Generate IDs for each segment value
            for segment_number in range(min_segment_value, max_segment_value + 1):
                constructed_id = segment_number

                # Repeat the segment
                for _ in range(1, repeat_times):
                    constructed_id = (
                        constructed_id * (10**segment_length) + segment_number
                    )

                if min_id <= constructed_id <= max_id:
                    found_ids.add(constructed_id)

    return found_ids


def compute_multi_repeat_id_sum(ranges: List[MultiRepeatRange]) -> int:
    """Compute sum of all multi-repeat IDs across all ranges."""
    cumulative_sum = 0

    for current_range in ranges:
        multi_repeat_ids = discover_multi_repeating_ids(
            current_range.lower_bound, current_range.upper_bound
        )
        for repeat_id in multi_repeat_ids:
            cumulative_sum += repeat_id

    return cumulative_sum


def solve_part2() -> None:
    """Main function to solve Part 2."""
    try:
        with open("input.txt", "r") as file:
            file_data = file.read()
    except FileNotFoundError:
        print("Error: input.txt not found")
        return
    except Exception as e:
        print(f"Error reading file: {e}")
        return

    range_collection = parse_multi_repeat_ranges(file_data)
    total_result = compute_multi_repeat_id_sum(range_collection)
    print(f"Part 2 Result: {total_result}")


if __name__ == "__main__":
    solve_part2()
