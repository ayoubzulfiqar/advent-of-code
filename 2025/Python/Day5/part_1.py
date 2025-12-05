def IDFreshIngrediants():
    try:
        with open("input.txt", "r", encoding="utf-8") as file:
            lines = [line.rstrip("\n") for line in file]
    except FileNotFoundError:
        print("Error: File 'input.txt' not found")
        return
    except IOError as e:
        print(f"Error opening file: {e}")
        return

    ranges = []
    ids = []
    parsing_ranges = True

    for line in lines:
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

            ranges.append({"Low": low, "High": high})
        else:
            try:
                id_num = int(line)
            except ValueError:
                print(f"Error: invalid ID: {line}")
                return

            ids.append(id_num)

    count = 0
    for id_num in ids:
        for r in ranges:
            if r["Low"] <= id_num <= r["High"]:
                count += 1
                break

    print(f"result = {count}")


if __name__ == "__main__":
    IDFreshIngrediants()
