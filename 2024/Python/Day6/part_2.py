# Read the input file and process it
with open("input.txt", "r") as file:
    input_data = [line.strip() for line in file.readlines()]

rows = len(input_data)
cols = len(input_data[0])

# Direction vectors (up, right, down, left)
directions = [
    (-1, 0),  # Up
    (0, 1),  # Right
    (1, 0),  # Down
    (0, -1),  # Left
]

# Locate the guard's initial position and direction
start_row, start_col, start_dir = None, None, None
for r in range(rows):
    for c in range(cols):
        if input_data[r][c] in "^>v<":
            start_row, start_col = r, c
            start_dir = "^>v<".index(input_data[r][c])  # Map direction symbol to index
            break


# Function to simulate guard movement with an optional extra obstacle
def simulate_with_obstacle(obstacle_row, obstacle_col):
    guard_row, guard_col, guard_dir = start_row, start_col, start_dir
    visited = set()
    visited.add(f"{guard_row},{guard_col},{guard_dir}")

    while True:
        dr, dc = directions[guard_dir]
        next_row = guard_row + dr
        next_col = guard_col + dc

        # Check if the next position is outside the grid
        if next_row < 0 or next_row >= rows or next_col < 0 or next_col >= cols:
            return False  # Guard exits the grid

        # Treat the additional obstacle as if it's a `#`
        next_cell = (
            "#"
            if (next_row == obstacle_row and next_col == obstacle_col)
            else input_data[next_row][next_col]
        )
        if next_cell == "#":
            # Obstacle ahead, turn right
            guard_dir = (guard_dir + 1) % 4
        else:
            # Move forward
            guard_row = next_row
            guard_col = next_col

        state = f"{guard_row},{guard_col},{guard_dir}"
        if state in visited:
            return True  # Loop detected
        visited.add(state)


# Count valid positions for the new obstruction
valid_positions = 0

for r in range(rows):
    for c in range(cols):
        # Skip positions that are not empty or are the starting position
        if input_data[r][c] == "#" or (r == start_row and c == start_col):
            continue

        # Simulate guard movement with an obstacle at (r, c)
        if simulate_with_obstacle(r, c):
            valid_positions += 1

print(f"Number of valid positions for a new obstruction: {valid_positions}")
