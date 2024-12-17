from typing import Dict, List, Set

Point = Dict[str, int]

DIRECTIONS: Dict[str, Point] = {
    "^": {"x": 0, "y": -1},
    ">": {"x": 1, "y": 0},
    "v": {"x": 0, "y": 1},
    "<": {"x": -1, "y": 0},
}


def allSumOfGPSCoordinates() -> int:
    with open("./input.txt") as file:
        con = file.read()
    # Parse the input into grid and instructions
    parts = [lines.split("\n") for lines in con.strip().split("\n\n")]
    grid = [list(line) for line in parts[0]]
    instructions = "".join(parts[1])
    width, height = len(grid[0]), len(grid)

    walls: Set[str] = set()
    boxes: List[Point] = []
    robot = {"x": 0, "y": 0}

    # Initialize walls, boxes, and robot position
    for y in range(height):
        for x in range(width):
            if grid[y][x] == "@":
                robot = {"x": x * 2, "y": y}
            if grid[y][x] == "#":
                walls.add(f"{x * 2},{y}")
                walls.add(f"{x * 2 + 1},{y}")
            if grid[y][x] == "O":
                boxes.append({"x": x * 2, "y": y})

    # Recursive function to try moving all boxes
    def move_box(
        collided_box: Point, direction: Point, movements: List[Dict[str, Point]]
    ) -> bool:
        # Try both positions of the moved box
        next_positions = [
            {
                "x": collided_box["x"] + direction["x"],
                "y": collided_box["y"] + direction["y"],
            },
            {
                "x": collided_box["x"] + 1 + direction["x"],
                "y": collided_box["y"] + direction["y"],
            },
        ]

        # If collided with a wall, stop all movements
        for next_pos in next_positions:
            if f"{next_pos['x']},{next_pos['y']}" in walls:
                return False

        # Find all boxes that are collided with
        collided_boxes = [
            box
            for box in boxes
            if any(
                (box["x"] == collided_box["x"] and box["y"] == collided_box["y"])
                is False
                and (
                    (box["x"] == next_pos["x"] or box["x"] + 1 == next_pos["x"])
                    and box["y"] == next_pos["y"]
                )
                for next_pos in next_positions
            )
        ]

        # If there are no collided boxes, all movements are good
        if not collided_boxes:
            return True

        # Check for conflicts
        conflicts = False
        for box in collided_boxes:
            if move_box(box, direction, movements):
                # If box can move and not already processed, add to movements
                if not any(
                    b["x"] == box["x"] and b["y"] == box["y"]
                    for b in [m["box"] for m in movements]
                ):
                    movements.append({"box": box, "direction": direction})
            else:
                # If box can't move, prevent any movements
                conflicts = True
                break

        return not conflicts

    # Process each instruction
    for instruction in instructions:
        direction = DIRECTIONS[instruction]
        position = {"x": robot["x"] + direction["x"], "y": robot["y"] + direction["y"]}

        # Only try to move if no wall is in the way
        if f"{position['x']},{position['y']}" not in walls:
            collided_box = next(
                (
                    box
                    for box in boxes
                    if (box["x"] == position["x"] or box["x"] + 1 == position["x"])
                    and box["y"] == position["y"]
                ),
                None,
            )

            # If there is a collided box, try to move all affected
            if collided_box is not None:
                movements: List[Dict[str, Point]] = []
                if move_box(collided_box, direction, movements):
                    for movement in movements:
                        movement["box"]["x"] += movement["direction"]["x"]
                        movement["box"]["y"] += movement["direction"]["y"]
                    collided_box["x"] += direction["x"]
                    collided_box["y"] += direction["y"]
                    robot = position
            else:
                robot = position

    # Calculate the score
    score = sum(box["y"] * 100 + box["x"] for box in boxes)
    print(score)
    return score
