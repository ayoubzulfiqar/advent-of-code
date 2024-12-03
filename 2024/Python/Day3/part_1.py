import re


def additionOfMultiplication():
    with open("./input.txt", "r") as file:
        corruptedMemory = file.read()
    pattern = r"mul\((\d+),(\d+)\)"
    matches = re.findall(pattern, corruptedMemory)

    # Calculate the sum of the products
    total = sum(int(x) * int(y) for x, y in matches)
    print(total)
    return total
