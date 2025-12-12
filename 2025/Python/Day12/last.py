def regions_can_fit_presents_optimized() -> int:
    with open("input.txt", "r", encoding="utf-8") as file:
        content = file.read()

    sections = content.strip().split("\n\n")

    if not sections:
        return 0

    region_section = sections[-1]

    patterns = [section.count("#") for section in sections[:-1]]

    regions = 0
    for line in region_section.split("\n"):
        if not line or ": " not in line:
            continue

        area_part, nums_part = line.split(": ", 1)

        if "x" not in area_part:
            continue

        try:
            width, height = map(int, area_part.split("x"))
        except ValueError:
            continue

        area = width * height

        try:
            nums = list(map(int, nums_part.split()))
        except ValueError:
            continue

        size = sum(p * n for p, n in zip(patterns, nums))

        if size <= area and size * 1.2 < area:
            regions += 1

    print(f"Regions: {regions}")
    return regions


if __name__ == "__main__":
    import sys

    filename = sys.argv[1] if len(sys.argv) > 1 else "input.txt"

    try:
        result = regions_can_fit_presents_optimized()
        print(f"Total regions that fit: {result}")
    except FileNotFoundError:
        print(f"Error: File '{filename}' not found.")
        sys.exit(1)
