from typing import List


def parseInput() -> List[List[str]]:
    with open("input.txt", "r") as f:
        return [part.split("\n") for part in f.read().strip().split("\n\n")]


def find(a, b, operator, gates):
    for gate in gates:
        if gate.startswith(f"{a} {operator} {b}") or gate.startswith(
            f"{b} {operator} {a}"
        ):
            return gate.split(" -> ").pop()
    return None


def swapANDJoinWires():
    data = parseInput()
    swapped = []
    c0 = None
    for i in range(45):
        n = str(i).zfill(2)
        m1, n1, r1, z1, c1 = None, None, None, None, None

        # Half adder logic
        m1 = find(f"x{n}", f"y{n}", "XOR", data[1])
        n1 = find(f"x{n}", f"y{n}", "AND", data[1])

        if c0:
            r1 = find(c0, m1, "AND", data[1])
            if not r1:
                m1, n1 = n1, m1
                swapped.extend([m1, n1])
                r1 = find(c0, m1, "AND", data[1])

            z1 = find(c0, m1, "XOR", data[1])

            if m1 and m1.startswith("z"):
                m1, z1 = z1, m1
                swapped.extend([m1, z1])

            if n1 and n1.startswith("z"):
                n1, z1 = z1, n1
                swapped.extend([n1, z1])

            if r1 and r1.startswith("z"):
                r1, z1 = z1, r1
                swapped.extend([r1, z1])

            c1 = find(r1, n1, "OR", data[1])

        if c1 and c1.startswith("z") and c1 != "z45":
            c1, z1 = z1, c1
            swapped.extend([c1, z1])

        c0 = c1 if c0 else n1

    return ",".join(sorted(swapped))


# print("Part 2", swapANDJoinWires())
