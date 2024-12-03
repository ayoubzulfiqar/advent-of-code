import re


def resultOfEnabledMultiply():
    with open("./input.txt", "r") as file:
        corruptedMemory = file.read()
    # Regular expressions for each instruction type
    mulPattern = r"mul\((\d+),(\d+)\)"
    doPattern = r"do\(\)"
    dontPattern = r"don't\(\)"

    # Split memory into individual characters or instructions
    instructions = re.split(r"(do\(\)|don't\(\)|mul\(\d+,\d+\))", corruptedMemory)

    # Remove empty strings and irrelevant characters
    instructions = [instr for instr in instructions if instr and not instr.isspace()]

    # Tracking the state and the result
    enabled = True
    total = 0

    for instr in instructions:
        if re.match(doPattern, instr):
            enabled = True
        elif re.match(dontPattern, instr):
            enabled = False
        elif re.match(mulPattern, instr) and enabled:
            # Extract numbers from the mul(X, Y) instruction
            x, y = map(int, re.findall(r"\d+", instr))
            total += x * y
    return total
