def randomNumber(seed: int) -> int:
    seed = ((seed << 6) ^ seed) % 16777216
    seed = ((seed >> 5) ^ seed) % 16777216
    seed = ((seed << 11) ^ seed) % 16777216
    return seed


def sumOfTwoThousandthNumber() -> int:
    with open("input.txt") as file:
        con = file.read()
    numbers = map(int, con.split("\n"))
    total = 0
    for num in numbers:
        seed = num
        for _ in range(2000):
            seed = randomNumber(seed)
        total += seed
    return total
