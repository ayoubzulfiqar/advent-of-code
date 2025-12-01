import sys
from typing import List, Union


class L:
    def __init__(self, n: int):
        self.n = n

    def __repr__(self):
        return f"L({self.n})"


class R:
    def __init__(self, n: int):
        self.n = n

    def __repr__(self):
        return f"R({self.n})"


Direction = Union[L, R]


def parse_turn(s: str) -> Direction:
    if not s:
        raise ValueError("empty string")

    try:
        n = int(s[1:])
    except ValueError:
        raise ValueError(f"invalid number: {s[1:]}")

    if s[0] == "L":
        return L(n)
    elif s[0] == "R":
        return R(n)
    else:
        raise ValueError("malformed input: expected L<n> or R<n>")


def step(dial: int, zeros: int, turn: Direction) -> tuple[int, int]:
    if isinstance(turn, L):
        dial_prime = (dial - turn.n) % 100
    elif isinstance(turn, R):
        dial_prime = (dial + turn.n) % 100
    else:
        raise TypeError(f"Unknown direction type: {type(turn)}")

    zeros_prime = zeros
    if dial_prime == 0:
        zeros_prime += 1

    return dial_prime, zeros_prime


def process(content: str) -> int:
    lines = content.strip().split("\n")
    turns: List[Direction] = []

    for line in lines:
        line = line.strip()
        if not line:
            continue

        turn = parse_turn(line)
        turns.append(turn)

    dial = 50
    zeros = 0

    for turn in turns:
        dial, zeros = step(dial, zeros, turn)

    return zeros


def ActualPasswordOfTheDoor():
    try:
        with open("input.txt", "r") as file:
            content = file.read()
    except FileNotFoundError:
        print("error: input.txt file not found", file=sys.stderr)
        sys.exit(1)
    except IOError as e:
        print(f"error reading file: {e}", file=sys.stderr)
        sys.exit(1)

    try:
        result = process(content)
        print(f"result = {result}")
    except ValueError as e:
        print(f"error processing input: {e}", file=sys.stderr)
        sys.exit(1)


if __name__ == "__main__":
    ActualPasswordOfTheDoor()
