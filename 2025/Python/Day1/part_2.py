import sys
from typing import List, Protocol, runtime_checkable


@runtime_checkable
class TurnDirection(Protocol):
    def get_amount(self) -> int: ...
    def with_amount(self, amount: int) -> "TurnDirection": ...


class LeftTurn:
    def __init__(self, amount: int):
        self.amount = amount

    def get_amount(self) -> int:
        return self.amount

    def with_amount(self, amount: int) -> "TurnDirection":
        return LeftTurn(amount)

    def __repr__(self):
        return f"LeftTurn({self.amount})"


class RightTurn:
    def __init__(self, amount: int):
        self.amount = amount

    def get_amount(self) -> int:
        return self.amount

    def with_amount(self, amount: int) -> "TurnDirection":
        return RightTurn(amount)

    def __repr__(self):
        return f"RightTurn({self.amount})"


def parse_direction(s: str) -> TurnDirection:
    if not s:
        raise ValueError("empty string")

    try:
        amount = int(s[1:])
    except ValueError:
        raise ValueError(f"invalid number: {s[1:]}")

    if s[0] == "L":
        return LeftTurn(amount)
    elif s[0] == "R":
        return RightTurn(amount)
    else:
        raise ValueError("malformed input: expected L<n> or R<n>")


def apply_turn(
    pointer: int, zero_count: int, direction: TurnDirection
) -> tuple[int, int]:
    # Iterative version to avoid recursion depth issues
    p = pointer
    z = zero_count
    d = direction

    while d.get_amount() > 0:
        if isinstance(d, LeftTurn):
            if p == 0:
                p = (p - 1) % 100
                if p < 0:
                    p += 100
                z += 1
            else:
                p = (p - 1) % 100
                if p < 0:
                    p += 100
            d = d.with_amount(d.amount - 1)
        elif isinstance(d, RightTurn):
            if p == 0:
                p = (p + 1) % 100
                z += 1
            else:
                p = (p + 1) % 100
            d = d.with_amount(d.amount - 1)
        else:
            raise TypeError(f"Unknown direction type: {type(d)}")

    return p, z


def compute_result(input_str: str) -> int:
    lines = input_str.strip().split("\n")
    directions: List[TurnDirection] = []

    for line in lines:
        line = line.strip()
        if not line:
            continue

        direction = parse_direction(line)
        directions.append(direction)

    pointer = 50
    zero_count = 0

    for direction in directions:
        pointer, zero_count = apply_turn(pointer, zero_count, direction)

    return zero_count


def Method0x434C49434BToOpenTheDoor():
    try:
        with open("input.txt", "r") as file:
            data = file.read()
    except FileNotFoundError:
        print("error: input.txt file not found", file=sys.stderr)
        sys.exit(1)
    except IOError as e:
        print(f"error reading file: {e}", file=sys.stderr)
        sys.exit(1)

    try:
        output = compute_result(data)
        print(f"result = {output}")
    except ValueError as e:
        print(f"error processing input: {e}", file=sys.stderr)
        sys.exit(1)


if __name__ == "__main__":
    Method0x434C49434BToOpenTheDoor()
