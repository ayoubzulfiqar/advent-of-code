def timesXMasAppears() -> int:
    with open("./input.txt") as file:
        content = file.read()
    lines = content.strip().split("\n")
    grid = [list(line) for line in lines]
    count = 0

    for i in range(1, len(grid) - 1):
        for j in range(1, len(grid[0]) - 1):
            if grid[i][j] == "A":
                # Check pattern M.M
                #              .A.
                #              S.S
                cond = (
                    grid[i - 1][j - 1] == "M"
                    and grid[i - 1][j + 1] == "M"
                    and grid[i + 1][j + 1] == "S"
                    and grid[i + 1][j - 1] == "S"
                )
                if cond:
                    count += 1

                # Check pattern S.S
                #              .A.
                #              M.M
                cond = (
                    grid[i - 1][j - 1] == "S"
                    and grid[i - 1][j + 1] == "S"
                    and grid[i + 1][j + 1] == "M"
                    and grid[i + 1][j - 1] == "M"
                )
                if cond:
                    count += 1

                # Check pattern M.S
                #              .A.
                #              M.S
                cond = (
                    grid[i - 1][j - 1] == "M"
                    and grid[i - 1][j + 1] == "S"
                    and grid[i + 1][j + 1] == "S"
                    and grid[i + 1][j - 1] == "M"
                )
                if cond:
                    count += 1

                # Check pattern S.M
                #              .A.
                #              S.M
                cond = (
                    grid[i - 1][j - 1] == "S"
                    and grid[i - 1][j + 1] == "M"
                    and grid[i + 1][j + 1] == "M"
                    and grid[i + 1][j - 1] == "S"
                )
                if cond:
                    count += 1

    return count
