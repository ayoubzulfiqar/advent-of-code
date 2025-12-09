from typing import List, Tuple


def getCoordinates(input_str: str) -> List[Tuple[int, int]]:
    coords = []
    for line in input_str.strip().split("\n"):
        if not line:
            continue
        comma_idx = line.find(",")
        if comma_idx == -1:
            continue
        try:
            x = int(line[:comma_idx].strip())
            y = int(line[comma_idx + 1 :].strip())
            coords.append((x, y))
        except ValueError:
            continue
    return coords


def largestAreaWithGreenAndRedTiles(coords: List[Tuple[int, int]]) -> int:
    if not coords:
        return 0

    n = len(coords)
    max_area = 0

    poly_edges = []
    for i in range(n):
        u = coords[i]
        v = coords[(i + 1) % n]
        poly_edges.append((u, v))

    xs = [p[0] for p in coords]
    ys = [p[1] for p in coords]
    min_x, max_x = min(xs), max(xs)
    min_y, max_y = min(ys), max(ys)

    pip_cache = {}

    for i in range(n):
        x1, y1 = coords[i]

        for j in range(i + 1, n):
            x2, y2 = coords[j]

            if x1 < x2:
                rx1, rx2 = x1, x2
            else:
                rx1, rx2 = x2, x1

            if y1 < y2:
                ry1, ry2 = y1, y2
            else:
                ry1, ry2 = y2, y1

            width = rx2 - rx1 + 1
            height = ry2 - ry1 + 1
            area = width * height

            if area <= max_area:
                continue

            if rx1 < min_x or rx2 > max_x or ry1 < min_y or ry2 > max_y:
                continue

            if isValidRectangleOptimized(
                rx1, rx2, ry1, ry2, coords, poly_edges, pip_cache
            ):
                max_area = area

    return max_area


def isValidRectangleOptimized(
    x1: int,
    x2: int,
    y1: int,
    y2: int,
    poly: List[Tuple[int, int]],
    poly_edges: List[Tuple[Tuple[int, int], Tuple[int, int]]],
    pip_cache: dict,
) -> bool:
    mx = x1 + x2
    my = y1 + y2
    cache_key = (mx, my)

    if cache_key in pip_cache:
        if not pip_cache[cache_key]:
            return False
    else:
        in_poly = isPointInPolyOptimized(mx, my, poly, poly_edges)
        pip_cache[cache_key] = in_poly
        if not in_poly:
            return False

    for (ux, uy), (vx, vy) in poly_edges:
        if ux == vx:
            ex = ux
            if ex <= x1 or ex >= x2:
                continue

            if uy < vy:
                ey_min, ey_max = uy, vy
            else:
                ey_min, ey_max = vy, uy

            overlap_y_min = y1 if y1 > ey_min else ey_min
            overlap_y_max = y2 if y2 < ey_max else ey_max
            if overlap_y_min < overlap_y_max:
                return False

        else:
            ey = uy
            if ey <= y1 or ey >= y2:
                continue

            if ux < vx:
                ex_min, ex_max = ux, vx
            else:
                ex_min, ex_max = vx, ux

            overlap_x_min = x1 if x1 > ex_min else ex_min
            overlap_x_max = x2 if x2 < ex_max else ex_max
            if overlap_x_min < overlap_x_max:
                return False

    return True


def isPointInPolyOptimized(
    x: int,
    y: int,
    poly: List[Tuple[int, int]],
    poly_edges: List[Tuple[Tuple[int, int], Tuple[int, int]]],
) -> bool:
    n = len(poly)

    for (ux, uy), (vx, vy) in poly_edges:
        ux2 = ux * 2
        uy2 = uy * 2
        vx2 = vx * 2
        vy2 = vy * 2

        if ux == vx and ux2 == x:
            if uy2 <= vy2:
                if uy2 <= y <= vy2:
                    return True
            elif vy2 <= y <= uy2:
                return True
        elif uy == vy and uy2 == y:
            if ux2 <= vx2:
                if ux2 <= x <= vx2:
                    return True
            elif vx2 <= x <= ux2:
                return True

    intersections = 0
    j = n - 1
    for i in range(n):
        ux, uy = poly[i]
        vx, vy = poly[j]

        if ux == vx:
            uy2 = uy * 2
            vy2 = vy * 2
            ex = ux * 2

            if uy2 <= vy2:
                min_y, max_y = uy2, vy2
            else:
                min_y, max_y = vy2, uy2

            if min_y <= y < max_y and ex > x:
                intersections += 1

        j = i

    return (intersections & 1) == 1


def main():
    with open("input.txt", "r", encoding="utf-8") as f:
        con = f.read()

    coords = getCoordinates(con)

    n = len(coords)
    print(f"Processing {n} points...")

    # For 496 points, we have ~122,760 rectangle checks (n*(n-1)/2)
    # Each check does O(m) work where m=496 edges
    # So total ~61 million operations, which should be ~1-2 seconds

    result = largestAreaWithGreenAndRedTiles(coords)
    print(f"Part 2 Answer: {result}")


if __name__ == "__main__":
    main()
