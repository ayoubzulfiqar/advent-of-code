bytes = [
    x + y * 1j
    for r in open("input.txt").readlines()
    for x, y in [map(int, r.strip().split(","))]
]


def minimumStepToExist() -> int:
    start, goal, steps = 0, 70 + 70 * 1j, 0
    front, seen, walls = {start}, {start}, set(bytes[:1024])
    while front:
        new = set()
        steps += 1
        for pos in front:
            for d in [1, -1, 1j, -1j]:
                nPosition = pos + d
                if nPosition == goal:
                    return steps
                if not (0 <= nPosition.real <= 70 and 0 <= nPosition.imag <= 70):
                    continue
                if nPosition in walls or nPosition in seen:
                    continue
                seen.add(nPosition)
                new.add(nPosition)
        front = new
    return 0


# print(minimumStepToExist())
