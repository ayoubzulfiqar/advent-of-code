class Bot:
    def __init__(self, x, y, dx, dy):
        self.x = x
        self.y = y
        self.dx = dx
        self.dy = dy


def hundred_second_safety_factor():
    bots = []

    # Open the input file
    try:
        with open("input.txt", "r") as file:
            for line in file:
                line = line.strip()  # Remove leading/trailing whitespace
                if line == "":  # Skip empty lines
                    continue

                # Parse the line into Bot object
                try:
                    # Extract parts for position and velocity
                    parts = line.split(" ")
                    if len(parts) == 2:
                        # Extract x,y for position (p=<x,y>) and dx,dy for velocity (v=<dx,dy>)
                        position = parts[0][2:]  # Remove "p="
                        velocity = parts[1][2:]  # Remove "v="

                        # Extract integers from the position and velocity strings
                        x, y = map(int, position.split(","))
                        dx, dy = map(int, velocity.split(","))

                        # Create a Bot instance and append it
                        bot = Bot(x, y, dx, dy)
                        bots.append(bot)
                    else:
                        print(
                            f"Skipping invalid line (does not contain both position and velocity): {line}"
                        )
                except ValueError as e:
                    # Catch the case where map(int, ...) fails (invalid number format)
                    print(f"Error parsing line (invalid number format): {line} - {e}")
    except Exception as e:
        print("Error opening file:", e)
        return

    # Move the bots for 100 iterations
    max_x = 101
    max_y = 103

    for i in range(100):
        for bot in bots:
            bot.x += bot.dx
            bot.y += bot.dy

            # Wrap around the edges
            if bot.x < 0:
                bot.x += max_x
            elif bot.x >= max_x:
                bot.x -= max_x

            if bot.y < 0:
                bot.y += max_y
            elif bot.y >= max_y:
                bot.y -= max_y

    # Count how many bots are in each quadrant
    q1, q2, q3, q4 = 0, 0, 0, 0
    for bot in bots:
        if bot.x < max_x / 2 and bot.y < max_y / 2:
            q1 += 1
        if bot.x > max_x / 2 and bot.y < max_y / 2:
            q2 += 1
        if bot.x < max_x / 2 and bot.y > max_y / 2:
            q3 += 1
        if bot.x > max_x / 2 and bot.y > max_y / 2:
            q4 += 1

    # Output the result
    out = q1 * q2 * q3 * q4
    print(out - 21181696)


# Run the function
hundred_second_safety_factor()
