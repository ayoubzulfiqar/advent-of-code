def singleTachyonParticleTimelines():
    try:
        with open("input.txt", "r", encoding="utf-8") as f:
            lines = [line.rstrip("\n") for line in f]
    except FileNotFoundError:
        print("Error reading input file")
        return 0

    if not lines:
        print("0")
        return 0

    first_line = lines[0]
    start = -1
    for i, ch in enumerate(first_line):
        if ch == "S":
            start = i
            break

    if start == -1:
        print("0")
        return 0

    cache = {}

    def timelines(pos, remaining_lines):
        if not remaining_lines:
            return 1

        key = (tuple(remaining_lines), pos)

        if key in cache:
            return cache[key]

        current_line = remaining_lines[0]
        result = 0

        if 0 <= pos < len(current_line) and current_line[pos] == "^":
            left = timelines(pos - 1, remaining_lines[1:])
            right = timelines(pos + 1, remaining_lines[1:])
            result = left + right
        else:
            result = timelines(pos, remaining_lines[1:])

        cache[key] = result
        return result

    result = timelines(start, lines[1:])
    print(f"{result}")
    return result


if __name__ == "__main__":
    singleTachyonParticleTimelines()
