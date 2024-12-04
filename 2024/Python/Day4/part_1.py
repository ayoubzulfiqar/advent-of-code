def xMasAppear() -> int:
    with open("./input.txt") as file:
        content = file.read()
    lines = content.strip().split("\n")
    grid = [list(line) for line in lines]
    count = 0
    XMAS = "XMAS"

    for i in range(len(grid)):
        for j in range(len(grid[0])):
            if grid[i][j] == "X":
                # horizontal
                found = True
                for k in range(4):
                    if j + k >= len(grid[0]) or grid[i][j + k] != XMAS[k]:
                        found = False
                        break
                if found:
                    count += 1

                # horizontal reverse
                found = True
                for k in range(4):
                    if j - k < 0 or grid[i][j - k] != XMAS[k]:
                        found = False
                        break
                if found:
                    count += 1

                # vertical
                found = True
                for k in range(4):
                    if i + k >= len(grid) or grid[i + k][j] != XMAS[k]:
                        found = False
                        break
                if found:
                    count += 1

                # vertical reverse
                found = True
                for k in range(4):
                    if i - k < 0 or grid[i - k][j] != XMAS[k]:
                        found = False
                        break
                if found:
                    count += 1

                # diagonal
                found = True
                for k in range(4):
                    if (
                        i + k >= len(grid)
                        or j + k >= len(grid[0])
                        or grid[i + k][j + k] != XMAS[k]
                    ):
                        found = False
                        break
                if found:
                    count += 1

                # diagonal reverse
                found = True
                for k in range(4):
                    if i - k < 0 or j - k < 0 or grid[i - k][j - k] != XMAS[k]:
                        found = False
                        break
                if found:
                    count += 1

                # off-diagonal
                found = True
                for k in range(4):
                    if (
                        i - k < 0
                        or j + k >= len(grid[0])
                        or grid[i - k][j + k] != XMAS[k]
                    ):
                        found = False
                        break
                if found:
                    count += 1

                # off-diagonal reverse
                found = True
                for k in range(4):
                    if i + k >= len(grid) or j - k < 0 or grid[i + k][j - k] != XMAS[k]:
                        found = False
                        break
                if found:
                    count += 1

    return count
