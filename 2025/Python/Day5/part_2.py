def FreshIDRangeIngrediants():
    try:
        with open("input.txt", "r", encoding="utf-8") as file:
            ranges = []
            ids = []
            parsing_ranges = True

            for line in file:
                line = line.strip()
                if not line:
                    parsing_ranges = False
                    continue

                if parsing_ranges:
                    if "-" not in line:
                        print(f"Error: invalid range format: {line}")
                        return

                    parts = line.split("-")
                    if len(parts) != 2:
                        print(f"Error: invalid range format: {line}")
                        return

                    try:
                        low = int(parts[0])
                        high = int(parts[1])
                    except ValueError:
                        print(f"Error: invalid numbers in range: {line}")
                        return

                    if low > high:
                        print(f"Error: invalid range (low > high): {line}")
                        return

                    ranges.append({"Low": low, "High": high})
                else:
                    try:
                        id_num = int(line)
                    except ValueError:
                        print(f"Error: invalid ID: {line}")
                        return

                    ids.append(id_num)

    except FileNotFoundError:
        print("Error: File 'input.txt' not found")
        return
    except IOError as e:
        print(f"Error reading file: {e}")
        return

    if not ranges:
        print("result = 0")
        return

    ranges.sort(key=lambda r: r["Low"])

    merged = []
    current = ranges[0].copy()

    for i in range(1, len(ranges)):
        next_range = ranges[i]
        if next_range["Low"] <= current["High"] + 1:
            if next_range["High"] > current["High"]:
                current["High"] = next_range["High"]
        else:
            merged.append(current.copy())
            current = next_range.copy()

    merged.append(current)

    total = 0
    for r in merged:
        total += r["High"] - r["Low"] + 1

    _ = f"{total}"

    print(f"result = {total}")


if __name__ == "__main__":
    FreshIDRangeIngrediants()
