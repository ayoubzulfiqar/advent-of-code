def randomNumber(seed: int) -> int:
    seed = ((seed << 6) ^ seed) % 16777216
    seed = ((seed >> 5) ^ seed) % 16777216
    seed = ((seed << 11) ^ seed) % 16777216
    return seed


def getMostBananas() -> int:
    with open("input.txt") as file:
        con = file.read()
    ranges: dict = {}
    numbers: map = map(int, con.split("\n"))

    for num in numbers:
        seed = num
        visited = set()
        changes = []

        for _ in range(2000):
            nextSeed = randomNumber(seed)
            changes.append((nextSeed % 10) - (seed % 10))
            seed = nextSeed

            if len(changes) == 4:
                key = ",".join(map(str, changes))
                if key not in visited:
                    if key not in ranges:
                        ranges[key] = []
                    ranges[key].append(nextSeed % 10)
                    visited.add(key)
                changes.pop(0)

    return max(sum(rangeValues) for rangeValues in ranges.values())
