class Bot:
    def __init__(self, x, y, dx, dy):
        self.x = x
        self.y = y
        self.dx = dx
        self.dy = dy


maxX = 101
maxY = 103


def robots_to_easter_egg():
    bots = []

    try:
        # Open the input file
        with open("input.txt", "r") as file:
            for line in file:
                line = line.strip()  # Remove leading/trailing whitespace
                if line:
                    try:
                        # Parse the line into Bot object (format: p=<x,y> v=<dx,dy>)
                        parts = line.split(" ")
                        position = parts[0][2:]  # Remove "p="
                        velocity = parts[1][2:]  # Remove "v="

                        # Extract integers from the position and velocity strings
                        x, y = map(int, position.split(","))
                        dx, dy = map(int, velocity.split(","))

                        # Create a Bot instance and append it
                        bot = Bot(x, y, dx, dy)
                        bots.append(bot)
                    except ValueError as e:
                        # Catch the case where map(int, ...) fails (invalid number format)
                        print(
                            f"Error parsing line (invalid number format): {line} - {e}"
                        )
    except Exception as e:
        print(f"Error opening file: {e}")
        return

    # Loop for 100,000 iterations
    for i in range(100000):
        # Initialize the grid and fill with '.'
        grid = [["." for _ in range(maxX)] for _ in range(maxY)]

        distinct = True

        # Update bot positions and check for collisions
        for bot in bots:
            bot.x += bot.dx
            bot.y += bot.dy

            # Wrap around the edges
            if bot.x < 0:
                bot.x += maxX
            elif bot.x >= maxX:
                bot.x -= maxX

            if bot.y < 0:
                bot.y += maxY
            elif bot.y >= maxY:
                bot.y -= maxY

            # Mark the grid cell
            if grid[bot.y][bot.x] == ".":
                grid[bot.y][bot.x] = "#"
            else:
                distinct = False

        # If all bots were distinct, print the grid
        if distinct:
            print(f"\nIter: {i+1}")
            # Uncomment this block to print the grid
            # for row in grid:
            #     print(''.join(row))
            break  # If you only want the first iteration with distinct bots, break after the first print


# Run the function
robots_to_easter_egg()
