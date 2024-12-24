import re
from typing import Dict, List


def parseInput() -> List[List[str]]:
    with open("input.txt", "r") as f:
        return [part.split("\n") for part in f.read().strip().split("\n\n")]


def simulateGates(data):
    state: Dict[str, int] = {}
    for wire in data[0]:
        name, value = wire.split(": ")
        state[name] = int(value)

    loop = True
    while loop:
        should_loop_again = False
        for gate in data[1]:
            match = re.match(r"^(.*) (AND|OR|XOR) (.*) -> (.*)$", gate)
            if match:
                left, operator, right, output = match.groups()
                if left not in state or right not in state:
                    should_loop_again = True
                    continue
                if operator == "AND":
                    state[output] = state[left] & state[right]
                elif operator == "OR":
                    state[output] = state[left] | state[right]
                elif operator == "XOR":
                    state[output] = state[left] ^ state[right]
        loop = should_loop_again

    return state


# Part -1
def decimalWireZ() -> int:
    data = parseInput()
    state = simulateGates(data)
    bits = "".join(
        str(state[name])
        for name, _ in sorted(
            filter(lambda x: x[0].startswith("z"), state.items()),
            key=lambda x: x[0],
            reverse=True,
        )
    )
    return int(bits, 2)


# print("Part 1", decimalWireZ())
