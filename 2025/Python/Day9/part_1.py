from typing import List, Tuple


def getCoords(input_str: str) -> List[Tuple[int, int]]:
    coords = []

    for line in input_str.strip().split("\n"):
        line = line.strip()
        if not line:
            continue

        if "," not in line:
            continue

        parts = line.split(",")
        if len(parts) != 2:
            continue

        try:
            x = int(parts[0].strip())
            y = int(parts[1].strip())
            coords.append((x, y))
        except ValueError:
            continue

    return coords


def largestAreaForRectangle(coords: List[Tuple[int, int]]) -> int:
    max_area = 0

    for i in range(len(coords)):
        x1, y1 = coords[i]
        for j in range(i + 1, len(coords)):
            x2, y2 = coords[j]

            width = abs(x1 - x2) + 1
            height = abs(y1 - y2) + 1

            area = width * height
            if area > max_area:
                max_area = area

    return max_area


def main():
    with open("input.txt", "r", encoding="utf-8") as f:
        con = f.read()
    coords = getCoords(con)
    result = largestAreaForRectangle(coords)
    print(f"Part 1 Answer: {result}")


if __name__ == "__main__":
    main()
