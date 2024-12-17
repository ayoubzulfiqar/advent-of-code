from typing import Dict

# Define a Point type as a dictionary with x and y coordinates
Point = Dict[str, int]

# Define the directions as a dictionary mapping characters to Points
DIRECTIONS: Dict[str, Point] = {
    "^": {"x": 0, "y": -1},
    ">": {"x": 1, "y": 0},
    "v": {"x": 0, "y": 1},
    "<": {"x": -1, "y": 0},
}


def sumOfGPSCoordinates() -> int:
    with open("./input.txt") as file:
        con = file.read()
    # Parse the input into grid and instructions
    parts = [lines.split("\n") for lines in con.strip().split("\n\n")]
    grid = [list(line) for line in parts[0]]
    instructions = "".join(parts[1])
    width, height = len(grid[0]), len(grid)

    # Define the recursive function to move a box
    def move_box(position: Point, direction: Point) -> bool:
        next_pos = {
            "x": position["x"] + direction["x"],
            "y": position["y"] + direction["y"],
        }

        if grid[next_pos["y"]][next_pos["x"]] == ".":
            # If the next spot is empty, swap positions
            grid[position["y"]][position["x"]], grid[next_pos["y"]][next_pos["x"]] = (
                grid[next_pos["y"]][next_pos["x"]],
                grid[position["y"]][position["x"]],
            )
            return True
        elif grid[next_pos["y"]][next_pos["x"]] == "#":
            # If the next spot is a wall, stop all boxes from moving
            return False
        else:
            # Only move the current box if the next box can move
            if move_box(next_pos, direction):
                (
                    grid[position["y"]][position["x"]],
                    grid[next_pos["y"]][next_pos["x"]],
                ) = (
                    grid[next_pos["y"]][next_pos["x"]],
                    grid[position["y"]][position["x"]],
                )
                return True

        # This should never be reached
        return False

    # Find the robot and clear its position
    robot = {"x": 0, "y": 0}
    for y in range(height):
        for x in range(width):
            if grid[y][x] == "@":
                robot = {"x": x, "y": y}
                grid[y][x] = "."

    # Process each instruction
    for instruction in instructions:
        direction = DIRECTIONS[instruction]
        position = {"x": robot["x"] + direction["x"], "y": robot["y"] + direction["y"]}

        # If there is a wall, don't move
        if grid[position["y"]][position["x"]] != "#":
            # If there is an empty spot, move without moving boxes
            if grid[position["y"]][position["x"]] == ".":
                robot = position
            # If there is a box, try to move all the boxes, then move
            elif grid[position["y"]][position["x"]] == "O" and move_box(
                position, direction
            ):
                robot = position

    # Tally all the box positions
    score = 0
    for y in range(height):
        for x in range(width):
            if grid[y][x] == "O":
                score += y * 100 + x
    print(score)
    return score
