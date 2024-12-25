class Grid:
    @staticmethod
    def parse(block):
        return [list(line) for line in block.splitlines()]


class Point:
    def __init__(self, x, y):
        self.x = x
        self.y = y

    def __add__(self, other):
        return Point(self.x + other.x, self.y + other.y)


DOWN = Point(0, 1)
UP = Point(0, -1)
ORIGIN = Point(0, 0)


def parse() -> str:
    with open("input.txt", "r") as f:
        return f.read().strip()


# Part 1
def uniqueLockANDKey():
    data = parse()
    locks = []
    keys = []
    result = 0

    for block in data.split("\n\n"):
        grid = Grid.parse(block)
        heights = 0

        if grid[ORIGIN.y][ORIGIN.x] == "#":
            for x in range(5):
                position = Point(x, 1)

                while position.y < len(grid) and grid[position.y][position.x] == "#":
                    position += DOWN

                heights = (heights << 4) + (position.y - 1)

            locks.append(heights)
        else:
            for x in range(5):
                position = Point(x, 5)

                while position.y >= 0 and grid[position.y][position.x] == "#":
                    position += UP

                heights = (heights << 4) + (5 - position.y)

            keys.append(heights)

    for lock in locks:
        for key in keys:
            if (lock + key + 0x22222) & 0x88888 == 0:
                result += 1

    return result


print(uniqueLockANDKey())
