def beamSplitTime():
    try:
        with open("input.txt", "r") as f:
            lines = [line.rstrip("\n") for line in f]
    except FileNotFoundError:
        print("Error reading input file")
        return 0

    if not lines:
        print("Part 1: 0")
        return 0

    first_line = lines[0]
    start = -1
    for i, ch in enumerate(first_line):
        if ch == "S":
            start = i
            break

    if start == -1:
        print("Part 1: 0")
        return 0

    beams = {start: True}
    count = 0

    for line in lines[1:]:
        splitters = {}
        for i, ch in enumerate(line):
            if ch == "^":
                splitters[i] = True

        splits = {}
        split_count = 0
        for pos in beams:
            if pos in splitters:
                splits[pos] = True
                split_count += 1

        new_beams = {}
        for pos in splits:
            if pos + 1 < len(line):
                new_beams[pos + 1] = True
            if pos - 1 >= 0:
                new_beams[pos - 1] = True

        remaining_beams = {}
        for pos in beams:
            if pos not in splits:
                remaining_beams[pos] = True

        final_beams = {}
        for pos in remaining_beams:
            final_beams[pos] = True
        for pos in new_beams:
            final_beams[pos] = True

        count += split_count
        beams = final_beams

    print(f"Part 1: {count}")
    return count


if __name__ == "__main__":
    beamSplitTime()
